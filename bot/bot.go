package bot

import (
	"awesomeProject/pkg/env"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot, _ = tgbotapi.NewBotAPI(env.GetValue("TOKEN"))
var TechnicalWork bool

func init() {
	TechnicalWork = false
}

func Send(c tgbotapi.Chattable) {
	if _, err := Bot.Send(c); err != nil {
		log.Println(fmt.Sprintf("Error sending message: %s", err.Error()))
	}
}

func Request(c tgbotapi.Chattable) {
	_, err := Bot.Request(c)
	if err != nil {
		log.Println(fmt.Sprintf("Error sending message: %s", err.Error()))
	}
}
