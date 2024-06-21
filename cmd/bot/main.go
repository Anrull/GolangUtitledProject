package main

import (
	"awesomeProject/bot"
	"awesomeProject/data/db"

	handler "awesomeProject/bot/bot_timetable"
	trackerHandler "awesomeProject/bot/bot_tracker"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"log"
	"regexp"
	"strings"
)

var Bot = bot.Bot

func main() {
	//Bot.Debug = true

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := Bot.GetUpdatesChan(u)

	for update := range updates {
		go func() {
			if update.Message != nil {
			}
			if update.Message != nil {
				message := update.Message
				if message.IsCommand() {
					switch message.Command() {
					case "start":
						handler.Start(message)
					case "help":
						handler.Help(message)
					case "time":
						handler.Time(message)
					case "days":
						handler.Days(message)
					case "schedule":
						handler.Schedule(message, true)
					case "tomorrow":
						handler.Schedule(message, false)
					case "week":
						handler.Week(message, false)
					case "add":
						trackerHandler.AddRecord(message)
					default:
						tgbotapi.NewMessage(message.Chat.ID, message.Text)
					}
				} else {
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
								Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Неудалось заполнить базу данных"))
							} else {
								Bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Done!"))
							}
						}
					}
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
					switch lstQ[2] {
					case "olimp":
						trackerHandler.OlimpsCallbacksHandler(query.Message, lstQ[3], lstQ[4], lstQ[5], lstQ[6], lstQ[1])
					case "sub":
						trackerHandler.SubjectsCallbacksHandler(query.Message, lstQ[1], lstQ[3])
					}
				} else if lstQ[0] == "main" {
					if lstQ[1] == "choice" {
						bot.ChoiceBotHandler(query.Message, lstQ[2])
					}
				}
				Bot.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
			}
		}()
	}
}
