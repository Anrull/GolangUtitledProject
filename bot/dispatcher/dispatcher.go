package dispatcher

import (
	"awesomeProject/bot"
	"awesomeProject/bot/feedback"

	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Dispatcher(update *tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.Document != nil {
			message := update.Message
			if !isAdmin(message) { return }

			FileHandler(message)
			return
		}
		if !bot.TechnicalWork {
			message := update.Message
			if message.IsCommand() {
				CommandsHandling(message)
			} else {
				MessageHandler(message)
			}
			return
		}
		if update.Message.IsCommand() {
			if update.Message.Command() == "unlock" {
				if isAdmin(update.Message) {
					bot.TechnicalWork = false
					return
				}
			}
		}
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Идут тех работы"))
	}
	if update.CallbackQuery != nil {
		if bot.TechnicalWork {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Идут тех работы"))
			return
		}
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
		} else if lstQ[0] == "admin" {
			AdminPanelHandler(query, lstQ[1], lstQ...)
		} else if lstQ[0] == "lesson" {
			feedback.Handler(query.Message, lstQ...)
		} else if lstQ[0] == "profile" {
			HandlerProfileCallbacks(query)
		}

		bot.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
	}
}
