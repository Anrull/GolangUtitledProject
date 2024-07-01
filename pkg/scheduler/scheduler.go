package scheduler

import (
	"time"
)

type Time struct {
	Year   int
	Month  time.Month
	Day    int
	Hour   int
	Minute int
	Second int
	NSec   int
	Locale *time.Location
}

// NewScheduler
//
// The function takes arguments: day, hour, minute and func. Every week it performs the transferred function at the specified time
func NewScheduler(day, hour, minute int, do func()) {
	go func() {
		for {
			now := time.Now()
			nextRun := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.UTC)

			// Находим следующий нужный день недели
			daysUntilTarget := (day - int(now.Weekday()) + 7) % 7
			if daysUntilTarget == 0 && now.Hour() >= hour && now.Minute() >= minute {
				daysUntilTarget = 7 // Запуск на следующей неделе, если уже прошло время запуска на этой неделе
			}
			nextRun = nextRun.AddDate(0, 0, daysUntilTarget)

			time.Sleep(nextRun.Sub(now))
			do()
			time.Sleep(24 * time.Hour) // Ожидание до следующего дня
		}
	}()
}

func (t *Time) NewSchedulerV2(do func()) {
	go func() {
		for {
			now := time.Now()

			nextRun := time.Date(t.Year, t.Month, t.Day, t.Hour,
				t.Minute, t.Second, t.NSec, t.Locale)
			if now.After(nextRun) {
				nextRun = nextRun.Add(24 * time.Hour)
			}

			time.Sleep(nextRun.Sub(now))
			do()
		}
	}()
}
