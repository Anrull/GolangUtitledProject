package botSchedule

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/bot"
	"awesomeProject/bot/callbacks"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot = bot.Bot

func Start(message *tgbotapi.Message) {
	err := db.NewUser(*message)
	if err != nil {
		log.Println(err)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID,
		"Выберите бота\n\n<b><em>Бот-Расписание</em></b> позволяет просматривать расписание\n\n<b><em>Бот-Трекер</em></b> позволяет добавлять и отслеживать свой прогресс в РСОШ олимпиадах")
	msg.ReplyMarkup = callbacks.BuilderChoiceBot
	msg.ParseMode = tgbotapi.ModeHTML

	bot.Send(msg)
}

func DeleteMe(message *tgbotapi.Message) {
	err := db.DeleteTrackerUser(message)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Неудалось удалить аккаунт"))
		return
	}

	Start(message)
}

func Help(message *tgbotapi.Message) {
	bot.Send(tgbotapi.NewMessage(message.Chat.ID, lexicon.HelpMessage))
}

func Time(message *tgbotapi.Message, query bool) {
	if !query {
		msg := tgbotapi.NewMessage(message.Chat.ID, lexicon.TimetableTime)
		msg.ReplyMarkup = callbacks.BuilderTimetableEscape
		msg.ParseMode = tgbotapi.ModeHTML
		bot.Send(msg)
		return
	}

	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, lexicon.TimetableTime, callbacks.BuilderTimetableEscape)
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

func Days(message *tgbotapi.Message) {
	msg1 := tgbotapi.NewMessage(message.Chat.ID, "Расписание нечетной недели")
	msg2 := tgbotapi.NewMessage(message.Chat.ID, "Расписание четной недели")

	msg1.ReplyMarkup = callbacks.BuilderDays1
	msg2.ReplyMarkup = callbacks.BuilderDays0

	bot.Send(msg1)
	bot.Send(msg2)
}

func Schedule(message *tgbotapi.Message, today bool) {
	var week string
	role, err := db.Get(message.Chat.ID, "role")
	logging(message, err)

	if today {
		week, err = timetable.GetWeek(false, true)
	} else {
		week, err = timetable.GetWeek(true, true)
	}

	if err != nil {
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Кажется, идут каникулы :)"))
	} else {
		var day string
		var photoByte []byte

		if today {
			day = timetable.GetDayToday()
		} else {
			day = timetable.GetDayTomorrow()
		}

		colors, err := db.GetColorByUserID(message.Chat.ID)
		logging(message, err)

		if role == "student" {
			stage, err := db.Get(message.Chat.ID, "classes")
			logging(message, err)

			lessons, err := timetable.GetTimetableText(week, day, stage)
			logging(message, err)

			extraLessons, err := timetable.GetExtraTimetableText(week, day, stage)
			logging(message, err)

			lessons = timetable.Merge(lessons, extraLessons)

			photoByte, _ = timetable.DrawTimetable(
				lessons, fmt.Sprintf("%s, нед: %s, день: %s",
					stage, lexicon.Week[week], lexicon.Day[day]),
				false, colors...)

			// if !reflect.DeepEqual(extraLessons, [][]string{{}, {}, {}}) {
			// 	extraPhotoByte, _ := timetable.DrawTimetable(
			// 		extraLessons, "Внеурочные занятия",
			// 		false, colors...)
				
			// 	defer SendPhotoByte(message.Chat.ID, extraPhotoByte)
			// }
		} else {
			teacher, err := db.Get(message.Chat.ID, "name_teacher")
			logging(message, err)

			lessons, err := timetable.GetTimetableTeachersText(teacher, week, day)
			logging(message, err)

			photoByte, _ = timetable.DrawTimetable(
				lessons, fmt.Sprintf("%s, нед: %s, день: %s", teacher, lexicon.Week[week], lexicon.Day[day]),
				true, colors...)
		}

		SendPhotoByte(message.Chat.ID, photoByte)
	}
}

func Week(message *tgbotapi.Message, query bool) {
	week, err := timetable.GetWeek(false, false)
	logging(message, err)

	if week == "н" {
		week = "нечетная"
	} else {
		week = "четная"
	}

	text := fmt.Sprintf("Текущая неделя - <b><em>%s</em></b>", week)

	if !query {
		msg := tgbotapi.NewMessage(message.Chat.ID, text)
		msg.ParseMode = tgbotapi.ModeHTML
		msg.ReplyMarkup = callbacks.BuilderTimetableEscape
		bot.Send(msg)
		return
	}

	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, callbacks.BuilderTimetableEscape)
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

func SendPhotoByte(ChatID int64, photoBytes []byte) {
	_, err := Bot.Send(tgbotapi.NewPhoto(ChatID, tgbotapi.FileBytes{
		Name:  "photo.jpg",
		Bytes: photoBytes,
	}))
	if err != nil {
		log.Println(err)
	}
}

func logging(message *tgbotapi.Message, err error) {
	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка связи с db"))
	}
}
