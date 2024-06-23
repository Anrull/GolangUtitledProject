package dispatcher

import (
	"awesomeProject/bot"
	"awesomeProject/data/db"

	handler "awesomeProject/bot/bot_timetable"
	trackerHandler "awesomeProject/bot/bot_tracker"

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
	default:
		Bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Неизвестная команда (%s)", message.Text)))
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
			db.Update(message.Chat.ID, "temp", "")
			if err != nil {
				Bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Неудалось заполнить базу данных"))
			} else {
				msg := tgbotapi.NewMessage(message.Chat.ID, "Готово!\nВот некоторый функционал РСОШ Трекера")
				msg.ReplyMarkup = bot.BuilderMenuTracker
				Bot.Send(msg)
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
		Bot.Send(msg)
	}
}
