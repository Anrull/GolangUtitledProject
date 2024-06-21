package bot_tracker

import (
	"awesomeProject/bot"
	"awesomeProject/bot/callbacks"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"fmt"
	"log"
	"strconv"
)

var Bot = bot.Bot

func SubjectCallbacksHandler(message *tgbotapi.Message) {
	//...
}

func OlimpsCallbacksHandler(message *tgbotapi.Message, status, spifMin, spifMax, role, method string) {
	step := lexicon.OlimpListStep
	if status == "nil" {
		spif, err := strconv.Atoi(spifMin)
		logging(message, err)
		spifMax, err := strconv.Atoi(spifMax)
		logging(message, err)
		spif_ := 0
		if role == "min" {
			if !(spif-step < 0) {
				spif_ = spif - step
			}
		} else {
			if !(spif+step > spifMax) {
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

		will.InlineKeyboard = append(will.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.OlimpListLeft, fmt.Sprintf("tracker;add;olimp;nil;%d;%d;min", spif_, spifMax)),
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.OlimpListRight, fmt.Sprintf("tracker;add;olimp;nil;%d;%d;plus", spif_, spifMax))))

		msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, will)
		Bot.Send(msg)
	} else if method == "add" {
		index, err := strconv.Atoi(status)
		logging(message, err)
		olimp := lexicon.TrackerOlimps[index]
		err = db.AddTracker(message, "olimps", olimp)
		if err != nil {
			logging(message, err)
			return
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите достигнутый этап")
		msg.ReplyMarkup = callbacks.BuilderStageKeyboards
		Bot.Send(msg)
	}
}

func SubjectsCallbacksHandler(message *tgbotapi.Message, method, index string) {
	if method == "add" {
		i, err := strconv.Atoi(index)
		logging(message, err)
		value := lexicon.SubjectsForButton[i]
		err = db.AddTracker(message, "olimps", value)
		if err != nil {
			logging(message, err)
			return
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите олимпиаду")
		msg.ReplyMarkup = callbacks.BuilderOlimpsKeyboard
		Bot.Send(msg)
	}
}

func logging(message *tgbotapi.Message, err error) {
	if err != nil {
		log.Println(err)
		Bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка связи с db"))
	}
}
