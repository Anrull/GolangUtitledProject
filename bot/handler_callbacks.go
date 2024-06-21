package bot

import (
	"awesomeProject/bot/callbacks"
	"awesomeProject/data/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ChoiceBotHandler(message *tgbotapi.Message, value string) {
	err := db.Update(message.Chat.ID, "bot", value)
	if err != nil {
		Bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка связи с db"))
		return
	}
	var msg tgbotapi.EditMessageTextConfig
	if value == "bot-schedule" {
		msg = tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, "Выбран бот-расписание", callbacks.BuilderWhoAreYou)
	} else {
		db.Update(message.Chat.ID, "temp", "snils")
		msg = tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "Выбран бот РСОШ трекер.\nДля начала работы введите СНИЛС")
	}
	Bot.Send(msg)
}
