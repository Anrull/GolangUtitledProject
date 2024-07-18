package feedback

import (
	"awesomeProject/bot"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"
	"fmt"
	"github.com/agnivade/levenshtein"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math"
	"strings"
	"time"
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
	}

	messages = append(messages, tgbotapi.NewMessage(message.Chat.ID, text))

	for _, i := range lexicon.Stages {
		if params[0] == "nowWeek" {
			records, err = db.GetFBLessonsByWeekTest(i, getWeekDates(time.Now()))
		} else if params[0] == "lastWeek" {
			records, err = db.GetFBLessonsByWeekTest(i, getPreviousWeekDates(time.Now()))
		}

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
		fmt.Println(result)

		builder := &strings.Builder{}
		builder.WriteString("Класс: <b><em>")
		builder.WriteString(i)
		builder.WriteString("</em></b>\n\n")

		for key := range result {
			res := result[key]
			builder.WriteString("<b><em>")
			builder.WriteString(key)

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
				builder.WriteString(fmt.Sprintf("</em></b>\nСредняя оценка: %.2f\n", mark))
			} else {
				builder.WriteString("</em></b>\nНет информации\n")
			}

			builder.WriteString(fmt.Sprintf("Не было уроков %d раз\n", count))
			builder.WriteString(fmt.Sprintf("Количество оценок %d раз\n\n", count+countLessons))
		}

		messages = append(messages, tgbotapi.NewMessage(message.Chat.ID, builder.String()))
	}

	for _, msg := range messages {
		msg.ParseMode = tgbotapi.ModeHTML
		bot.Send(msg)
	}
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
			if len(group) >= minPrefixMatch && len(subjectLower) >= minPrefixMatch &&
				strings.HasPrefix(subjectLower, string(group[0:minPrefixMatch])) &&
				strings.HasSuffix(subjectLower, string(group[len(subjectLower)-minPrefixMatch])) {
				foundGroup = group
				break
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
func getWeekDates(date time.Time) []string {
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

func getPreviousWeekDates(date time.Time) []string {
	// Отнимаем 7 дней от текущей даты, чтобы получить
	// соответствующий день недели на прошлой неделе.
	lastWeekDate := date.AddDate(0, 0, -7)

	// Вызываем getWeekDates с датой на прошлой неделе.
	return getWeekDates(lastWeekDate)
}
