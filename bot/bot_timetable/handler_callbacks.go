package bot_timetable

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/data/db"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var days = map[string]string{"0": "пн", "1": "вт", "2": "ср", "3": "чт", "4": "пт"}
var weeks = map[string]string{"0": "чет", "1": "нечет"}

func DaysHandler(ChatID int64, week, day string) {
	role, err := db.Get(ChatID, "role")
	if err != nil {
		log.Println(err)
		Bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка свзяи с db"))
		return
	}

	var schedule [][]string
	n, err := count("data/temp/images")
	filename := fmt.Sprintf("data/temp/images/schedule%d.png", n)
	if role == "student" {
		res, err := db.Get(ChatID, "classes")
		if err != nil {
			log.Println(err)
			Bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка свзяи с db"))
			return
		}
		schedule, err = timetable.GetTimetableText(week, day, res)
		timetable.DrawTimetable(schedule,
			fmt.Sprintf("%s, нед: %s, день: %s", res, weeks[week], days[day]), false, n)
	} else {
		name, err := db.Get(ChatID, "name_teacher")
		if err != nil {
			log.Println(err)
			Bot.Send(tgbotapi.NewMessage(ChatID, "Произошла ошибка свзяи с db"))
			return
		}
		schedule, err = timetable.GetTimetableTeachersText(name, week, day)
		timetable.DrawTimetable(schedule,
			fmt.Sprintf("%s, нед: %s, день: %s", name, weeks[week], days[day]), true, n)
	}
	photoBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	Bot.Send(tgbotapi.NewPhoto(ChatID, tgbotapi.FileBytes{
		Name:  "photo.jpg",
		Bytes: photoBytes,
	}))
}

func count(path string) (int, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("не удалось прочитать директорию: %w", err)
	}

	count_ := 0
	for _, file := range files {
		if !file.IsDir() {
			count_++
		}
	}

	return count_, nil
}
