package botSchedule

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/bot"
	"awesomeProject/bot/feedback"
	"awesomeProject/bot/lexicon"
	"awesomeProject/bot/logger"
	"awesomeProject/data/db"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"awesomeProject/pkg/scheduler"

	"fmt"
)

func RunScheduler() {
	scheduler.NewScheduler(1, 11, 0, TasksSchedule)
	scheduler.NewScheduler(2, 11, 0, TasksSchedule)
	scheduler.NewScheduler(3, 11, 0, TasksSchedule)
	scheduler.NewScheduler(4, 11, 0, TasksSchedule)
	scheduler.NewScheduler(7, 13, 0, TasksSchedule)

	scheduler.NewScheduler(2, 18, 18, LessonOne)
	for i := 1; i <= 5; i++ {
		scheduler.NewScheduler(i, 5, 0, LessonOne)
		scheduler.NewScheduler(i, 6, 50, LessonTwo)
		scheduler.NewScheduler(i, 8, 40, LessonThree)
		scheduler.NewScheduler(i, 9, 40, LessonSeven)
	}

	logger.Info("Scheduler started")
}

func TasksSchedule() {
	users, _ := db.GetAllUsers()

	for _, user := range users {
		var photoByte []byte

		role, err := db.Get(user.UserID, "role")
		if err != nil {
			continue
		}
		week, err := timetable.GetWeek(false, true)
		if err != nil {
			continue
		}
		day := timetable.GetDayTomorrow()

		colors, err := db.GetColorByUserID(user.UserID)
		if err != nil {
			continue
		}

		if role == "student" {
			stage, err := db.Get(user.UserID, "classes")
			if err != nil {
				continue
			}
			lessons, err := timetable.GetTimetableText(week, day, stage)
			if err != nil {
				continue
			}
			
			extraLessons, _ := timetable.GetExtraTimetableText(week, day, stage)

			lessons = timetable.Merge(lessons, extraLessons)

			// if !reflect.DeepEqual(extraLessons, [][]string{{}, {}, {}}) {
			// 	extraPhotoByte, _ := timetable.DrawTimetable(
			// 		extraLessons, "Внеурочные занятия",
			// 		false, colors...)
				
			// 	defer SendPhotoByte(user.UserID, extraPhotoByte)
			// }

			photoByte, err = timetable.DrawTimetable(
				lessons, fmt.Sprintf("%s, нед: %s, день: %s",
					stage, lexicon.Week[week], lexicon.Day[day]),
				false, colors...)
			if err != nil {
				continue
			}
		} else {
			teacher, err := db.Get(user.UserID, "name_teacher")
			if err != nil {
				continue
			}
			lessons, err := timetable.GetTimetableTeachersText(teacher, week, day)
			if err != nil {
				continue
			}

			photoByte, err = timetable.DrawTimetable(
				lessons, fmt.Sprintf("%s, нед: %s, день: %s", teacher,
					lexicon.Week[week], lexicon.Day[day]),
				true, colors...)
			if err != nil {
				continue
			}
		}

		SendPhotoByte(user.UserID, photoByte)
	}
}

func SendFeedbackLessons(num int) {
	if !isTimeBefore(10, 0) {
		return
	}

	num--
	users, err := db.GetAllUsers()
	if err != nil {
		log.Println()
	}

	week, _ := timetable.GetWeek(false, true)
	day := timetable.GetDayTodayFeedback()

	for _, user := range users {
		if user.Role != "student" {
			continue
		}
		if user.Newsletter != "1" {
			continue
		}
		stage, _ := db.Get(user.UserID, "classes")
		lessons, _ := timetable.GetTimetableText(week, day, stage)

		if len(lessons) > num {
			msg := tgbotapi.NewMessage(user.UserID,
				fmt.Sprintf("Оцените как прошел урок: %s", lessons[num][0]))
			msg.ReplyMarkup = feedback.GetFeedbackCallback(stage,
				lessons[num][0], time.Now().Format("2006-01-02"))
			bot.Send(msg)
		}
	}
}

func LessonOne() {
	SendFeedbackLessons(1)
}

func LessonTwo() {
	SendFeedbackLessons(4)
}

func LessonThree() {
	SendFeedbackLessons(6)
}

func LessonSeven() {
	SendFeedbackLessons(7)
}

// функция принимает часы и минуты, сравнивает с текущим временем
// и возвращает true если текущее время меньше или равно заданному
func isTimeBefore(hour, minute int) bool {
	// получаем текущее время
	now := time.Now()

	// создаем объект Time с текущей датой и переданными часами и минутами
	targetTime := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.UTC)

	// сравниваем текущее время с заданным
	return now.Before(targetTime) || now.Equal(targetTime)
}
