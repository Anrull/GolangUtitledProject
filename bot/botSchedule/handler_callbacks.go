package botSchedule

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/bot"
	"awesomeProject/bot/callbacks"
	"awesomeProject/data/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"fmt"
	"log"
)

var days = map[string]string{"0": "пн", "1": "вт", "2": "ср", "3": "чт", "4": "пт"}
var weeks = map[string]string{"0": "чет", "1": "нечет"}

func WhoAreYouHandler(ChatId int64, msgId int, role string) {
	err := db.Update(ChatId, "role", role)
	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(ChatId, "Ошибка связи с db"))
		return
	}

	var msg tgbotapi.EditMessageTextConfig
	if role == "student" {
		msg = tgbotapi.NewEditMessageTextAndMarkup(ChatId, msgId, "Кто именно?", callbacks.BuilderChoiceStage)
	} else if role == "teacher" {
		msg = tgbotapi.NewEditMessageTextAndMarkup(ChatId, msgId,
			"Кто именно?", callbacks.BuilderChoiceTeacher)
	}

	bot.Send(msg)
}

func DaysHandler(ChatID int64, week, day string) {
	role, err := db.Get(ChatID, "role")
	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка свзяи с db"))
		return
	}

	var schedule [][]string
	var photoByte []byte
	colors, err := db.GetColorByUserID(ChatID)

	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка свзяи с db"))
		return
	}
	if role == "student" {
		res, err := db.Get(ChatID, "classes")
		if err != nil {
			log.Println(err)
			bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка связи с db"))
			return
		}

		schedule, err = timetable.GetTimetableText(week, day, res)
		if err != nil {
			log.Println(err)
			bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка парсинга расписания"))
			return
		}

		extraLessons, err := timetable.GetExtraTimetableText(week, day, res)
		if err != nil {
			log.Println(err)
			bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка парсинга внеурочек"))
			return
		}

		schedule = timetable.Merge(schedule, extraLessons)

		photoByte, err = timetable.DrawTimetable(schedule,
			fmt.Sprintf("%s, нед: %s, день: %s", res, weeks[week],
				days[day]), false, colors...)

		if err != nil {
			log.Println(err)
			bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка отрисовки расписания"))
			return
		}
	} else {
		name, err := db.Get(ChatID, "name_teacher")
		if err != nil {
			log.Println(err)
			bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка свзяи с db"))
			return
		}

		schedule, err = timetable.GetTimetableTeachersText(name, week, day)
		if err != nil {
			log.Println(err)
			bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка отрисовки расписания"))
			return
		}

		photoByte, _ = timetable.DrawTimetable(schedule,
			fmt.Sprintf("%s, нед: %s, день: %s",
				name, weeks[week], days[day]), true, colors...)
	}

	SendPhotoByte(ChatID, photoByte)
}

func ChoiceTimetableHandler(ChatId int64, msgId int, param, role string) {
	var msg = tgbotapi.EditMessageTextConfig{}

	var err error
	if role == "student" {
		err = db.Update(ChatId, "classes", param)
	} else {
		err = db.Update(ChatId, "name_teacher", param)
	}

	if err != nil {
		bot.Send(tgbotapi.NewMessage(ChatId, "Ошибка связи с db"))
		log.Println(err)
		return
	}

	if role == "student" {
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			ChatId, msgId, fmt.Sprintf("Принято, ученик %s класса", param), bot.MenuScheduleBotKeyboard)
	} else {
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			ChatId, msgId, fmt.Sprintf("Принято, %s", param), bot.MenuScheduleBotKeyboard)
	}

	bot.Send(msg)
}
