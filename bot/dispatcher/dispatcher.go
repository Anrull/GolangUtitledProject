package dispatcher

import (
	"awesomeProject/bot"
	"awesomeProject/bot/feedback"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"

	"strings"
)

var slice []time.Duration

func Dispatcher(update *tgbotapi.Update) {
	if len(slice) > 50 {
		var sum time.Duration

		mn := slice[0]
		mx := slice[0]
		for _, d := range slice {
			if d > mx {
				mx = d
			}
			if d < mn {
				mn = d
			}
			sum += d
		}
		average := sum / time.Duration(len(slice))
		fmt.Println("Среднее время выполнения:", average)
		fmt.Println("Максимальное время выполнения:", mx)
		fmt.Println("Минимальное время выполнения:", mn)
		slice = []time.Duration{}
	}
	startTime := time.Now()
	if update.Message != nil {
		message := update.Message
		if message.IsCommand() {
			CommandsHandling(message)
		} else {
			MessageHandler(message)
		}
		elapsedTime := time.Since(startTime)
		slice = append(slice, elapsedTime)
		return
	}
	if update.CallbackQuery != nil {
		query := update.CallbackQuery
		lstQ := strings.Split(query.Data, ";")
		fmt.Println(lstQ)
		if lstQ[0] == "timetable" {
			TimetableCallbackQuery(query, lstQ)
		} else if lstQ[0] == "tracker" {
			TrackerCallbackQuery(query, lstQ)
		} else if lstQ[0] == "main" {
			MainCallbackQuery(query, lstQ)
		} else if lstQ[0] == "yn" {
			YNCallbackQuery(query, lstQ)
		} else if lstQ[0] == "menu" {
			MenuCallbackQuery(query, lstQ)
		} else if lstQ[0] == "admin" {
			AdminPanelHandler(query, lstQ[1], lstQ...)
		} else if lstQ[0] == "lesson" {
			feedback.Handler(query.Message, lstQ...)
		}
		elapsedTime := time.Since(startTime)
		slice = append(slice, elapsedTime)
		bot.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
	}
}
