package main

import (
	"awesomeProject/bot"
	"awesomeProject/bot/dispatcher"

	handler "awesomeProject/bot/botSchedule"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"log"
)

var Bot = bot.Bot

func main() {
	log.Printf("Authorized on account %s", Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	//fmt.Println(u.Timeout)
	//u.Timeout = 1

	go handler.RunScheduler()

	updates := Bot.GetUpdatesChan(u)

	for update := range updates {
		go dispatcher.Dispatcher(&update)
	}
}
