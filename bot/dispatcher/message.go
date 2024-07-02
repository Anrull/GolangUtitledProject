package dispatcher

import (
	"awesomeProject/bot"
	"awesomeProject/bot/feedback"
	"awesomeProject/data/db"
	"os"

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
	case "shutdown":
		if message.Chat.ID == 1705933876 {
			bot.Send(tgbotapi.NewMessage(1705933876, "Бот выключен"))
			os.Exit(0)
		}
	case "admin":
		if message.Chat.ID == 1705933876 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Панель Администратора")
			msg.ReplyToMessageID = message.MessageID
			msg.ReplyMarkup = bot.AdminPanel
			bot.Send(msg)
		}
	case "db":
		msg := tgbotapi.NewMessage(message.Chat.ID, "dfgtyhui")
		msg.ReplyMarkup = feedback.GetFeedbackCallback("9B", "Геометрия", "09.05.2024")
		bot.Send(msg)
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
			msg.ReplyMarkup = bot.MenuScheduleBotKeyboard
		} else {
			msg.ReplyMarkup = bot.BuilderMenuTracker
		}
		bot.Send(msg)
	}
}
