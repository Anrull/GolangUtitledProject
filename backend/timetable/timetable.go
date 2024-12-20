package timetable

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

var table map[string]map[string]map[string][][]string
var extraTable map[string]map[string]map[string][][]string
var scheduleTeacher map[string]map[string]map[string][][]interface{}
var weeks map[string]string
var crutches = map[float64]string{1.0: "1", 2.0: "2", 3.0: "3", 4.0: "4", 5.0: "5", 6.0: "6", 7.0: "7", 8.0: "8"}
var Teachers []string

var dictDaysOfWeek = map[string]string{
	"Monday":    "0",
	"Tuesday":   "1",
	"Wednesday": "2",
	"Thursday":  "3",
	"Friday":    "4",
	"Saturday":  "0",
	"Sunday":    "0",
}

var dictDaysOfWeekTomorrow = map[string]string{
	"Monday":    "1",
	"Tuesday":   "2",
	"Wednesday": "3",
	"Thursday":  "4",
	"Friday":    "0",
	"Saturday":  "0",
	"Sunday":    "0",
}

func init() {
	data, err := os.ReadFile("data/dict.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла (предметы):", err)
	}
	err = json.Unmarshal(data, &table)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON (предметы):", err)
	}

	dataTeachers, err := os.ReadFile("data/teachers.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла (учителя):", err)
	}
	err = json.Unmarshal(dataTeachers, &scheduleTeacher)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON (учителя):", err)
	}

	dataWeeks, err := os.ReadFile("data/weeks.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла (недели):", err)
	}
	err = json.Unmarshal(dataWeeks, &weeks)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON (недели):", err)
	}

	dataExtraTable, err := os.ReadFile("data/extra_lessons.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла (доп. таблица):", err)
	}
	err = json.Unmarshal(dataExtraTable, &extraTable)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON (доп. таблица):", err)
	}

	for i := range scheduleTeacher["0"] {
		Teachers = append(Teachers, i)
	}

	sort.Strings(Teachers)
}

// GetTimetableText retrieves the timetable for a given week, day, and text class.
//
// Parameters:
// - week: The week of the timetable.
// - day: The day of the timetable.
// - textClass: The text class of the timetable.
//
// Returns:
// - [][]string: The timetable for the given week, day, and text class.
// - error: An error if the input parameters are invalid.
func GetTimetableText(week, day, textClass string) ([][]string, error) {
	if _, ok := table[week]; !ok {
		return nil, fmt.Errorf("неверная неделя: %s", week)
	}
	if _, ok := table[week][textClass]; !ok {
		return nil, fmt.Errorf("неверный text_class: %s", textClass)
	}
	if _, ok := table[week][textClass][day]; !ok {
		return nil, fmt.Errorf("неверный день: %s", day)
	}

	return table[week][textClass][day], nil
}


func GetExtraTimetableText(week, day, textClass string) ([][]string, error) {
	if _, ok := extraTable[week]; !ok {
		return nil, fmt.Errorf("неверная неделя: %s", week)
	}
	if _, ok := extraTable[week][textClass]; !ok {
		return nil, fmt.Errorf("неверный text_class: %s", textClass)
	}
	if _, ok := extraTable[week][textClass][day]; !ok {
		return nil, fmt.Errorf("неверный день: %s", day)
	}

	return extraTable[week][textClass][day], nil
}

// GetTimetableTeachersText retrieves the timetable for a given week, day, and teacher name.
//
// Parameters:
// - name: The name of the teacher.
// - week: The week of the timetable.
// - day: The day of the timetable.
//
// Returns:
// - [][]string: The timetable for the given week, day, and teacher name.
// - error: An error if the input parameters are invalid.
func GetTimetableTeachersText(name, week, day string) ([][]string, error) {
	// Проверка наличия ключей
	if _, ok := scheduleTeacher[week]; !ok {
		return nil, fmt.Errorf("неверная неделя: %s", week)
	}
	if _, ok := scheduleTeacher[week][name]; !ok {
		return nil, fmt.Errorf("преподаватель не найден: %s", name)
	}
	if _, ok := scheduleTeacher[week][name][day]; !ok {
		return nil, fmt.Errorf("неверный день: %s", day)
	}

	res := scheduleTeacher[week][name][day]

	stringSlice := make([][]string, len(res))
	for i, row := range res {
		strRow := make([]string, 2)

		if s, ok := row[0].(string); ok {
			strRow[0] = s
		} else {
			strRow[0] = ""
		}
		if ss, ok := row[1].(float64); ok {
			strRow[1] = fmt.Sprintf("%s", crutches[ss])
		}
		stringSlice[i] = strRow
	}

	return stringSlice, nil
}


// GetWeek returns the week string based on the given flag and res.
//
// Parameters:
// - flag: a boolean indicating whether to add or subtract days from the current date.
// - res: a boolean indicating whether to return "н" or "0" or "1" based on the week.
//
// Returns:
// - string: the week string.
// - error: an error if the week is not found.
func GetWeek(flag, res bool) (string, error) {
	date := time.Now()
	if !flag {
		if date.Weekday() == time.Saturday {
			date = date.AddDate(0, 0, 2)
		} else if date.Weekday() == time.Sunday {
			date = date.AddDate(0, 0, 3)
		}
	} else {
		if date.Weekday() >= time.Friday {
			date = date.AddDate(0, 0, 3)
		}
	}

	dateString := date.Format("2006-01-02")
	week, exists := weeks[dateString]
	if !exists {
		return "", fmt.Errorf("неделя не найдена для даты: %s", dateString)
	}
	if res {
		if week == "н" {
			return "1", nil
		}
		return "0", nil
	}
	return week, nil
}

func GetNextWeek() (string, error) {
	date := time.Now()
	if date.Weekday() == time.Saturday {
		date = date.AddDate(0, 0, 1)
	} else if date.Weekday() == time.Friday {
		date = date.AddDate(0, 0, 2)
	} else if date.Weekday() == time.Thursday {
		date = date.AddDate(0, 0, 3)
	}

	dateString := date.Format("2006-01-02")
	week, exists := weeks[dateString]
	if !exists {
		return "", fmt.Errorf("неделя не найдена для даты: %s", dateString)
	}
	return week, nil
}

// GetDayToday returns the day of the week for today.
//
// No parameters.
// Returns a string representing the day of the week.
func GetDayToday() string {
	date := time.Now()
	if date.Hour() > 12 || (date.Hour() == 12 && date.Minute() >= 30) {
		date = date.AddDate(0, 0, 1)
	}
	return dictDaysOfWeek[date.Weekday().String()]
}

// GetDayTomorrow returns the day of the week for tomorrow.
//
// No parameters.
// Returns a string representing the day of the week for tomorrow.
func GetDayTomorrow() string {
	return dictDaysOfWeekTomorrow[time.Now().Weekday().String()]
}

func GetDayTodayFeedback() string {
	date := time.Now()
	return dictDaysOfWeek[date.Weekday().String()]
}
