package dispatcher

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/bot"
	"awesomeProject/bot/feedback"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"
	"os"

	handler "awesomeProject/bot/botSchedule"
	trackerHandler "awesomeProject/bot/botTracker"

	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
)

func CommandsHandling(message *tgbotapi.Message) {
	switch message.Command() {
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
			log.Println(err)
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
			log.Println(err)
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
	default:
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Неизвестная команда (%s)\nВоспользуйтесь /help", message.Text)))
	}
}

func MessageHandler(message *tgbotapi.Message) {
	res, err := db.Get(message.Chat.ID, "temp")
	if err != nil {
		log.Println(err)
	}
	if res == "snils" {
		re := regexp.MustCompile(message.Text)
		snils := strings.Join(re.FindAllString(message.Text, -1), "")
		status, name, stage := db.CheckSnils(snils)
		if status {
			err = db.CreateNewTrackerUser(message, name, stage)
			_ = db.Update(message.Chat.ID, "temp", "")
			if err != nil {
				bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Неудалось заполнить базу данных"))
			} else {
				msg := tgbotapi.NewMessage(message.Chat.ID, "Готово!\nВот некоторый функционал РСОШ Трекера")
				msg.ReplyMarkup = bot.BuilderMenuTracker
				bot.Send(msg)
			}
		}
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите опцию")
		model, err := db.Get(message.Chat.ID, "bot")
		if err != nil {
			log.Println(err)
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
						log.Println(err)
						return false
					}
					image, err := timetable.DrawTimetable(lessons,
						fmt.Sprintf("%s, нед: %s, день: %s", slice[0], slice[1], slice[2]), false)
					if err != nil {
						log.Println(err)
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
					log.Println(err)
					return false
				}
				image, err := timetable.DrawTimetable(lessons,
					fmt.Sprintf("%s, нед: %s, день: %s",
						slice[0], lexicon.Week[week], lexicon.Day[day]), false)
				if err != nil {
					log.Println(err)
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
			log.Println(err)
			return
		}

		xlsx := excelize.NewFile()
		sheet, err := xlsx.NewSheet("Sheet1")
		if err != nil {
			log.Println(err)
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
			log.Println(err)
			return
		}

		bot.SendFile(message.Chat.ID, filename, "Все записи.xlsx", "time")
	}
}