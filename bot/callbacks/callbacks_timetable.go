package callbacks

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/bot/lexicon"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var BuilderDays0 = tgbotapi.NewInlineKeyboardMarkup() // BuilderDays0 расписание по дням
var BuilderDays1 = tgbotapi.NewInlineKeyboardMarkup() // BuilderDays1 расписание по дням

// BuilderWhoAreYou Выбор между учителем и учеником
var BuilderWhoAreYou = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Учитель", "timetable;who;teacher"),
		tgbotapi.NewInlineKeyboardButtonData("Ученик", "timetable;who;student"),
	),
)

var BuilderChoiceStage = tgbotapi.NewInlineKeyboardMarkup()   // BuilderChoiceStage Выбор класса для расписания
var BuilderChoiceTeacher = tgbotapi.NewInlineKeyboardMarkup() // BuilderChoiceTeacher Выбор учителя для расписания

func init() {
	slice0 := []tgbotapi.InlineKeyboardButton{}
	slice1 := []tgbotapi.InlineKeyboardButton{}
	for i := range lexicon.ListDays {
		slice0 = append(slice0,
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.ListDays[i],
				fmt.Sprintf("timetable;days;%d;%d", 0, i)))
		slice1 = append(slice1,
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.ListDays[i],
				fmt.Sprintf("timetable;days;%d;%d", 1, i)))
	}
	BuilderDays0 = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(slice0...))
	BuilderDays1 = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(slice1...))

	slice2 := []tgbotapi.InlineKeyboardButton{}
	for i := range lexicon.Stages {
		slice2 = append(slice2,
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.Stages[i],
				fmt.Sprintf("timetable;choice;student;%s", lexicon.Stages[i])))
	}
	BuilderChoiceStage = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(slice2[:4]...),
		tgbotapi.NewInlineKeyboardRow(slice2[4:8]...),
		tgbotapi.NewInlineKeyboardRow(slice2[8:12]...),
		tgbotapi.NewInlineKeyboardRow(slice2[12:]...),
	)

	slice2 = []tgbotapi.InlineKeyboardButton{}
	for i := range timetable.Teachers {
		slice2 = append(slice2,
			tgbotapi.NewInlineKeyboardButtonData(
				timetable.Teachers[i],
				fmt.Sprintf("timetable;choice;teacher;%s", timetable.Teachers[i])))
	}
	BuilderChoiceTeacher = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(slice2[:3]...),
		tgbotapi.NewInlineKeyboardRow(slice2[3:6]...),
		tgbotapi.NewInlineKeyboardRow(slice2[6:9]...),
		tgbotapi.NewInlineKeyboardRow(slice2[9:12]...),
		tgbotapi.NewInlineKeyboardRow(slice2[12:15]...),
		tgbotapi.NewInlineKeyboardRow(slice2[15:18]...),
		tgbotapi.NewInlineKeyboardRow(slice2[18:21]...),
		tgbotapi.NewInlineKeyboardRow(slice2[21:24]...),
		tgbotapi.NewInlineKeyboardRow(slice2[24:27]...),
		tgbotapi.NewInlineKeyboardRow(slice2[27:30]...),
		tgbotapi.NewInlineKeyboardRow(slice2[30:]...),
	)
}
