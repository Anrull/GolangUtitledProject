package bot_timetable

import (
	"awesomeProject/bot"
	"awesomeProject/bot/callbacks"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var Bot = bot.Bot

func Start(message *tgbotapi.Message) {
	err := db.NewUser(*message)
	if err != nil {
		log.Println(err)
		return
	}
	Bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Добро пожаловать в новую версию бота, написанную на языке программирования Golang!"))
}

func Help(message *tgbotapi.Message) {
	Bot.Send(tgbotapi.NewMessage(message.Chat.ID, lexicon.HelpMessage))
}

func Time(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, lexicon.TimetableTime)
	msg.ParseMode = tgbotapi.ModeHTML
	Bot.Send(msg)
}

func Days(message *tgbotapi.Message) {
	msg1 := tgbotapi.NewMessage(message.Chat.ID, "Расписание нечетной недели")
	msg2 := tgbotapi.NewMessage(message.Chat.ID, "Расписание четной недели")
	msg1.ReplyMarkup = callbacks.BuilderDays1
	msg2.ReplyMarkup = callbacks.BuilderDays0
	Bot.Send(msg1)
	Bot.Send(msg2)
}
