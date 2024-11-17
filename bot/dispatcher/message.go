package dispatcher

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/bot"
	"awesomeProject/bot/feedback"
	"awesomeProject/bot/lexicon"
	"awesomeProject/bot/logger"
	"awesomeProject/data/db"

	handler "awesomeProject/bot/botSchedule"
	trackerHandler "awesomeProject/bot/botTracker"

	"os"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
)

func CommandsHandling(message *tgbotapi.Message) {
	switch message.Command() {
	case "delete_me":
		handler.DeleteMe(message)
	case "start":
		handler.Start(message)
	case "help":
		handler.Help(message)
	case "time":
		handler.Time(message, false)
	case "days":
		handler.Days(message)
	case "schedule":
		handler.Schedule(message, true)
	case "tomorrow":
		handler.Schedule(message, false)
	case "week":
		handler.Week(message, false)
	case "add":
		trackerHandler.AddRecord(message, false)
	case "newsletter":
		value, err := db.Get(message.Chat.ID, "newsletter")
		if err != nil {
			logger.Error("", err)
			return
		}

		if value == "1" {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Уведомления отключены")
			msg.ReplyToMessageID = message.MessageID
			bot.Send(msg)
			err = db.Update(message.Chat.ID, "newsletter", "0")
		} else {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Уведомления включены")
			msg.ReplyToMessageID = message.MessageID
			bot.Send(msg)
			err = db.Update(message.Chat.ID, "newsletter", "1")
		}
		if err != nil {
			bot.Send(tgbotapi.NewMessage(1705933876, err.Error()))
			logger.Error("", err)
		}
	case "my_olimps":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите фильтр")
		msg.ReplyMarkup = bot.BuilderChoiceTrackerFilter
		bot.Send(msg)
	case "shutdown":
		shutdown(message)
	case "admin":
		if isAdmin(message) {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Панель Администратора")
			msg.ReplyToMessageID = message.MessageID
			msg.ReplyMarkup = bot.AdminPanel
			bot.Send(msg)
		}
	case "lock":
		if isAdmin(message) {
			bot.TechnicalWork = true
		}
	case "unlock":
		if isAdmin(message) {
			bot.TechnicalWork = false
		}
	case "profile":
		Profile(message)
	case "fb":
		if isAdmin(message) {
			feedback.HandlerInfo(message, "nowWeek")
		}
	case "add_admin":
		AddAdmin(message)
	case "get_tracker":
		getTracker(message)
		defer os.Remove("data/temp/Все записи.xlsx")
	case "add_student":
		defaultTextMessage := "Образец:\n/add_student\nФамилия Имя Отчество\n10А\n12345678910"
		if isAdmin(message) {
			lstQ := strings.Split(message.Text, "\n")
			if len(lstQ) == 4 {
				if len(strings.Split(lstQ[1], " ")) == 3 {
					re := regexp.MustCompile(lstQ[3])
					snils := strings.Join(re.FindAllString(lstQ[3], -1), "")
					snils = strings.Map(func(r rune) rune {
						if r >= '0' && r <= '9' {
							return r
						}
						return -1
					}, snils)
					if func() int {var count int
						for range snils {
							count++
						}
						return count }() != 11 {
						bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Неверный формат СНИЛС"))
						return
					}
					stages := lexicon.ExampleStages
					if stages[lstQ[2]] != "" {
						lstQ[2] = stages[lstQ[2]]
					} else {
						bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("%s (%s)\n\n%s", "Неверный формат класса", lstQ[2], defaultTextMessage)))
						return
					}
					err := db.AddStudent(lstQ[1], snils, lstQ[2])
					if err != nil {
						bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("%s\n\n%s", err.Error(), defaultTextMessage)))
						return
					}
					bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Студент добавлен"))
				} else {
					bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("%s\n\n%s", "Неверный формат ФИО", defaultTextMessage)))
				}
			} else {
				bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("%s\n\n%s", "Неверный формат запроса", defaultTextMessage)))
			}
		} else {
			bot.Send(tgbotapi.NewMessage(message.Chat.ID, "У вас нет прав администратора"))
			return
		}
	default:
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Неизвестная команда (%s)\n\nВоспользуйтесь /help", message.Text)))
	}
}

func MessageHandler(message *tgbotapi.Message) {
	res, err := db.Get(message.Chat.ID, "temp")
	if err != nil {
		logger.Error("", err)
	}
	if res == "snils" {
		re := regexp.MustCompile(message.Text)
		snils := strings.Join(re.FindAllString(message.Text, -1), "")
		snils = strings.Map(func(r rune) rune {
			if r >= '0' && r <= '9' {
				return r
			}
			return -1
		}, snils)
		status, name, stage := db.CheckSnils(snils)
		if status {
			err = db.CreateNewTrackerUser(message, name, stage)
			_ = db.Update(message.Chat.ID, "temp", "")
			if err != nil {
				logger.Error("", err)
				bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Неудалось заполнить базу данных"))
			} else {
				msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Готово!\nВот некоторый функционал РСОШ Трекера\n\n%s, %s", name, stage))
				t := bot.CopyInlineKeyboard(bot.BuilderMenuTracker)
				msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
					append(t.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Удалить аккаунт", "menu;tracker;delete_me"),
					))...,
				)
				bot.Send(msg)
			}
		} else {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Неверный СНИЛС")
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Отмена", "menu;tracker;snils"),
				),
			)
			bot.Send(msg)
		}
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите опцию")
		model, err := db.Get(message.Chat.ID, "bot")
		if err != nil {
			logger.Error("", err)
		}
		if model == "bot-schedule" {
			sliceMessage := strings.Split(message.Text, " ")
			if isValid(message, sliceMessage, len(sliceMessage)) {
				return
			}
			msg.ReplyMarkup = bot.MenuScheduleBotKeyboard
		} else {
			msg.ReplyMarkup = bot.BuilderMenuTracker
		}
		bot.Send(msg)
	}
}

func isValid(message *tgbotapi.Message, slice []string, num int) bool {
	colors, _ := db.GetColorByUserID(message.Chat.ID)
	if num == 3 {
		if in(lexicon.Stages, slice[0]) {
			if in([]string{"н", "ч"}, slice[1]) {
				if in(lexicon.ListDays, slice[2]) {
					if slice[1] == "н" {
						slice[1] = "1"
					} else {
						slice[1] = "0"
					}
					lessons, err := timetable.GetTimetableText(slice[1],
						lexicon.DayTextToInt[slice[2]], slice[0])
					if err != nil {
						logger.Error("", err)
						return false
					}

					extraLessons, err := timetable.GetExtraTimetableText(slice[1],
						lexicon.DayTextToInt[slice[2]], slice[0])
					
					if err != nil {
						logger.Error("", err)
						return false
					}

					lessons = timetable.Merge(lessons, extraLessons)

					image, err := timetable.DrawTimetable(lessons,
						fmt.Sprintf("%s, нед: %s, день: %s", slice[0], func(s string) string {if s == "0" {return "чет"}; return "нечет"}(slice[1]), slice[2]), false, colors...)
					if err != nil {
						logger.Error("", err)
						return false
					}
					handler.SendPhotoByte(message.Chat.ID, image)
					return true
				}
			}
		}
		return false
	}
	if num == 2 {
		if in(lexicon.Stages, slice[0]) {
			day := strings.ToLower(slice[1])
			if in([]string{"сегодня", "завтра"}, day) {
				var week string

				if day == "сегодня" {
					week, _ = timetable.GetWeek(false, true)
					day = timetable.GetDayToday()
				} else {
					week, _ = timetable.GetWeek(true, true)
					day = timetable.GetDayTomorrow()
				}
				lessons, err := timetable.GetTimetableText(week, day, slice[0])
				if err != nil {
					logger.Error("", err)
					return false
				}

				extraLessons, err := timetable.GetExtraTimetableText(week, day, slice[0])
				
				if err != nil {
					logger.Error("", err)
					return false
				}

				lessons = timetable.Merge(lessons, extraLessons)

				image, err := timetable.DrawTimetable(lessons,
					fmt.Sprintf("%s, нед: %s, день: %s",
						slice[0], lexicon.Week[week], lexicon.Day[day]), false, colors...)
				if err != nil {
					logger.Error("", err)
					return false
				}
				handler.SendPhotoByte(message.Chat.ID, image)
				return true
			}
		}
	}

	return false
}

func in(slice []string, value string) bool {
	return slices.Contains(slice, value)
}

func isAdmin(message *tgbotapi.Message) bool {
	ids, err := db.GetAdminIds()
	if err != nil {
		return false
	}

	if slices.Contains(ids, message.Chat.ID) {
		return true
	}
	return false
}

func getTracker(message *tgbotapi.Message, filenames ...string) {
	if isAdmin(message) {
		var filename string
		if len(filenames) > 0 {
			filename = filenames[0]
		} else {
			filename = "data/temp/Все записи.xlsx"
		}
		records, err := db.GetAllRecords()
		if err != nil {
			logger.Error("", err)
			return
		}

		xlsx := excelize.NewFile()
		sheet, err := xlsx.NewSheet("Sheet1")
		if err != nil {
			logger.Error("", err)
			return
		}
		xlsx.SetActiveSheet(sheet)

		row := 1
		for _, record := range *records {
			xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(row), record.Date)
			xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(row), record.Name)
			xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(row), record.Class)
			xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(row), record.Olimps)
			xlsx.SetCellValue("Sheet1", "E"+strconv.Itoa(row), record.Stage)
			xlsx.SetCellValue("Sheet1", "F"+strconv.Itoa(row), record.Subjects)
			xlsx.SetCellValue("Sheet1", "G"+strconv.Itoa(row), record.Teachers)
			row++
		}

		err = xlsx.SaveAs(filename)
		if err != nil {
			logger.Error("", err)
			return
		}

		bot.SendFile(message.Chat.ID, filename, "Все записи.xlsx", "time")
	} else {
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Недостаточно прав доступа"))
	}
}
