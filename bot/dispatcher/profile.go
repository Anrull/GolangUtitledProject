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
	var user db.User
	var err error
	var username string
	if isAdmin(message) {
		lst := strings.Split(message.Text, " ")
		if len(lst) > 1 {
			username = lst[1]
			id, err := db.GetChatID(username)
			if err != nil {
				log.Println(err)
				bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Пользователь не найден"))
			}

			user, err = db.GetInfoAboutPerson(id)
			if err != nil {
				log.Println(err)
				bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Пользователь не найден"))
				return
			}
		}
	} else {
		user, err = db.GetInfoAboutPerson(message.Chat.ID)
		if err != nil {
			log.Println(1, err)
			return
		}
		username = message.From.UserName
	}


	builder, err := prepareProfileText(user, username, message.Chat.ID)
	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка подготовки данных"))
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, builder)
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

func prepareProfileText(user db.User, username string, id int64) (string, error) {
	builder := &strings.Builder{}
	builder.WriteString("<b>")
	builder.WriteString(username)
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
	if _, err := db.GetTrackerById(id, "name"); err != nil {
		builder.WriteString("<em>Вы не зарегестрированны в РСОШ.Трекере</em>")
	} else {
		tracker, err := db.GetInfoAboutPersonTracker(id)
		if err != nil {
			log.Println(2, err)
			return "", err
		}

		builder.WriteString("Имя: <em>")
		builder.WriteString(tracker.Name)
		builder.WriteString("</em>\nКласс: ")
		builder.WriteString(tracker.Stage)
		builder.WriteString("\nКол-во записей: <b><em>")

		num, err := db.GetRecordsCount(tracker.Name, "nil", "nil", "nil", "nil")
		if err != nil {
			log.Println(3, err)
			return "", err
		}

		builder.WriteString(fmt.Sprintf("%d</em></b>", num))
	}

	return builder.String(), nil
}
