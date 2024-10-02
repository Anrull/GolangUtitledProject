package bot

import (
	"awesomeProject/bot/callbacks"
	"awesomeProject/data/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ChoiceBotHandler(message *tgbotapi.Message, value string) {
	err := db.Update(message.Chat.ID, "bot", value)
	if err != nil {
		Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка связи с db"))
		return
	}

	var msg tgbotapi.EditMessageTextConfig
	if value == "bot-schedule" {
		msg = tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID,
			message.MessageID, "Выбран бот-расписание", callbacks.BuilderWhoAreYou)
	} else {
		_, err = db.GetTracker(message, "name")
		if err != nil {
			_ = db.Update(message.Chat.ID, "temp", "snils")
		}
		_, err = db.GetTracker(message, "name")
		if err != nil {
			msg = tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID,
				"Выбран бот РСОШ трекер.\nДля начала работы введите СНИЛС", tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Отмена", "menu;tracker;snils"),
					),
				))
		} else {
			msg = tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID,
				message.MessageID, "Выбран бот РСОШ трекер.", BuilderMenuTracker)
		}
	}
	Send(msg)
}
