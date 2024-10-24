package feedback

import (
	"awesomeProject/bot"
	"awesomeProject/bot/logger"
	"awesomeProject/data/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(message *tgbotapi.Message, params ...string) {
	if params[1] == "escape" {
		bot.Bot.Send(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
		return
	}
	userName := message.Chat.UserName
	stage := params[1]
	sub, err := db.GetTempFbNameByID(params[2])
	if err != nil {
		bot.Send(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
		logger.Error("Error in getTempFbNameByID", err)
		return
	}
	date := params[3]
	role := params[4]
	if role == "-2" {
		role = "Не буду отвечать"
		bot.Send(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
		return
	} else if role == "-1" {
		role = "Не было учителя/урока"
	}
	err = db.CreateFBLessons(message.Chat.ID, userName, stage, sub, date, role)
	if err != nil {
		logger.Error("Ошибка заполнения бд (feedback)", err)
		return
	}
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID,
		"Спасибо за оставленный отзыв!", Escape)
	bot.Send(msg)
}
