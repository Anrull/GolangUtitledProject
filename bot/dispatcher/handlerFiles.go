package dispatcher

import (
	"awesomeProject/bot"
	"awesomeProject/bot/logger"
	"awesomeProject/data/db"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
)

func FileHandler(message *tgbotapi.Message) {
	document := message.Document
	if strings.HasSuffix(message.Document.FileName, ".xlsx") {
		// var role string
		if slices.Contains([]string{"обновить", "/update", "update"}, strings.ToLower(message.Caption)) {
			// role = "update"
		} else if slices.Contains([]string{"заменить", "/replace", "replace"}, strings.ToLower(message.Caption)) {
			// role = "replace"
			getTracker(message, "data/AllRecords.xlsx")
			err := db.DeleteAllRecords()
			if err != nil {
				logger.Error("", err)
				return
			}
		} else {
			return
		}
		msg, err := bot.Bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Обновляем"))
		if err != nil {
			logger.Error("", err)
			return
		}
		// Define the path where the document will be downloaded
		downloadPath := "data/temp/download/" + document.FileName
		err = downloadFile(downloadPath, document)
		if err != nil {
			logger.Error("", err)
			return
		}
		defer os.Remove(downloadPath)

		log.Println("File downloaded to:", downloadPath)

		f, err := excelize.OpenFile(downloadPath)
		if err != nil {
			logger.Error("", err)
			return
		}
		defer f.Close()

		// Получаем список строк в первом листе
		rows, err := f.GetRows("Sheet1")
		if err != nil {
			logger.Error("", err)
			return
		}
		var startIndex = 0
		if rows[0][0] == "Дата" {
			startIndex = 1
		}

		for i := startIndex; i < len(rows); i++ {
			row := rows[i]
			// Читаем значения в каждой колонке
			if len(row) < 7 {
				bot.Send(tgbotapi.NewEditMessageText(message.Chat.ID, msg.MessageID, "Неверный формат"))
				return
			}
			date := row[0]
			name := row[1]
			class := row[2]
			olimp := row[3]
			stage := row[4]
			subject := row[5]
			teacher := row[6]

			err = db.AddRecord(name, class, olimp, subject, teacher, stage, date)
			if err != nil {
				logger.Error("Error add record: ", err)
				return
			}
		}
		bot.Send(tgbotapi.NewEditMessageText(message.Chat.ID, msg.MessageID, "Обновлено"))
	}
}

func downloadFile(path string, document *tgbotapi.Document) error {
	// Create a new file in the specified path
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %e", err)
	}
	defer file.Close()

	// Get the file from the Telegram API
	fileBytes, err := bot.Bot.GetFileDirectURL(document.FileID)
	if err != nil {
		return fmt.Errorf("error getting file URL: %e", err)
	}

	// Download the file
	resp, err := http.Get(fileBytes)
	if err != nil {
		return fmt.Errorf("error downloading file: %e", err)
	}
	defer resp.Body.Close()

	// Write the file to disk
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving file: %e", err)
	}

	return nil
}
