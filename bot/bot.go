package bot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// var Bot, _ = tgbotapi.NewBotAPI("6510904282:AAHPDg_LN7_edScl1NoBbATdHKKQegpa8yI")
var Bot, _ = tgbotapi.NewBotAPI("5965445771:AAElQV8vDRy8h0mRKniDCbL9E_uIb6feeUc")

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
