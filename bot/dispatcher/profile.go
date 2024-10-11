package dispatcher

import (
	"awesomeProject/bot"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Profile(message *tgbotapi.Message, q ...bool) {
	var user db.User
	var err error
	var username string
	flag := true
	if isAdmin(message) {
		lst := strings.Split(message.Text, " ")
		if len(lst) > 1 && len(q) == 0 {
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
			flag = false
		} else {
			user, err = db.GetInfoAboutPerson(message.Chat.ID)
			if err != nil {
				log.Println(1, err)
				return
			}
			username = message.Chat.UserName
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

	if len(q) == 0 {
		msg := tgbotapi.NewMessage(message.Chat.ID, builder)
		msg.ParseMode = tgbotapi.ModeHTML
		if flag {
			msg.ReplyMarkup = bot.ProfilePanel
		}

		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, builder, bot.ProfilePanel)
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
	builder.WriteString("<b><em>Бот-расписание:</em></b>\nПалитра: ")
	builder.WriteString(lexicon.RgbToColorsConfig[user.Color])
	builder.WriteString("\n")
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

func HandlerProfileCallbacks(query *tgbotapi.CallbackQuery) {
	lstQ := strings.Split(query.Data, ";")
	switch lstQ[1] {
	case "escape":
		Profile(query.Message, true)
	case "choice-color":
		if len(lstQ) < 3 {
			return
		}
		switch lstQ[2] {
		case "main":
			msg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, "Выберите цвет палитры", bot.ProfileColorsPanel)
			bot.Send(msg)
			return
		case "white":
			err := db.Update(query.Message.Chat.ID, "color", lexicon.ColorsToRgbConfig["white"])
			if err != nil {
				log.Println(err)
			}
		case "green":
			err := db.Update(query.Message.Chat.ID, "color", lexicon.ColorsToRgbConfig["green"])
			if err != nil {
				log.Println(err)
			}
		case "blue":
			err := db.Update(query.Message.Chat.ID, "color", lexicon.ColorsToRgbConfig["blue"])
			if err != nil {
				log.Println(err)
			}
		case "yellow":
			err := db.Update(query.Message.Chat.ID, "color", lexicon.ColorsToRgbConfig["yellow"])
			if err != nil {
				log.Println(err)
			}
		case "purple":
			err := db.Update(query.Message.Chat.ID, "color", lexicon.ColorsToRgbConfig["purple"])
			if err != nil {
				log.Println(err)
			}
		}
		Profile(query.Message, true)
	}

}