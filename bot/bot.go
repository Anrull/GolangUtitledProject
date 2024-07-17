package bot

import (
	"awesomeProject/pkg/env"
	"fmt"
	"log"
	"os"
	"time"

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

func SendFile(ChatID int64, filename, title, Caption string) {
	fileReader, _ := os.Open(filename)
	defer fileReader.Close()

	inputFile := tgbotapi.FileReader{
		Name:   title,
		Reader: fileReader,
	}

	msg := tgbotapi.NewDocument(ChatID, inputFile)
	if Caption != "time" {
		msg.Caption = Caption
	} else {
		msg.Caption = time.Now().Format("2006-01-02 15:04:05")
	}

	Send(msg)
}
