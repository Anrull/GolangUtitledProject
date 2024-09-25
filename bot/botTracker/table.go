package botTracker

import (
	"awesomeProject/bot"
	"awesomeProject/data/db"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

func GetTable(message *tgbotapi.Message) {
	tracker, err := db.GetInfoAboutPersonTracker(message.Chat.ID)
	if err != nil {
		log.Println("GetTable:", err)
		return
	}

	textSlice, err := db.GetTracker(message, "filter")
	if err != nil || textSlice == "" {
		return
	}
	slice := strings.Split(textSlice, ";;")
	sub, olimp, stage, teacher := getSettings(slice)

	records, err := db.GetRecords(tracker.Name, sub, olimp, stage, teacher)
	if err != nil {
		log.Println("GetTable:", err)
		return
	}

	f := excelize.NewFile()
	filename := fmt.Sprintf("data/temp/%s.xlsx", tracker.Name)

	// Устанавливаем заголовки столбцов
	f.SetCellValue("sheet1", "A1", "Дата")
	f.SetCellValue("sheet1", "B1", "Олимпиада")
	f.SetCellValue("sheet1", "C1", "Этап")
	f.SetCellValue("sheet1", "D1", "Предметы")
	f.SetCellValue("sheet1", "E1", "Преподаватели")

	// Записываем данные из слайса в файл
	for i, record := range *records {
		row := i + 2 // Начинаем со второй строки, т.к. первая строка - заголовки

		// Записываем данные в ячейки
		f.SetCellValue("sheet1", fmt.Sprintf("A%d", row), record.Date)
		f.SetCellValue("sheet1", fmt.Sprintf("B%d", row), record.Olimps)
		f.SetCellValue("sheet1", fmt.Sprintf("C%d", row), record.Stage)
		f.SetCellValue("sheet1", fmt.Sprintf("D%d", row), record.Subjects)
		f.SetCellValue("sheet1", fmt.Sprintf("E%d", row), record.Teachers)
	}

	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
	}

	bot.SendFile(message.Chat.ID, filename, fmt.Sprintf("%s.xlsx", tracker.Name), tracker.Name)
}
