package feedback

import (
	"awesomeProject/bot"
	"awesomeProject/data/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Handler(message *tgbotapi.Message, params ...string) {
	if params[1] == "escape" {
		bot.Bot.Send(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
		return
	}
	userName := message.Chat.UserName
	stage := params[1]
	sub := params[2]
	date := params[3]
	role := params[4]
	if role == "-2" {
		role = "Не буду отвечать"
		bot.Send(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
		return
	} else if role == "-1" {
		role = "Не было учителя/урока"
	}
	err := db.CreateFBLessons(message.Chat.ID, userName, stage, sub, date, role)
	if err != nil {
		log.Println("Ошибка заполнения бд (feedback)", err)
		return
	}
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID,
		"Спасибо за оставленный отзыв!", Escape)
	bot.Send(msg)
}
