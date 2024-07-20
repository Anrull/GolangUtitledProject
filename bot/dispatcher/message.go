package dispatcher

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/bot"
	"awesomeProject/bot/feedback"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"
	"os"
	"slices"

	handler "awesomeProject/bot/botSchedule"
	trackerHandler "awesomeProject/bot/botTracker"

	"fmt"
	"log"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	case "my_olimps":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите фильтр")
		msg.ReplyMarkup = bot.BuilderChoiceTrackerFilter
		bot.Send(msg)
	case "shutdown":
		ids, err := db.GetAdminIds()
		if err != nil {
			return
		}

		if slices.Contains(ids, message.Chat.ID) {
			bot.Send(tgbotapi.NewMessage(1705933876, "Бот выключен"))
			os.Exit(0)
		}
	case "admin":
		ids, err := db.GetAdminIds()
		if err != nil {
			return
		}

		if slices.Contains(ids, message.Chat.ID) {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Панель Администратора")
			msg.ReplyToMessageID = message.MessageID
			msg.ReplyMarkup = bot.AdminPanel
			bot.Send(msg)
		}
	case "lock":
		ids, err := db.GetAdminIds()
		if err != nil {
			return
		}

		if slices.Contains(ids, message.Chat.ID) {
			bot.TechnicalWork = true
		}
	case "unlock":
		ids, err := db.GetAdminIds()
		if err != nil {
			return
		}

		if slices.Contains(ids, message.Chat.ID) {
			bot.TechnicalWork = false
		}
	case "profile":
		Profile(message)
	case "fb":
		ids, err := db.GetAdminIds()
		if err != nil {
			return
		}

		if slices.Contains(ids, message.Chat.ID) {
			feedback.HandlerInfo(message, "nowWeek")
		}
	case "add_admin":
		user, err := db.GetInfoAboutPerson(message.Chat.ID)
		if err != nil {
			return
		}

		if user.Admin != "SuperAdmin" {
			bot.Send(tgbotapi.NewMessage(message.Chat.ID, "У вас нет прав супер администратора"))
			return
		}

		msg := strings.Replace(message.Text, "/add_admin ", "", 1)
		role := "id"
		value := msg
		if strings.Contains(msg, "@") {
			role = "nick"
			value = strings.Replace(msg, "@", "", 1)
		}
		err = db.AddAdmin(value, role)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Не удалось добавить администратора (%s)", err.Error())))
			log.Println(err)
			return
		}
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Администратор %s добавлен", value)))
	default:
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Неизвестная команда (%s)", message.Text)))
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
					image, err := timetable.DrawTimetableTest(lessons,
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
				image, err := timetable.DrawTimetableTest(lessons,
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
	for _, str := range slice {
		if str == value {
			return true
		}
	}
	return false
}
