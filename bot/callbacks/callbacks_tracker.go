package callbacks

import (
	"awesomeProject/bot/lexicon"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// создание клавиатуры выбора предмета "tracker;add;sub;%d"

var BuilderSubjectsForTreker tgbotapi.InlineKeyboardMarkup
var SortBuilderSubjectsKeyboard tgbotapi.InlineKeyboardMarkup
var BuilderDeleteSubjectsForTreker tgbotapi.InlineKeyboardMarkup

// создание клавиатуры выбора по этапу

var BuilderDeleteStageKeyboards tgbotapi.InlineKeyboardMarkup
var BuilderStageKeyboards tgbotapi.InlineKeyboardMarkup
var SortBuilderStageKeyboard tgbotapi.InlineKeyboardMarkup

// создание клавиатуры выбора наставника

var BuilderTeacherKeyboards tgbotapi.InlineKeyboardMarkup
var BuilderGetTeacherKeyboard tgbotapi.InlineKeyboardMarkup
var BuilderDeleteTeacherKeyboard tgbotapi.InlineKeyboardMarkup

var BuilderOlimpsKeyboard tgbotapi.InlineKeyboardMarkup
var BuilderGetOlimpsKeyboard tgbotapi.InlineKeyboardMarkup
var BuilderDeleteOlimpsKeyboard tgbotapi.InlineKeyboardMarkup

func init() {
	BuilderSubjectsForTreker, SortBuilderSubjectsKeyboard, BuilderDeleteSubjectsForTreker = buttons("sub", lexicon.SubjectsForButton, 2)
	// ......... \\
	BuilderStageKeyboards, SortBuilderStageKeyboard, BuilderDeleteStageKeyboards = buttons("stage", lexicon.StagesTracker, 2)
	// ......... \\
	BuilderTeacherKeyboards, BuilderGetTeacherKeyboard, BuilderDeleteTeacherKeyboard = buttons("teacher", lexicon.TeacherTracker, 3)
	// ......... \\
	BuilderOlimpsKeyboard, BuilderGetOlimpsKeyboard, BuilderDeleteOlimpsKeyboard = buttons("olimp", lexicon.TrackerOlimps, 1)
}

func buttons(data string, slice []string, step int) (tgbotapi.InlineKeyboardMarkup, tgbotapi.InlineKeyboardMarkup, tgbotapi.InlineKeyboardMarkup) {
	slice0, slice1, slice2 := [][]tgbotapi.InlineKeyboardButton{}, [][]tgbotapi.InlineKeyboardButton{}, [][]tgbotapi.InlineKeyboardButton{}
	slice00, slice11, slice22 := []tgbotapi.InlineKeyboardButton{}, []tgbotapi.InlineKeyboardButton{}, []tgbotapi.InlineKeyboardButton{}
	for i := range slice {
		if i%step == 0 && i > 1 {
			slice0 = append(slice0, slice00)
			slice1 = append(slice1, slice11)
			slice2 = append(slice2, slice22)
			slice00, slice11, slice22 = []tgbotapi.InlineKeyboardButton{}, []tgbotapi.InlineKeyboardButton{}, []tgbotapi.InlineKeyboardButton{}
		}
		slice00 = append(slice00, tgbotapi.NewInlineKeyboardButtonData(
			slice[i], fmt.Sprintf("tracker;add;%s;%d", data, i)))
		slice11 = append(slice11, tgbotapi.NewInlineKeyboardButtonData(
			slice[i], fmt.Sprintf("tracker;get;%s;%d", data, i)))
		slice22 = append(slice22, tgbotapi.NewInlineKeyboardButtonData(
			slice[i], fmt.Sprintf("tracker;delete;%s;%d", data, i)))
	}
	builder1 := tgbotapi.NewInlineKeyboardMarkup(slice0...)
	builder2 := tgbotapi.NewInlineKeyboardMarkup(slice1...)
	builder3 := tgbotapi.NewInlineKeyboardMarkup(slice2...)
	return builder1, builder2, builder3
}
