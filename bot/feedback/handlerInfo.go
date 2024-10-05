package feedback

import (
	"awesomeProject/bot"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/agnivade/levenshtein"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
)

func HandlerInfo(message *tgbotapi.Message, params ...string) {
	var records *[]db.FeedbackLesson
	var err error

	var messages []tgbotapi.MessageConfig
	var text string

	if params[0] == "nowWeek" {
		text = "Анализ уроков за эту неделю"
	} else if params[0] == "lastWeek" {
		text = "Анализ уроков за прошлую неделю"
	} else if params[0] == "lastMonth" {
        text = "Анализ уроков за прошлый месяц"
    } else if params[0] == "nowMonth" {
        text = "Анализ уроков за этот месяц"
    }

	messages = append(messages, tgbotapi.NewMessage(message.Chat.ID, text))

    f := excelize.NewFile()
    
    err = f.MergeCell("Sheet1", "A1", "H4")
    if err != nil {
        log.Println(err)
        return
    }
    
    f.SetCellValue("Sheet1", "A1", lexicon.InformationCaption)

	for _, i := range lexicon.Stages {
		records, err = db.GetFBLessonsByWeekTest(i, getDates(params[0]))

		if err != nil {
			log.Println(err)
			return
		}

		if records == nil {
			continue
		}

		result := findMostCommonSubjects(*records)
		if len(result) == 0 {
			continue
		}

        startIndex := 1

        index, _ := f.NewSheet(i)
        f.SetActiveSheet(index)

        f.SetCellValue(i, "A1", "Предмет")
        f.SetCellValue(i, "B1", "Средняя оценка")
        f.SetCellValue(i, "C1", "Не было уроков")
        f.SetCellValue(i, "D1", "Количество оценок")

		builder := &strings.Builder{}
		builder.WriteString("Класс: <b><em>")
		builder.WriteString(i)
		builder.WriteString("</em></b>\n\n")

		for key := range result {
            startIndex++
			res := result[key]

			builder.WriteString("<b><em>")
			builder.WriteString(key)

            f.SetCellValue(i, fmt.Sprintf("A%d", startIndex), key)

			var count, count1, count2, count3, count4, count5 int
			for _, j := range res {
				if j == "5" {
					count5++
				} else if j == "4" {
					count4++
				} else if j == "3" {
					count3++
				} else if j == "2" {
					count2++
				} else if j == "1" {
					count1++
				} else {
					count++
				}
			}

			markLesson := float64(count1 + (count2 * 2) + (count3 * 3) + (count4 * 4) + (count5 * 5))
			countLessons := count1 + count2 + count3 + count4 + count5
			mark := markLesson / float64(countLessons)

			if !math.IsNaN(mark) {
                f.SetCellValue(i, fmt.Sprintf("B%d", startIndex), fmt.Sprintf("%.2f", mark))
				builder.WriteString(fmt.Sprintf("</em></b>\nСредняя оценка: %.2f\n", mark))
			} else {
                f.SetCellValue(i, fmt.Sprintf("B%d", startIndex), "Нет информации")
				builder.WriteString("</em></b>\nНет информации\n")
			}

			builder.WriteString(fmt.Sprintf("Не было уроков %d раз\n", count))
			builder.WriteString(fmt.Sprintf("Количество оценок %d раз\n\n", count+countLessons))

            // f.SetCellValue(i, fmt.Sprintf("B%d", startIndex), fmt.Sprintf("%.2f", mark))
            f.SetCellValue(i, fmt.Sprintf("C%d", startIndex), fmt.Sprintf("%d", count))
            f.SetCellValue(i, fmt.Sprintf("D%d", startIndex), fmt.Sprintf("%d", count+countLessons))
		}
        
		messages = append(messages, tgbotapi.NewMessage(message.Chat.ID, builder.String()))
	}

	for _, msg := range messages {
		msg.ParseMode = tgbotapi.ModeHTML
		bot.Send(msg)
	}

    filename := "data/temp/fb.xlsx"
    defer os.Remove(filename)

    if err := f.SaveAs(filename); err != nil {
        log.Println(err)
        return
    }

    bot.SendFile(message.Chat.ID, filename, "Анализ уроков.xlsx", "time")
}

func findMostCommonSubjects(subjects []db.FeedbackLesson) map[string][]string {
	// Создаем словарь для группировки похожих названий предметов
	groupedSubjects := make(map[string][]string)
	groupedSubjectsCount := make(map[string][]string)

	minPrefixMatch := 6

	// Проходим по всем предметам в слайсе
	for _, subject := range subjects {
		// Приводим название предмета к нижнему регистру
		subjectLower := strings.ToLower(subject.TitleLesson)
		// Находим группу, к которой принадлежит текущий предмет
		foundGroup := ""
		for group := range groupedSubjects {
			// Проверяем совпадение префикса с учетом minPrefixMatch
			if llen(group) > minPrefixMatch && llen(subjectLower) < minPrefixMatch && llen(subjectLower)-minPrefixMatch < llen(group) {
				if strings.HasPrefix(subjectLower, string(group[0:minPrefixMatch])) &&
				strings.HasSuffix(subjectLower, 
					string(group[llen(subjectLower)-minPrefixMatch:])) {
				foundGroup = group
				break
			}
			}
			// Если префикс не совпадает, проверяем расстояние Левенштейна
			distance := levenshtein.ComputeDistance(subjectLower, group)
			// Если расстояние меньше 3 (можно настроить), считаем, что это одна группа
			if distance <= 3 {
				foundGroup = group
				break
			}
		}
		// Если группа найдена, добавляем к ней текущий предмет
		if foundGroup != "" {
			groupedSubjectsCount[foundGroup] = append(groupedSubjectsCount[foundGroup], subject.Status)
			groupedSubjects[foundGroup] = append(groupedSubjects[foundGroup], subject.TitleLesson)
		} else {
			// Иначе создаем новую группу с текущим предметом
			groupedSubjectsCount[subjectLower] = []string{subject.Status}
			groupedSubjects[subjectLower] = []string{subject.TitleLesson}
		}
	}

	// Создаем словарь для хранения наиболее часто встречающихся названий предметов
	mostCommonSubjects := make(map[string][]string)

	// Проходим по каждой группе предметов
	for i, group := range groupedSubjects {
		// Находим наиболее часто встречающийся предмет в текущей группе
		mostCommonSubject := findMostCommonSubject(group)
		// Добавляем его в результирующий словарь
		mostCommonSubjects[mostCommonSubject] = groupedSubjectsCount[i]
		//fmt.Println(i, groupedSubjectsCount[i])
	}

	// Возвращаем словарь с наиболее часто встречающимися названиями предметов
	return mostCommonSubjects
}

// Вспомогательная функция для нахождения самого частого элемента в слайсе
func findMostCommonSubject(subjects []string) string {
	subjectCounts := make(map[string]int)
	for _, subject := range subjects {
		subjectCounts[subject]++
	}
	maxCount := 0
	mostCommonSubject := ""
	for subject, count := range subjectCounts {
		if count > maxCount {
			maxCount = count
			mostCommonSubject = subject
		}
	}
	return mostCommonSubject
}

// getWeekDates возвращает слайс из строк, представляющих даты всех дней недели,
// начиная с воскресенья, для заданного дня недели.
func getWeekDates(date time.Time, param ...string) []string {
	if len(param) != 0 {
		if param[0] == "lastWeek" {
			date = date.AddDate(0, 0, -7)
		}
	}
	// Определяем номер дня недели (0 - воскресенье, 1 - понедельник и т.д.)
	dayOfWeek := int(date.Weekday())

	// Вычисляем количество дней, которое нужно отнять от текущей даты,
	// чтобы получить дату воскресенья.
	daysToSubtract := (dayOfWeek + 6) % 7 // Прибавляем 6, чтобы воскресенье было 0, а не 7

	// Получаем дату воскресенья.
	sunday := date.AddDate(0, 0, -daysToSubtract)

	// Создаем слайс для хранения дат.
	weekDates := make([]string, 5)

	// Заполняем слайс датами на всю неделю.
	for i := 0; i < 5; i++ {
		weekDates[i] = sunday.AddDate(0, 0, i).Format("2006-01-02")
	}

	return weekDates
}

// getMonthDates возвращает слайс из строк, представляющих даты всех дней месяца,
// для заданного дня.
func getMonthDates(date time.Time, param ...string) []string {
	if len(param) != 0 {
		if param[0] == "lastMonth" {
			firstDayOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())

			// Вычитаем 1 день, чтобы получить последний день предыдущего месяца.
			date = firstDayOfMonth.AddDate(0, 0, -1)
		}
	}

	// Получаем первый день месяца.
	firstDayOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())

	// Получаем количество дней в месяце.
	daysInMonth := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, date.Location()).Day()

	// Создаем слайс для хранения дат.
	monthDates := make([]string, daysInMonth)

	// Заполняем слайс датами на весь месяц.
	for i := 0; i < daysInMonth; i++ {
		monthDates[i] = firstDayOfMonth.AddDate(0, 0, i).Format("2006-01-02")
	}

	return monthDates
}

func getDates(param string) []string {
	if strings.Contains(param, "Week") {
		return getWeekDates(time.Now(), param)
	} else if strings.Contains(param, "Month") {
		return getMonthDates(time.Now(), param)
	}
	return []string{}
}

func llen(s string) int {
	var count int
	for range s {
		count++
	}
	return count
}