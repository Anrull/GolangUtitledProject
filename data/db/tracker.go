package db

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var TrackerDB, _ = gorm.Open(sqlite.Open("data/db/tracker.db"), &gorm.Config{})

type Tracker struct {
	gorm.Model
	UserID       int64  `gorm:"column:user_id"`
	Time         string `gorm:"column:time"`
	Name         string `gorm:"column:name"`
	Stage        string `gorm:"column:stage;default=''"`
	Snils        int64  `gorm:"column:snils"`
	Olimps       string `gorm:"column:olimps"`
	LastIdOlimps string `gorm:"column:last_id_olimps"`
	DeleteOlimps string `gorm:"column:delete_olimps"`
}

func init() {
	TrackerDB.AutoMigrate(&Tracker{})
}

func AddTracker(message *tgbotapi.Message, column, value string) error {
	var tracker Tracker
	result := TrackerDB.First(&tracker, "user_id = ?", message.Chat.ID)
	if result.Error != nil {
		return result.Error
	}
	switch column {
	case "delete_olimps":
		tracker.DeleteOlimps = value
	case "olimps":
		if tracker.Olimps == "" {
			tracker.Olimps = value
		} else {
			tracker.Olimps = tracker.Olimps + ";" + value
		}
	}
	return TrackerDB.Save(&tracker).Error
}

func CreateNewTrackerUser(message *tgbotapi.Message, name, stage string) error {
	newUser := Tracker{
		UserID: message.Chat.ID,
		Name:   name,
		Stage:  stage,
	}
	return TrackerDB.Create(&newUser).Error
}
