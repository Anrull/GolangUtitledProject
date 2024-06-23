package bot_tracker

import (
	"awesomeProject/bot/callbacks"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddRecord(message *tgbotapi.Message, q bool) {
	if !q {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите предмет")
		msg.ReplyMarkup = callbacks.BuilderSubjectsForTracker
		Bot.Send(msg)
		return
	}
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, "Выберите предмет", callbacks.BuilderSubjectsForTracker)
	Bot.Send(msg)

}
