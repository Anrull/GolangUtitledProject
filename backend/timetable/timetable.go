package timetable

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var table map[string]map[string]map[string][][]string
var scheduleTeacher map[string]map[string]map[string][][]interface{}
var weeks map[string]string
var crutches = map[float64]string{1.0: "1", 2.0: "2", 3.0: "3", 4.0: "4", 5.0: "5", 6.0: "6", 7.0: "7", 8.0: "8"}

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
}

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

func GetWeek() string {
	return time.Now().Format("2006-01-02")
}
