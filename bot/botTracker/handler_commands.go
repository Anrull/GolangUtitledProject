package botTracker

import (
	"awesomeProject/bot"
	"awesomeProject/bot/callbacks"
	"awesomeProject/data/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"strings"
)

func AddRecord(message *tgbotapi.Message, q bool) {
	_ = db.AddTracker(message, "olimps", "")
	
	if !q {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите предмет")
		msg.ReplyMarkup = callbacks.BuilderSubjectsForTracker
		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID,
		message.MessageID, "Выберите предмет", callbacks.BuilderSubjectsForTracker)
	bot.Send(msg)

}

func HandlerDeleteOlimpsMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите нужные для удаления олимпиады\n\n<b><em>Удалять можно только последние олимпиды</em></b>")
	msg.ParseMode = tgbotapi.ModeHTML

	name, _ := db.GetTracker(message, "name")
	textSlice, err := db.GetTracker(message, "filter")
	if err != nil || textSlice == "" {
		return
	}
	slice := strings.Split(textSlice, ";;")

	var sub, olimp, stage, teacher string
	for _, v := range slice {
		vv := strings.Split(v, "||")
		if vv[0] == "sub" {
			sub = vv[1]
		} else if vv[0] == "olimp" {
			olimp = vv[1]
		} else if vv[0] == "stage" {
			stage = vv[1]
		} else if vv[0] == "teacher" {
			teacher = vv[1]
		}
	}

	records, err := db.GetRecords(name, sub, olimp, stage, teacher)
	if err != nil {
		return
	}

	markup := callbacks.CreateButtonsDelete(len(*records))
	if len(*records) > 100 {
		msg = tgbotapi.NewMessage(message.Chat.ID, "Выберите поменьше фильтров")
		bot.Send(msg)
		return
	}
	if len(*records) > 49 {
		msg.ReplyMarkup = markup.InlineKeyboard[:50]
		msg1 := tgbotapi.NewMessage(message.Chat.ID, "Вторая часть")
		msg1.ReplyMarkup = markup.InlineKeyboard[50:]
		bot.Send(msg)
		bot.Send(msg1)
		return
	}
	msg.ReplyMarkup = markup.InlineKeyboard
	bot.Send(msg)
}
