package bot_tracker

import (
	"awesomeProject/bot"
	"awesomeProject/bot/callbacks"
	"awesomeProject/bot/lexicon"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"fmt"
	"log"
	"strconv"
)

var Bot = bot.Bot

func OlimpsCallbacksHandler(message *tgbotapi.Message, status, spifMin, spifMax, role string) {
	step := lexicon.OlimpListStep
	if status == "nil" {
		spif, err := strconv.Atoi(spifMin)
		logging(message, err)
		spifMax, err := strconv.Atoi(spifMax)
		logging(message, err)
		spif_ := 0
		if role == "min" {
			if !(spif-step < 0) {
				fmt.Println(true)
				spif_ = spif - step
			}
		} else {
			if !(spif+step > spifMax) {
				fmt.Println(false)
				spif_ = spif + step
			}
		}
		mx := spifMax
		if !(spif_+step > spifMax) {
			mx = spif_ + step
		}
		if !(spif-step < 0) {
			spif = spif - step
		}
		will := tgbotapi.InlineKeyboardMarkup{}
		willCopy_ := make([][]tgbotapi.InlineKeyboardButton, len(callbacks.BuilderOlimpsKeyboard.InlineKeyboard))
		copy(willCopy_, callbacks.BuilderOlimpsKeyboard.InlineKeyboard)
		will.InlineKeyboard = willCopy_[spif_:mx]
		//fmt.Println(spif_, mx, spifMax)

		will.InlineKeyboard = append(will.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.OlimpListLeft, fmt.Sprintf("tracker;add;olimp;nil;%d;%d;min", spif_, spifMax)),
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.OlimpListRight, fmt.Sprintf("tracker;add;olimp;nil;%d;%d;plus", spif_, spifMax))))

		msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, will)
		Bot.Send(msg)
	}
}

func logging(message *tgbotapi.Message, err error) {
	if err != nil {
		log.Println(err)
		Bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка связи с db"))
	}
}
