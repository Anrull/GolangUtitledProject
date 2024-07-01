package dispatcher

import (
	"awesomeProject/bot"
	"awesomeProject/bot/callbacks"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"
	"fmt"

	handler "awesomeProject/bot/botSchedule"
	trackerHandler "awesomeProject/bot/botTracker"

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
	message := query.Message
	if lstQ[1] == "schedule" {
		switch lstQ[2] {
		case "Получить расписание":
			msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, "Какое именно", bot.MenuChoiceModeScheduleKeyboard)
			bot.Send(msg)
		case "Сменить бота":
			msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, "Выберите бота", callbacks.BuilderChoiceBot)
			bot.Send(msg)
		case "Прочее":
			msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, bot.MenuScheduleOtherBotKeyboard)
			bot.Send(msg)
		case "Сегодня":
			handler.Schedule(message, true)
		case "Завтра":
			handler.Schedule(message, false)
		case "По дням":
			handler.Days(message)
			bot.Request(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
		case "Назад":
			msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, bot.MenuScheduleBotKeyboard)
			bot.Send(msg)
		case "Посмотреть неделю":
			handler.Week(message, true)
		case "Расписание звонков":
			handler.Time(message, true)
		}
	} else if lstQ[1] == "filter" {
		switch lstQ[2] {
		case "Без фильтров":
			//trackerHandler.AddRecord(query.Message, true)
			_ = db.AddTracker(message, "filter", "")
			bot.Send(
				tgbotapi.NewEditMessageTextAndMarkup(
					message.Chat.ID, message.MessageID,
					"Выберите фильтр", callbacks.SomeGetSubjectsTracker))
		case "Несколько фильтров":
			_ = db.AddTracker(message, "filter", "")
			bot.Send(
				tgbotapi.NewEditMessageTextAndMarkup(
					message.Chat.ID, message.MessageID,
					"Выберите фильтр", callbacks.SomeGetSubjectsTracker))
		case "Отфильтровать по олимпиаде":
			will := bot.CopyInlineKeyboard(callbacks.BuilderGetOlimpsKeyboard)
			will.InlineKeyboard = will.InlineKeyboard[:lexicon.OlimpListStep]
			will.InlineKeyboard = append(will.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					lexicon.OlimpListLeft, fmt.Sprintf("tracker;get;olimp;nil;0;%d;min",
						len(callbacks.BuilderGetOlimpsKeyboard.InlineKeyboard))),
				tgbotapi.NewInlineKeyboardButtonData(
					lexicon.OlimpListRight, fmt.Sprintf("tracker;get;olimp;nil;0;%d;plus",
						len(callbacks.BuilderGetOlimpsKeyboard.InlineKeyboard)))))
			bot.Send(tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID,
				query.Message.MessageID, "Выберите нужную олимпиаду", will))
		case "Отфильтровать по предмету":
			bot.Send(
				tgbotapi.NewEditMessageTextAndMarkup(
					query.Message.Chat.ID, query.Message.MessageID,
					"Выберите нужную олимпиаду", callbacks.SortBuilderSubjectsKeyboard))
		case "Отфильтровать по этапу":
			bot.Send(
				tgbotapi.NewEditMessageTextAndMarkup(
					query.Message.Chat.ID, query.Message.MessageID,
					"Выберите нужную олимпиаду", callbacks.SortBuilderStageKeyboard))
		case "Отфильтровать по наставнику":
			bot.Send(
				tgbotapi.NewEditMessageTextAndMarkup(
					query.Message.Chat.ID, query.Message.MessageID,
					"Выберите нужную олимпиаду", callbacks.BuilderGetTeacherKeyboard))
		case "Назад":
			bot.Send(tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID,
				message.MessageID, "Вот некоторые опции",
				bot.BuilderMenuTracker))
		case "Удалить":
			trackerHandler.HandlerDeleteOlimpsMessage(message)
		}
	} else if lstQ[1] == "tracker" {
		switch lstQ[2] {
		case "Добавить запись":
			trackerHandler.AddRecord(message, true)
		case "Просмотр записей":
			bot.Send(
				tgbotapi.NewEditMessageTextAndMarkup(
					message.Chat.ID, message.MessageID,
					"Выберите фильтр", bot.BuilderChoiceTrackerFilter))
		case "Назад":
			bot.Send(
				tgbotapi.NewEditMessageTextAndMarkup(
					message.Chat.ID, message.MessageID,
					"Выберите бота", callbacks.BuilderChoiceBot))
		}
	}
}
