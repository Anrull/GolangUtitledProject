package bot_tracker

import (
	"awesomeProject/data/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func YNAddRecordHandler(message *tgbotapi.Message, role string) {
	var msg tgbotapi.EditMessageTextConfig
	if role == "yes" {
		textSlice, err := db.GetTracker(message, "olimps")
		if err != nil || textSlice == "" {
			return
		}
		slice := strings.Split(textSlice, ";")
		name, err := db.GetTracker(message, "name")
		logging(message, err)
		class, err := db.GetTracker(message, "stage")
		logging(message, err)
		err = db.AddRecord(name, class, slice[1], slice[0], slice[3], slice[2])
		if err != nil {
			logging(message, err)
			return
		}
		msg = tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "Запись добавлена")
	} else {
		msg = tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "Запись не была добавлена")
	}
	Bot.Send(msg)
	db.AddTracker(message, "olimps", "")
}
