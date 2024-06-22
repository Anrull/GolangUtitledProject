package bot_timetable

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/bot"
	"awesomeProject/bot/callbacks"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var Bot = bot.Bot

func Start(message *tgbotapi.Message) {
	err := db.NewUser(*message)
	if err != nil {
		log.Println(err)
		return
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "Добро пожаловать в новую версию бота, написанную на языке программирования Golang!")
	//msg.ReplyMarkup = callbacks.BuilderWhoAreYou
	//will := tgbotapi.InlineKeyboardMarkup{}
	////will.InlineKeyboard = callbacks.BuilderOlimpsKeyboard.InlineKeyboard
	//willCopy_ := make([][]tgbotapi.InlineKeyboardButton, len(callbacks.BuilderOlimpsKeyboard.InlineKeyboard))
	//copy(willCopy_, callbacks.BuilderOlimpsKeyboard.InlineKeyboard)
	//will.InlineKeyboard = willCopy_[:lexicon.OlimpListStep]
	////will.InlineKeyboard = callbacks.BuilderOlimpsKeyboard.InlineKeyboard[:lexicon.OlimpListStep]
	//will.InlineKeyboard = append(will.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
	//	tgbotapi.NewInlineKeyboardButtonData(
	//		lexicon.OlimpListLeft, fmt.Sprintf("tracker;add;olimp;nil;0;%d;min", len(callbacks.BuilderOlimpsKeyboard.InlineKeyboard))),
	//	tgbotapi.NewInlineKeyboardButtonData(
	//		lexicon.OlimpListRight, fmt.Sprintf("tracker;add;olimp;nil;0;%d;plus", len(callbacks.BuilderOlimpsKeyboard.InlineKeyboard)))))
	//msg.ReplyMarkup = will
	msg.ReplyMarkup = callbacks.BuilderChoiceBot
	Bot.Send(msg)
}

func Help(message *tgbotapi.Message) {
	Bot.Send(tgbotapi.NewMessage(message.Chat.ID, lexicon.HelpMessage))
}

func Time(message *tgbotapi.Message, q bool) {
	if !q {
		msg := tgbotapi.NewMessage(message.Chat.ID, lexicon.TimetableTime)
		msg.ParseMode = tgbotapi.ModeHTML
		Bot.Send(msg)
		return
	}
	msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, lexicon.TimetableTime)
	msg.ParseMode = tgbotapi.ModeHTML
	Bot.Send(msg)
}

func Days(message *tgbotapi.Message) {
	msg1 := tgbotapi.NewMessage(message.Chat.ID, "Расписание нечетной недели")
	msg2 := tgbotapi.NewMessage(message.Chat.ID, "Расписание четной недели")
	msg1.ReplyMarkup = callbacks.BuilderDays1
	msg2.ReplyMarkup = callbacks.BuilderDays0
	Bot.Send(msg1)
	Bot.Send(msg2)
}

func Schedule(message *tgbotapi.Message, today bool) {
	role, err := db.Get(message.Chat.ID, "role")
	logging(message, err)
	week, err := timetable.GetWeek(false, true)
	if err != nil {
		Bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Кажется, идут каникулы :)"))
	} else {
		n, err := count("data/temp/images")
		logging(message, err)
		filename, day := fmt.Sprintf("data/temp/images/schedule%d.png", n), ""
		if today {
			day = timetable.GetDayToday()
		} else {
			day = timetable.GetDayTomorrow()
		}
		if role == "student" {
			stage, err := db.Get(message.Chat.ID, "classes")
			logging(message, err)
			lessons, err := timetable.GetTimetableText(week, day, stage)
			logging(message, err)
			timetable.DrawTimetable(
				lessons, fmt.Sprintf("%s, нед: %s, день: %s", stage, lexicon.Week[week], lexicon.Day[day]),
				false, n)
		} else {
			teacher, err := db.Get(message.Chat.ID, "name_teacher")
			logging(message, err)
			lessons, err := timetable.GetTimetableTeachersText(teacher, week, day)
			logging(message, err)
			timetable.DrawTimetable(
				lessons, fmt.Sprintf("%s, нед: %s, день: %s", teacher, lexicon.Week[week], lexicon.Day[day]),
				true, n)
		}
		sendPhoto(message.Chat.ID, filename)
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
		Bot.Send(msg)
	} else {
		msg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, text)
		msg.ParseMode = tgbotapi.ModeHTML
		Bot.Send(msg)
	}
}

func sendPhoto(ChatID int64, filename string) {
	photoBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	Bot.Send(tgbotapi.NewPhoto(ChatID, tgbotapi.FileBytes{
		Name:  "photo.jpg",
		Bytes: photoBytes,
	}))
}

func logging(message *tgbotapi.Message, err error) {
	if err != nil {
		log.Println(err)
		Bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка связи с db"))
	}
}
