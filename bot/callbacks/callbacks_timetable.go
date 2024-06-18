package callbacks

import (
	"awesomeProject/bot/lexicon"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var BuilderDays0 = tgbotapi.NewInlineKeyboardMarkup()
var BuilderDays1 = tgbotapi.NewInlineKeyboardMarkup()

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
}
