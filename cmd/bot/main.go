package main

import (
	"awesomeProject/bot"
	"awesomeProject/bot/dispatcher"
	"awesomeProject/bot/logger"
	"log"

	handler "awesomeProject/bot/botSchedule"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/manucorporat/try"

	"fmt"
	"log/slog"
)

var Bot = bot.Bot

func main() {
	log.Printf("Authorized on account %s", Bot.Self.UserName)
	logger.New(slog.LevelInfo, "data/logs/bot.log")
	
	logger.Info(fmt.Sprintf("Authorized on account %s", Bot.Self.UserName))
	logger.Info("Logger started")

	go handler.RunScheduler()

	u := tgbotapi.NewUpdate(0)
	updates := Bot.GetUpdatesChan(u)

	for update := range updates {
		try.This(func() {
			go dispatcher.Dispatcher(&update)
		}).Catch(func(e try.E) {
			fmt.Println("Caught an error:", e)
		})
	}
}
