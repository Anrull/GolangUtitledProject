package dispatcher

import (
	"awesomeProject/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"strings"
)

var Bot = bot.Bot

func Dispatcher(update *tgbotapi.Update) {
	if update.Message != nil {
		message := update.Message
		if message.IsCommand() {
			CommandsHandling(message)
		} else {
			MessageHandler(message)
		}
	} else if update.CallbackQuery != nil {
		query := update.CallbackQuery
		lstQ := strings.Split(query.Data, ";")
		if lstQ[0] == "timetable" {
			TimetableCallbackQuery(query, lstQ)
		} else if lstQ[0] == "tracker" {
			TrackerCallbackQuery(query, lstQ)
		} else if lstQ[0] == "main" {
			MainCallbackQuery(query, lstQ)
		} else if lstQ[0] == "yn" {
			YNCallbackQuery(query, lstQ)
		} else if lstQ[0] == "menu" {
			MenuCallbackQuery(query, lstQ)
		}
		Bot.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
	}
}
