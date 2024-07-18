package dispatcher

import (
	"awesomeProject/bot"
	"awesomeProject/bot/feedback"
	"awesomeProject/data/db"
	"database/sql"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func AdminPanelHandler(query *tgbotapi.CallbackQuery, role string, someParams ...string) {
	message := query.Message

	switch role {
	case "count":
		n, err := db.GetUserCount()
		if err != nil {
			log.Println("Ошибка в подсчете пользователей")
		}
		msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID,
			fmt.Sprintf("Текущее количество пользователей — %d", n), bot.AdminPanelEscape)
		bot.Send(msg)
	case "escape":
		msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID,
			"Панель Администратора", bot.AdminPanel)
		bot.Send(msg)
	case "other":
		msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, bot.AdminPanelShutdown)
		bot.Send(msg)
	case "get_db":
		msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID,
			"Панель Администратора\n\nВыберите формат", bot.AdminPanelXLSX)
		bot.Send(msg)
	case "mode":
		msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID,
			fmt.Sprintf("Панель Администратора\n\nВыберите таблицу\n\nФормат: %s", someParams[2]),
			bot.AdminPanelDB)
		bot.Send(msg)
	case "get_logs":
		bot.Request(tgbotapi.NewCallback(query.ID, "Логирование пока не настроено"))
	case "shutdown":
		bot.Request(tgbotapi.NewCallback(query.ID, "Бот выключен"))
		os.Exit(0)
	case "fb":
		msg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, bot.AdminFB)
		bot.Send(msg)
	case "fbHandler":
		feedback.HandlerInfo(message, someParams[2:]...)
	default:
		getDB(message.Chat.ID, role, message.Text[len(message.Text)-4:])
	}
}

func getDB(ChatID int64, mode, format string) {
	switch mode {
	case "all":
		if format == " .db" {
			bot.SendFile(ChatID, "data/db/users.db", "users.db", "")
			bot.SendFile(ChatID, "data/db/records.db", "records.db", "")
			bot.SendFile(ChatID, "data/db/tracker.db", "tracker.db", "")
			bot.SendFile(ChatID, "data/db/student.db", "students.db", "")
			bot.SendFile(ChatID, "data/db/feedback.db", "feedback.db", "time")
		} else {
			err := exportToXLSX("data/db/users.db", "data/temp/users.xlsx", "users")
			if err != nil {
				log.Println(err)
				return
			}
			err = exportToXLSX("data/db/records.db", "data/temp/records.xlsx", "records")
			if err != nil {
				log.Println(err)
				return
			}
			err = exportToXLSX("data/db/tracker.db", "data/temp/tracker.xlsx", "trackers")
			if err != nil {
				log.Println(err)
				return
			}
			err = exportToXLSX("data/db/student.db", "data/temp/student.xlsx", "students")
			if err != nil {
				log.Println(err)
				return
			}
			err = exportToXLSX("data/db/feedback.db", "data/temp/feedback.xlsx", "feedback_lessons")
			if err != nil {
				log.Println(err)
				return
			}

			bot.SendFile(ChatID, "data/temp/feedback.xlsx", "feedback.xlsx", "time")
			bot.SendFile(ChatID, "data/temp/users.xlsx", "users.xlsx", "time")
			bot.SendFile(ChatID, "data/temp/records.xlsx", "records.xlsx", "time")
			bot.SendFile(ChatID, "data/temp/tracker.xlsx", "tracker.xlsx", "time")
			bot.SendFile(ChatID, "data/temp/student.xlsx", "student.xlsx", "time")
		}
	case "users":
		if format == " .db" {
			bot.SendFile(ChatID, "data/db/users.db", "users.db", "time")
		} else {
			err := exportToXLSX("data/db/users.db", "data/temp/users.xlsx", "users")
			if err != nil {
				log.Println(err)
				return
			}
			bot.SendFile(ChatID, "data/temp/users.xlsx", "users.xlsx", "time")
		}
	case "records":
		if format == " .db" {
			bot.SendFile(ChatID, "data/db/records.db", "records.db", "time")
		} else {
			err := exportToXLSX("data/db/records.db", "data/temp/records.xlsx", "records")
			if err != nil {
				log.Println(err)
				return
			}
			bot.SendFile(ChatID, "data/temp/records.xlsx", "records.xlsx", "time")
		}
	case "tracker":
		if format == " .db" {
			bot.SendFile(ChatID, "data/db/tracker.db", "tracker.db", "time")
		} else {
			err := exportToXLSX("data/db/tracker.db", "data/temp/tracker.xlsx", "trackers")
			if err != nil {
				log.Println(err)
				return
			}
			bot.SendFile(ChatID, "data/temp/tracker.xlsx", "tracker.xlsx", "time")
		}
	case "students":
		if format == " .db" {
			bot.SendFile(ChatID, "data/db/student.db", "students.db", "time")
		} else {
			err := exportToXLSX("data/db/student.db", "data/temp/student.xlsx", "students")
			if err != nil {
				log.Println(err)
				return
			}
			bot.SendFile(ChatID, "data/temp/student.xlsx", "student.xlsx", "time")
		}
	case "fb":
		if format == " .db" {
			bot.SendFile(ChatID, "data/db/feedback.db", "feedback.db", "time")
		} else {
			err := exportToXLSX("data/db/feedback.db", "data/temp/feedback.xlsx", "feedback_lessons")
			if err != nil {
				log.Println(err)
				return
			}
			bot.SendFile(ChatID, "data/temp/feedback.xlsx", "feedback.xlsx", "time")
		}
	}
}

func exportToXLSX(dbPath, xlsxPath, tableName string) error {
	data, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("ошибка открытия базы данных: %w", err)
	}
	defer data.Close()

	rows, err := data.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	xlsx := excelize.NewFile()
	sheet := xlsx.GetSheetName(xlsx.GetActiveSheetIndex())

	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("ошибка получения заголовков столбцов: %w", err)
	}
	for i, column := range columns {
		xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", string(rune('A'+i)), 1), column)
	}

	rowNum := 2
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err = rows.Scan(valuePtrs...); err != nil {
			return fmt.Errorf("ошибка сканирования строки: %w", err)
		}

		for i, value := range values {
			xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", string(rune('A'+i)), rowNum), value)
		}
		rowNum++
	}

	if err = xlsx.SaveAs(xlsxPath); err != nil {
		return fmt.Errorf("ошибка сохранения файла .xlsx: %w", err)
	}

	return nil
}
