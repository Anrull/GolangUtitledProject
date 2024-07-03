package botSchedule

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/bot"
	"awesomeProject/bot/feedback"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"

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
}

func TasksSchedule() {
	users, _ := db.GetAllUsers()

	for _, user := range users {
		var photoByte []byte

		role, _ := db.Get(user.UserID, "role")
		week, _ := timetable.GetWeek(false, true)
		day := timetable.GetDayTomorrow()

		if role == "student" {
			stage, _ := db.Get(user.UserID, "classes")
			lessons, _ := timetable.GetTimetableText(week, day, stage)

			photoByte, _ = timetable.DrawTimetableTest(
				lessons, fmt.Sprintf("%s, нед: %s, день: %s",
					stage, lexicon.Week[week], lexicon.Day[day]),
				false)
		} else {
			teacher, _ := db.Get(user.UserID, "name_teacher")
			lessons, _ := timetable.GetTimetableTeachersText(teacher, week, day)

			photoByte, _ = timetable.DrawTimetableTest(
				lessons, fmt.Sprintf("%s, нед: %s, день: %s", teacher,
					lexicon.Week[week], lexicon.Day[day]),
				true)
		}

		sendPhotoByte(user.UserID, photoByte)
	}
}

func SendFeedbackLessons(num int) {
	num--
	users, err := db.GetAllUsers()
	if err != nil {
		log.Println()
	}

	week, err := timetable.GetWeek(false, true)
	day := timetable.GetDayToday()

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