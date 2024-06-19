package callbacks

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var BuilderChoiceBot = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Бот-расписание", "main;choice;bot-schedule"),
		tgbotapi.NewInlineKeyboardButtonData("Бот-трекер", "main;choice;bot-treker"),
	),
)

//func init() {
//
//}
