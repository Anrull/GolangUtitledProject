package botTracker

import (
	"awesomeProject/bot"
	"awesomeProject/bot/callbacks"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"fmt"
	"log"
	"strconv"
)

func SubjectsCallbacksHandler(message *tgbotapi.Message, method, index string) {
	if method == "add" {
		i, err := strconv.Atoi(index)
		bot.Logging(message, err)
		value := lexicon.SubjectsForButton[i]
		err = db.AddTracker(message, "olimps", value)
		if err != nil {
			bot.Logging(message, err)
			return
		}

		will := bot.CopyInlineKeyboard(callbacks.BuilderOlimpsKeyboard)
		will.InlineKeyboard = will.InlineKeyboard[:lexicon.OlimpListStep]
		will.InlineKeyboard = append(will.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.OlimpListLeft, fmt.Sprintf("tracker;add;olimp;nil;0;%d;min", len(callbacks.BuilderOlimpsKeyboard.InlineKeyboard))),
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.OlimpListRight, fmt.Sprintf("tracker;add;olimp;nil;0;%d;plus", len(callbacks.BuilderOlimpsKeyboard.InlineKeyboard)))))
		msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, "Выберите олимпиаду", will)

		bot.Send(msg)
		return
	}
	if method == "someget" {
		i, err := strconv.Atoi(index)
		bot.Logging(message, err)
		value := "sub||nil"
		if i != 999 {
			value = "sub||" + lexicon.SubjectsForButton[i]
		}
		err = db.AddTracker(message, "filter", value)
		bot.Logging(message, err)
		will := bot.CopyInlineKeyboard(callbacks.SomeGetBuilderOlimpsKeyboard)
		will.InlineKeyboard = will.InlineKeyboard[:lexicon.OlimpListStep]
		will.InlineKeyboard = append(will.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.OlimpListLeft, fmt.Sprintf("tracker;someget;olimp;nil;0;%d;min", len(callbacks.SomeGetBuilderOlimpsKeyboard.InlineKeyboard))),
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.OlimpListRight, fmt.Sprintf("tracker;someget;olimp;nil;0;%d;plus", len(callbacks.SomeGetBuilderOlimpsKeyboard.InlineKeyboard)))))
		msg := tgbotapi.NewEditMessageReplyMarkup(
			message.Chat.ID, message.MessageID, will)
		bot.Send(msg)
		return
	}
	if method == "get" {
		i, err := strconv.Atoi(index)
		bot.Logging(message, err)
		sub := lexicon.SubjectsForButton[i]
		name, err := db.GetTracker(message, "name")
		err = db.AddTracker(message, "get_olimps",
			fmt.Sprintf("sub||%s;;olimp||nil;;stage||nil;;teacher||nil", sub))
		if err != nil {
			bot.Logging(message, err)
			return
		}
		records, err := db.GetRecords(name, sub, "nil", "nil", "nil")
		if err != nil {
			bot.Logging(message, err)
			return
		}
		sendOlimps(message, name, records)
	}
}

func OlimpsCallbacksHandler(message *tgbotapi.Message, status, spifMin, spifMax, role, method string) {
	step := lexicon.OlimpListStep
	if status == "nil" {
		spif, err := strconv.Atoi(spifMin)
		bot.Logging(message, err)
		spifMax, err := strconv.Atoi(spifMax)
		bot.Logging(message, err)
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
		var will tgbotapi.InlineKeyboardMarkup
		if method == "add" {
			will = bot.CopyInlineKeyboard(callbacks.BuilderOlimpsKeyboard)
		} else if method == "someget" {
			will = bot.CopyInlineKeyboard(callbacks.SomeGetBuilderOlimpsKeyboard)
		} else if method == "get" {
			will = bot.CopyInlineKeyboard(callbacks.BuilderGetOlimpsKeyboard)
		}
		will.InlineKeyboard = will.InlineKeyboard[spif_:mx]

		will.InlineKeyboard = append(will.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.OlimpListLeft, fmt.Sprintf("tracker;%s;olimp;nil;%d;%d;min", method, spif_, spifMax)),
			tgbotapi.NewInlineKeyboardButtonData(
				lexicon.OlimpListRight, fmt.Sprintf("tracker;%s;olimp;nil;%d;%d;plus", method, spif_, spifMax))))

		msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, will)
		bot.Send(msg)
		return
	}
	if method == "add" {
		index, err := strconv.Atoi(status)
		bot.Logging(message, err)
		olimp := lexicon.TrackerOlimps[index]
		err = db.AddTracker(message, "olimps", olimp)
		if err != nil {
			bot.Logging(message, err)
			return
		}
		msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, "Выберите достигнутый этап", callbacks.BuilderStageKeyboards)
		bot.Send(msg)
		return
	}
	if method == "someget" {
		index, err := strconv.Atoi(status)
		bot.Logging(message, err)
		olimp := "olimp||nil"
		if index != 999 {
			olimp = "olimp||" + lexicon.TrackerOlimps[index]
		}
		err = db.AddTracker(message, "filter", olimp)
		if err != nil {
			bot.Logging(message, err)
			return
		}
		msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, callbacks.SomeGetBuilderStageKeyboards)
		bot.Send(msg)
		return
	}
	if method == "get" {
		index, err := strconv.Atoi(status)
		bot.Logging(message, err)
		olimp := lexicon.TrackerOlimps[index]
		name, err := db.GetTracker(message, "name")
		err = db.AddTracker(message, "get_olimps",
			fmt.Sprintf("sub||nil;;olimp||%s;;stage||nil;;teacher||nil", olimp))
		if err != nil {
			bot.Logging(message, err)
			return
		}
		records, err := db.GetRecords(name, "nil", olimp, "nil", "nil")
		if err != nil {
			bot.Logging(message, err)
			return
		}
		sendOlimps(message, name, records)
	}
}

func StageCallbacksHandler(message *tgbotapi.Message, method, index string) {
	if method == "add" {
		i, err := strconv.Atoi(index)
		bot.Logging(message, err)
		value := lexicon.StagesTracker[i]
		err = db.AddTracker(message, "olimps", value)
		if err != nil {
			bot.Logging(message, err)
			return
		}
		msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, "Выберите наставника", callbacks.BuilderTeacherKeyboards)
		bot.Send(msg)
		return
	}
	if method == "someget" {
		i, err := strconv.Atoi(index)
		bot.Logging(message, err)
		value := "stage||nil"
		if i != 999 {
			value = "stage||" + lexicon.StagesTracker[i]
		}
		err = db.AddTracker(message, "filter", value)
		if err != nil {
			bot.Logging(message, err)
			return
		}
		msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, callbacks.SomeGetBuilderTeacherKeyboards)
		bot.Send(msg)
		return
	}
	if method == "get" {
		i, err := strconv.Atoi(index)
		bot.Logging(message, err)
		stage := lexicon.StagesTracker[i]
		name, err := db.GetTracker(message, "name")
		err = db.AddTracker(message, "get_olimps",
			fmt.Sprintf("sub||nil;;olimp||nil;;stage||%s;;teacher||nil", stage))
		if err != nil {
			bot.Logging(message, err)
			return
		}
		records, err := db.GetRecords(name, "nil", "nil", stage, "nil")
		if err != nil {
			bot.Logging(message, err)
			return
		}
		sendOlimps(message, name, records)
	}
}

func TeachersCallbacksHandler(message *tgbotapi.Message, method, index string) {
	if method == "add" {
		i, err := strconv.Atoi(index)
		bot.Logging(message, err)
		value := lexicon.TeacherTracker[i]
		err = db.AddTracker(message, "olimps", value)
		if err != nil {
			bot.Logging(message, err)
			return
		}
		textSlice, err := db.GetTracker(message, "olimps")
		if err != nil || textSlice == "" {
			return
		}
		slice := strings.Split(textSlice, ";")
		text := fmt.Sprintf("Подтвердите правильность введенных вами данных:\n• %s\n• %s\n• %s\n• %s", slice[0], slice[1], slice[2], slice[3])
		msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, callbacks.BuilderYNAddRecord)
		bot.Send(msg)
		return
	}
	if method == "get" {
		i, err := strconv.Atoi(index)
		bot.Logging(message, err)
		teacher := lexicon.TeacherTracker[i]
		name, err := db.GetTracker(message, "name")
		err = db.AddTracker(message, "get_olimps",
			fmt.Sprintf("sub||nil;;olimp||nil;;stage||nil;;teacher||%s", teacher))
		if err != nil {
			bot.Logging(message, err)
			return
		}
		records, err := db.GetRecords(name, "nil", "nil", "nil", teacher)
		if err != nil {
			bot.Logging(message, err)
			return
		}
		sendOlimps(message, name, records)
		return
	}
	if method == "someget" {
		i, err := strconv.Atoi(index)
		bot.Logging(message, err)
		value := "teacher||nil"
		if i != 999 {
			value = "teacher||" + lexicon.TeacherTracker[i]
		}
		err = db.AddTracker(message, "filter", value)
		if err != nil {
			bot.Logging(message, err)
			return
		}
		textSlice, err := db.GetTracker(message, "filter")
		if err != nil || textSlice == "" {
			return
		}
		slice := strings.Split(textSlice, ";;")

		sub, olimp, stage, teacher := getSettings(slice)

		name, err := db.GetTracker(message, "name")
		if err != nil {
			bot.Logging(message, err)
			return
		}
		records, err := db.GetRecords(name, sub, olimp, stage, teacher)
		if err != nil {
			bot.Logging(message, err)
			return
		}
		sendOlimps(message, name, records)
	}
}

func WithoutFiltersCallbacksHandler(message *tgbotapi.Message) {
	name, _ := db.GetTracker(message, "name")
	err := db.AddTracker(message, "get_olimps",
		"sub||nil;;olimp||nil;;stage||nil;;teacher||nil")
	if err != nil {
		bot.Logging(message, err)
		return
	}
	records, err := db.GetRecords(name, "nil", "nil", "nil", "nil")
	if err != nil {
		bot.Logging(message, err)
		return
	}
	sendOlimps(message, name, records)
}

func logging(message *tgbotapi.Message, err error) {
	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка связи с db"))
	}
}

func format(name, stage string, records *[]db.Records) []string {
	if len(*records) == 0 {
		return []string{"Нет олимпиад"}
	}
	builder := &strings.Builder{}
	builder.WriteString(name + ", " + stage)
	builder.WriteString("\n\n💬 Ваши олимпиады:")
	var result []string
	for i := len(*records) - 1; i >= 0; i-- {
		record := (*records)[i]
		if len(builder.String()) > 3800 {
			result = append(result, builder.String())
			builder.Reset()
		}
		builder.WriteString("\n\n")
		builder.WriteString("<blockquote>🧩 ")
		builder.WriteString(record.Olimps)
		builder.WriteString("\n🏆 ")
		builder.WriteString(record.Stage)
		builder.WriteString("\n📚 ")
		builder.WriteString(record.Subjects)
		builder.WriteString("\n👨‍🏫 ")
		builder.WriteString(record.Teachers)
		builder.WriteString("</blockquote>")
	}
	result = append(result, builder.String())
	return result
}

func sendOlimps(message *tgbotapi.Message, name string, records *[]db.Records) {
	class, _ := db.GetTracker(message, "stage")
	res := format(name, class, records)
	if len(res) == 1 {
		msg := tgbotapi.NewEditMessageTextAndMarkup(
			message.Chat.ID, message.MessageID, res[0], callbacks.ButtonsAfterOlimps)
		msg.ParseMode = tgbotapi.ModeHTML
		bot.Send(msg)
		return
	}
	for i, elem := range res[:len(res)-1] {
		if i == 0 {
			msg := tgbotapi.NewEditMessageText(
				message.Chat.ID, message.MessageID, elem)
			msg.ParseMode = tgbotapi.ModeHTML
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(message.Chat.ID, elem)
			msg.ParseMode = tgbotapi.ModeHTML
			bot.Send(msg)
		}
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, res[len(res)-1])
	msg.ReplyMarkup = callbacks.ButtonsAfterOlimps
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

func getSettings(slice []string) (string, string, string, string) {
	var sub, olimp, stage, teacher string

	for _, v := range slice {
		vv := strings.Split(v, "||")
		if vv[0] == "sub" {
			sub = vv[1]
		} else if vv[0] == "olimp" {
			olimp = vv[1]
		} else if vv[0] == "stage" {
			stage = vv[1]
		} else if vv[0] == "teacher" {
			teacher = vv[1]
		}
	}

	return sub, olimp, stage, teacher
}
