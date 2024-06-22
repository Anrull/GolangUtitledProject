package dispatcher

import (
	"awesomeProject/bot"
	"awesomeProject/bot/callbacks"

	handler "awesomeProject/bot/bot_timetable"
	trackerHandler "awesomeProject/bot/bot_tracker"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TimetableCallbackQuery(query *tgbotapi.CallbackQuery, lstQ []string) {
	switch lstQ[1] {
	case "days":
		handler.DaysHandler(query.From.ID, lstQ[2], lstQ[3])
	case "who":
		handler.WhoAreYouHandler(query.Message.Chat.ID, query.Message.MessageID, lstQ[2])
	case "choice":
		handler.ChoiceTimetableHandler(query.Message.Chat.ID, query.Message.MessageID, lstQ[3], lstQ[2])
	}
}

func TrackerCallbackQuery(query *tgbotapi.CallbackQuery, lstQ []string) {
	switch lstQ[2] {
	case "olimp":
		trackerHandler.OlimpsCallbacksHandler(query.Message, lstQ[3], lstQ[4], lstQ[5], lstQ[6], lstQ[1])
	case "sub":
		trackerHandler.SubjectsCallbacksHandler(query.Message, lstQ[1], lstQ[3])
	case "stage":
		trackerHandler.StageCallbacksHandler(query.Message, lstQ[1], lstQ[3])
	case "teacher":
		trackerHandler.TeachersCallbacksHandler(query.Message, lstQ[1], lstQ[3])
	}
}

func MainCallbackQuery(query *tgbotapi.CallbackQuery, lstQ []string) {
	if lstQ[1] == "choice" {
		bot.ChoiceBotHandler(query.Message, lstQ[2])
	}
}

func YNCallbackQuery(query *tgbotapi.CallbackQuery, lstQ []string) {
	if lstQ[1] == "AddRecord" {
		trackerHandler.YNAddRecordHandler(query.Message, lstQ[2])
	}
}

func MenuCallbackQuery(query *tgbotapi.CallbackQuery, lstQ []string) {
	if lstQ[1] == "schedule" {
		message := query.Message
		switch lstQ[2] {
		case "Получить расписание":
			msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, "Какое именно", bot.MenuChoiceModeScheduleKeyboard)
			Bot.Send(msg)
		case "Сменить бота":
			msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, "Выберите бота", callbacks.BuilderChoiceBot)
			Bot.Send(msg)
		case "Прочее":
			msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, bot.MenuScheduleOtherBotKeyboard)
			Bot.Send(msg)
		case "Сегодня":
			handler.Schedule(message, true)
		case "Завтра":
			handler.Schedule(message, false)
		case "По дням":
			handler.Days(message)
			Bot.Request(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
		case "Назад":
			msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, bot.MenuScheduleBotKeyboard)
			Bot.Send(msg)
		case "Посмотреть неделю":
			handler.Week(message, true)
		case "Расписание звонков":
			handler.Time(message, true)
		}
	}
}
