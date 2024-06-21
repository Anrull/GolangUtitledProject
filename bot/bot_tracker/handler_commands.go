package bot_tracker

import (
	"awesomeProject/bot/callbacks"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddRecord(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите предмет")
	msg.ReplyMarkup = callbacks.BuilderSubjectsForTreker
	Bot.Send(msg)
}
