package bot_timetable

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/bot/callbacks"
	"awesomeProject/data/db"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var days = map[string]string{"0": "пн", "1": "вт", "2": "ср", "3": "чт", "4": "пт"}
var weeks = map[string]string{"0": "чет", "1": "нечет"}

func WhoAreYouHandler(ChatId int64, msgId int, role string) {
	err := db.Update(ChatId, "role", role)
	if err != nil {
		log.Println(err)
		Bot.Send(tgbotapi.NewMessage(ChatId, "Ошибка связи с db"))
		return
	}
	var msg tgbotapi.EditMessageTextConfig
	if role == "student" {
		//tgbotapi.NewEditMessageText(ChatId, msgId, "Выберите класс")
		//msg.ReplyMarkup = callbacks.BuilderChoiceStage
		msg = tgbotapi.NewEditMessageTextAndMarkup(ChatId, msgId, "Кто именно?", callbacks.BuilderChoiceStage)
	} else if role == "teacher" {
		//tgbotapi.NewEditMessageText(ChatId, msgId, "Кто именно?")
		msg = tgbotapi.NewEditMessageTextAndMarkup(ChatId, msgId, "Кто именно?", callbacks.BuilderChoiceTeacher)
		//msg.ReplyMarkup = callbacks.BuilderChoiceTeacher
	}
	Bot.Send(msg)
}

func DaysHandler(ChatID int64, week, day string) {
	role, err := db.Get(ChatID, "role")
	if err != nil {
		log.Println(err)
		Bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка свзяи с db"))
		return
	}

	var schedule [][]string
	n, err := count("data/temp/images")
	filename := fmt.Sprintf("data/temp/images/schedule%d.png", n)
	if role == "student" {
		res, err := db.Get(ChatID, "classes")
		if err != nil {
			log.Println(err)
			Bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка свзяи с db"))
			return
		}
		schedule, err = timetable.GetTimetableText(week, day, res)
		timetable.DrawTimetable(schedule,
			fmt.Sprintf("%s, нед: %s, день: %s", res, weeks[week], days[day]), false, n)
	} else {
		name, err := db.Get(ChatID, "name_teacher")
		if err != nil {
			log.Println(err)
			Bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка свзяи с db"))
			return
		}
		schedule, err = timetable.GetTimetableTeachersText(name, week, day)
		timetable.DrawTimetable(schedule,
			fmt.Sprintf("%s, нед: %s, день: %s", name, weeks[week], days[day]), true, n)
	}
	sendPhoto(ChatID, filename)
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
		Bot.Send(tgbotapi.NewMessage(ChatId, "Ошибка связи с db"))
		log.Println(err)
		return
	}

	if role == "student" {
		msg = tgbotapi.NewEditMessageText(
			ChatId, msgId, fmt.Sprintf("Принято, ученик %s класса", param))
	} else {
		msg = tgbotapi.NewEditMessageText(
			ChatId, msgId, fmt.Sprintf("Принято, %s", param))
	}
	Bot.Send(msg)
}

func count(path string) (int, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("не удалось прочитать директорию: %w", err)
	}

	count_ := 0
	for _, file := range files {
		if !file.IsDir() {
			count_++
		}
	}

	return count_, nil
}
