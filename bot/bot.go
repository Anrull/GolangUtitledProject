package bot

import (
	"awesomeProject/pkg/env"
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

// Send sends a message using the Bot.
//
// The function takes a parameter of type tgbotapi.Chattable, which represents the message to be sent.
// It returns nothing.
// If there is an error while sending the message, it logs the error message.
func Send(c tgbotapi.Chattable) {
	if _, err := Bot.Send(c); err != nil {
		log.Printf("Error sending message: %s", err.Error())
	}
}

// Request sends a request to the Bot.
//
// The function takes a parameter of type tgbotapi.Chattable, which represents the request to be sent.
// It returns nothing.
// If there is an error while sending the request, it logs the error message.
func Request(c tgbotapi.Chattable) {
	_, err := Bot.Request(c)
	if err != nil {
		log.Printf("Error sending message: %s", err.Error())
	}
}

// SendFile sends a file to the specified ChatID with the given filename, title, and caption.
//
// Parameters:
// - ChatID: the ID of the chat to send the file to
// - filename: the name of the file to send
// - title: the title of the file
// - Caption: the caption for the file. If Caption is "time", the caption will be the current time in the format "2006-01-02 15:04:05".
func SendFile(ChatID int64, filename, title, Caption string, flag ...bool) {
	fileReader, _ := os.Open(filename)
	if len(flag) == 0 {
		err := os.Remove(filename)
		if err != nil {
			log.Println("Ошибка при удалении файла:", err)
		}
	}
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

func Logging(message *tgbotapi.Message, err error) {
	if err != nil {
		log.Println(err)
		Bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка связи с db"))
	}
}