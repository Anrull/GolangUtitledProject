package main

import (
	"awesomeProject/bot"
	handler "awesomeProject/bot/bot_timetable"
	trackerHandler "awesomeProject/bot/bot_tracker"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"log"
	"strings"
)

var Bot = bot.Bot

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
				case "/schedule":
					handler.Schedule(message, true)
				case "/tomorrow":
					handler.Schedule(message, false)
				case "/week":
					handler.Week(message, false)
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
					case "who":
						handler.WhoAreYouHandler(query.Message.Chat.ID, query.Message.MessageID, lstQ[2])
					case "choice":
						handler.ChoiceTimetableHandler(query.Message.Chat.ID, query.Message.MessageID, lstQ[3], lstQ[2])
					}
				} else if lstQ[0] == "tracker" {
					fmt.Println(lstQ)
					switch lstQ[2] {
					case "olimp":
						//fmt.Println("work")
						trackerHandler.OlimpsCallbacksHandler(query.Message, lstQ[3], lstQ[4], lstQ[5], lstQ[6])

					}
				}
			}
		}()
	}
}
