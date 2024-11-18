package bot

import (
	"awesomeProject/bot/logger"
	"awesomeProject/pkg/env"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mymmrac/telego"
)

var Bot, _ = tgbotapi.NewBotAPI(env.GetValue("TOKEN"))
var BotTelego, _ = telego.NewBot(env.GetValue("TOKEN"))
var TechnicalWork bool

const (
	ModeHTML       = "HTML"
	ModeMarkdown   = "Markdown"
	ModeMarkdownV2 = "MarkdownV2"
)

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
		logger.Error("", err)
		Bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка связи с db"))
	}
}


func SendMiniApp(ChatID int64, textButton, textMessage, url, escapeText, escapeData, parseMode string, messageIds ...int) {
	button := telego.InlineKeyboardButton{
		Text:   textButton,                    // Текст кнопки
		WebApp: &telego.WebAppInfo{URL: url}, // URL вашего Mini App
	}

	buttonEscape := telego.InlineKeyboardButton{
		Text:         escapeText,
		CallbackData: escapeData,
	}

	keyboard := telego.InlineKeyboardMarkup{
		InlineKeyboard: [][]telego.InlineKeyboardButton{
			{button}, // Добавляем кнопку
		},
	}
	
	if escapeData != "nil" {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []telego.InlineKeyboardButton{buttonEscape})
	}
	if len(messageIds) > 0 {
		editMsg := telego.EditMessageTextParams{
			ChatID:    telego.ChatID{ID: ChatID},
			MessageID: messageIds[0],
			Text:      textMessage,
			ParseMode: parseMode,
			ReplyMarkup: &keyboard,
		}
		_, err := BotTelego.EditMessageText(&editMsg)

		if err != nil {
			log.Println(err)
		}
		return
	}
	// Отправляем сообщение с кнопкой
	params := telego.SendMessageParams{
		ChatID:    telego.ChatID{ID: ChatID},
		Text:      textMessage,
		ParseMode: parseMode,
		ReplyMarkup: &keyboard,
	}
	_, err := BotTelego.SendMessage(&params)

	if err != nil {
		log.Println(err)
	}
}

func TelegoSendWithKeyboard(ChatID int64, text, parseMode string, buttons telego.InlineKeyboardMarkup, messageIds ...int) {
	if len(messageIds) > 0 {
		editMsg := telego.EditMessageTextParams{
			ChatID:    telego.ChatID{ID: ChatID},
			MessageID: messageIds[0],
			Text:      text,
			ParseMode: parseMode,
		}
		if buttons.InlineKeyboard != nil {
			editMsg.ReplyMarkup = &buttons
		}
		_, err := BotTelego.EditMessageText(&editMsg)

		if err != nil {
			log.Println(err)
		}
		return
	}
	// Отправляем сообщение с кнопкой
	params := telego.SendMessageParams{
		ChatID:    telego.ChatID{ID: ChatID},
		Text:      text,
		ParseMode: parseMode,
	}
	if buttons.InlineKeyboard != nil {
		params.ReplyMarkup = &buttons
	}
	_, err := BotTelego.SendMessage(&params)

	if err != nil {
		log.Println(err)
	}
}