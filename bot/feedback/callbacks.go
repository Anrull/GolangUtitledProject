package feedback

import (
	"awesomeProject/data/db"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Escape = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(
	tgbotapi.NewInlineKeyboardButtonData("Скрыть уведомление", "lesson;escape")))

func GetFeedbackCallback(stage, sub, date string) tgbotapi.InlineKeyboardMarkup {
	var slice []tgbotapi.InlineKeyboardButton

	sub_ := db.AddTempFb(sub)

	for i := 1; i <= 5; i++ {
		slice = append(slice,
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d⭐️", i),
				fmt.Sprintf("lesson;%s;%s;%s;%d", stage, sub_, date, i)))
	}

	button1 := tgbotapi.NewInlineKeyboardButtonData(
		"Не было учителя/урока", fmt.Sprintf("lesson;%s;%s;%s;-1", stage, sub_, date))
	button2 := tgbotapi.NewInlineKeyboardButtonData(
		"Не буду отвечать", fmt.Sprintf("lesson;%s;%s;%s;-2", stage, sub_, date))
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(button1),
		slice,
		tgbotapi.NewInlineKeyboardRow(button2))
}
