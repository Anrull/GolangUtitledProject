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
		_, err := db.GetTracker(message, "name")
		if err != nil {
			msg = tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "Выбран бот РСОШ трекер.\nДля начала работы введите СНИЛС")
		} else {
			msg = tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "Выбран бот РСОШ трекер.")
		}
	}
	Bot.Send(msg)
}

//func HandlerMenuButtons(message *tgbotapi.Message, role string) {
//switch role {
//case "Получить расписание":
//	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, "Какое именно", MenuChoiceModeScheduleKeyboard)
//	Bot.Send(msg)
//case "Сменить бота":
//	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, "Выберите бота", callbacks.BuilderChoiceBot)
//	Bot.Send(msg)
//case "Прочее":
//	msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, MenuScheduleOtherBotKeyboard)
//	Bot.Send(msg)
//case "Сегодня":
//	bot_timetable.Schedule(message, true)
//case "Завтра":
//	bot_timetable.Schedule(message, false)
//case "По дням":
//	bot_timetable.Days(message)
//	Bot.Request(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
//case "Назад":
//	msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, MenuScheduleBotKeyboard)
//	Bot.Send(msg)
//case "Посмотреть неделю":
//	bot_timetable.Week(message, true)
//case "Расписание звонков":
//	bot_timetable.Time(message, true)
//}
//}
