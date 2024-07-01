package callbacks

import (
	"awesomeProject/bot/lexicon"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"fmt"
)

// создание клавиатуры выбора предмета "tracker;add;sub;%d"

var SomeGetSubjectsTracker tgbotapi.InlineKeyboardMarkup
var BuilderSubjectsForTracker tgbotapi.InlineKeyboardMarkup
var SortBuilderSubjectsKeyboard tgbotapi.InlineKeyboardMarkup
var BuilderDeleteSubjectsForTreker tgbotapi.InlineKeyboardMarkup

// создание клавиатуры выбора по этапу

var BuilderDeleteStageKeyboards tgbotapi.InlineKeyboardMarkup
var BuilderStageKeyboards tgbotapi.InlineKeyboardMarkup
var SomeGetBuilderStageKeyboards tgbotapi.InlineKeyboardMarkup
var SortBuilderStageKeyboard tgbotapi.InlineKeyboardMarkup

// создание клавиатуры выбора наставника

var SomeGetBuilderTeacherKeyboards tgbotapi.InlineKeyboardMarkup
var BuilderTeacherKeyboards tgbotapi.InlineKeyboardMarkup
var BuilderGetTeacherKeyboard tgbotapi.InlineKeyboardMarkup
var BuilderDeleteTeacherKeyboard tgbotapi.InlineKeyboardMarkup

var SomeGetBuilderOlimpsKeyboard tgbotapi.InlineKeyboardMarkup
var BuilderOlimpsKeyboard tgbotapi.InlineKeyboardMarkup
var BuilderGetOlimpsKeyboard tgbotapi.InlineKeyboardMarkup
var BuilderDeleteOlimpsKeyboard tgbotapi.InlineKeyboardMarkup

var DeleteButton = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Удалить", "menu;filter;Удалить")))

var BuilderYNAddRecord = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(
	tgbotapi.NewInlineKeyboardButtonData("Да", "yn;AddRecord;yes"),
	tgbotapi.NewInlineKeyboardButtonData("Нет", "yn;AddRecord;no"),
))

var BuilderEscMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Назад", "menu;filter;Назад")))

var ButtonsAfterOlimps = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Назад", "menu;filter;Назад"),
		tgbotapi.NewInlineKeyboardButtonData(
			"Удалить", "menu;filter;Удалить")))

func init() {
	BuilderSubjectsForTracker, SortBuilderSubjectsKeyboard, BuilderDeleteSubjectsForTreker, SomeGetSubjectsTracker = buttons("sub", "", lexicon.SubjectsForButton, 2, 1)
	// ......... \\
	BuilderStageKeyboards, SortBuilderStageKeyboard, BuilderDeleteStageKeyboards, SomeGetBuilderStageKeyboards = buttons("stage", "", lexicon.StagesTracker, 2, 1)
	// ......... \\
	BuilderTeacherKeyboards, BuilderGetTeacherKeyboard, BuilderDeleteTeacherKeyboard, SomeGetBuilderTeacherKeyboards = buttons("teacher", "", lexicon.TeacherTracker, 3, 1)
	// ......... \\
	BuilderOlimpsKeyboard, BuilderGetOlimpsKeyboard, BuilderDeleteOlimpsKeyboard, SomeGetBuilderOlimpsKeyboard = buttons("olimp", ";nil;nil;nil", lexicon.TrackerOlimps, 1, 0)
}

func CreateButtonsDelete(max int) tgbotapi.InlineKeyboardMarkup {
	var sliceButtons []tgbotapi.InlineKeyboardButton
	var sliceRows [][]tgbotapi.InlineKeyboardButton

	for i := 0; i < max; i++ {
		if i%3 == 0 && i != 0 {
			sliceRows = append(sliceRows, sliceButtons)
			sliceButtons = []tgbotapi.InlineKeyboardButton{}
		}
		sliceButtons = append(sliceButtons, tgbotapi.NewInlineKeyboardButtonData(
			fmt.Sprintf("%d", i), fmt.Sprintf("del;filter;%d;%d", i, max)))
	}

	return tgbotapi.NewInlineKeyboardMarkup(sliceRows...)
}

func buttons(data, subData string, slice []string, step, minParam int) (tgbotapi.InlineKeyboardMarkup,
	tgbotapi.InlineKeyboardMarkup, tgbotapi.InlineKeyboardMarkup,
	tgbotapi.InlineKeyboardMarkup) {
	var slice0, slice1, slice2, slice3 [][]tgbotapi.InlineKeyboardButton
	var slice00, slice11, slice22, slice33 []tgbotapi.InlineKeyboardButton
	slice3 = append(slice3, tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Пропустить",
		fmt.Sprintf("tracker;someget;%s;%d%s", data, 999, subData))))
	for i := range slice {
		if i%step == 0 && i > minParam {
			slice0 = append(slice0, slice00)
			slice1 = append(slice1, slice11)
			slice2 = append(slice2, slice22)
			slice3 = append(slice3, slice33)
			slice00, slice11, slice22, slice33 = []tgbotapi.InlineKeyboardButton{}, []tgbotapi.InlineKeyboardButton{}, []tgbotapi.InlineKeyboardButton{}, []tgbotapi.InlineKeyboardButton{}
		}
		slice00 = append(slice00, tgbotapi.NewInlineKeyboardButtonData(
			slice[i], fmt.Sprintf("tracker;add;%s;%d%s", data, i, subData)))
		slice11 = append(slice11, tgbotapi.NewInlineKeyboardButtonData(
			slice[i], fmt.Sprintf("tracker;get;%s;%d%s", data, i, subData)))
		slice22 = append(slice22, tgbotapi.NewInlineKeyboardButtonData(
			slice[i], fmt.Sprintf("tracker;delete;%s;%d%s", data, i, subData)))
		slice33 = append(slice33, tgbotapi.NewInlineKeyboardButtonData(
			slice[i], fmt.Sprintf("tracker;someget;%s;%d%s", data, i, subData)))
	}
	builder1 := tgbotapi.NewInlineKeyboardMarkup(slice0...)
	builder2 := tgbotapi.NewInlineKeyboardMarkup(slice1...)
	builder3 := tgbotapi.NewInlineKeyboardMarkup(slice2...)
	builder4 := tgbotapi.NewInlineKeyboardMarkup(slice3...)
	return builder1, builder2, builder3, builder4
}
