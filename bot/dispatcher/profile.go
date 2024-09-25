package dispatcher

import (
	"awesomeProject/bot"
	"awesomeProject/data/db"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func Profile(message *tgbotapi.Message) {
	user, err := db.GetInfoAboutPerson(message.Chat.ID)
	if err != nil {
		log.Println(1, err)
		return
	}

	builder := &strings.Builder{}
	builder.WriteString("<b>")
	builder.WriteString(message.Chat.UserName)
	builder.WriteString("</b>\nБот: <em>")
	builder.WriteString(user.Bot)
	var text string
	if user.Newsletter == "1" {
		text = "</em>\nУведомления <em>включены</em>\n\n"
	} else {
		text = "</em>\nУведомления <em>отлючены</em>\n\n"
	}
	builder.WriteString(text)
	builder.WriteString("<b><em>Бот-расписание:</em></b>\n")
	if user.Role == "student" {
		builder.WriteString("Класс: ")
		builder.WriteString(user.Classes)
	} else {
		builder.WriteString("Имя учителя: ")
		builder.WriteString(user.NameTeacher)
	}
	builder.WriteString("\n\n<b><em>РСОШ.Трекер</em></b>\n")
	if _, err = db.GetTracker(message, "name"); err != nil {
		builder.WriteString("<em>Вы не зарегестрированны в РСОШ.Трекере</em>")
	} else {
		tracker, err := db.GetInfoAboutPersonTracker(message.Chat.ID)
		if err != nil {
			log.Println(2, err)
			return
		}

		builder.WriteString("Имя: <em>")
		builder.WriteString(tracker.Name)
		builder.WriteString("</em>\nКласс: ")
		builder.WriteString(tracker.Stage)
		builder.WriteString("\nКол-во записей: <b><em>")

		num, err := db.GetRecordsCount(tracker.Name, "nil", "nil", "nil", "nil")
		if err != nil {
			log.Println(3, err)
			return
		}

		builder.WriteString(fmt.Sprintf("%d</em></b>", num))
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, builder.String())
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}
