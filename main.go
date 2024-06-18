package main

import (
	"awesomeProject/bot"
	handler "awesomeProject/bot/bot_timetable"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"log"
	"strings"
)

var Bot = bot.Bot

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://github.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	//Bot.Debug = true

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := Bot.GetUpdatesChan(u)

	for update := range updates {
		go func() {
			if update.Message != nil {
				message := update.Message
				switch message.Text {
				case "/start":
					handler.Start(message)
				case "/help":
					handler.Help(message)
				case "/time":
					handler.Time(message)
				case "/days":
					handler.Days(message)
				default:
					tgbotapi.NewMessage(message.Chat.ID, message.Text)
				}
			} else if update.CallbackQuery != nil {
				query := update.CallbackQuery
				lstQ := strings.Split(query.Data, ";")
				if lstQ[0] == "timetable" {
					switch lstQ[1] {
					case "days":
						handler.DaysHandler(query.From.ID, lstQ[2], lstQ[3])
					}
				}
			}
		}()
	}
}
