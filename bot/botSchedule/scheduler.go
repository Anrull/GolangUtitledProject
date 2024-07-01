package botSchedule

import (
	"awesomeProject/backend/timetable"
	"awesomeProject/bot/lexicon"
	"awesomeProject/data/db"

	"awesomeProject/pkg/scheduler"

	"fmt"
)

func RunScheduler() {
	scheduler.NewScheduler(1, 11, 0, TasksSchedule)
	scheduler.NewScheduler(2, 11, 0, TasksSchedule)
	scheduler.NewScheduler(3, 11, 0, TasksSchedule)
	scheduler.NewScheduler(4, 11, 0, TasksSchedule)
	scheduler.NewScheduler(7, 13, 0, TasksSchedule)
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
