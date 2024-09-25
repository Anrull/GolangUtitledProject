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
		if value != "" {
			if tracker.Olimps == "" {
				tracker.Olimps = value
			} else {
				tracker.Olimps = tracker.Olimps + ";" + value
			}
		} else {
			tracker.Olimps = ""
		}
	case "get_olimps":
		tracker.LastIdOlimps = value
	case "filter":
		if value != "" {
			if tracker.LastIdOlimps == "" {
				tracker.LastIdOlimps = value
			} else {
				tracker.LastIdOlimps = tracker.LastIdOlimps + ";;" + value
			}
		} else {
			tracker.LastIdOlimps = ""
		}
	}
	return TrackerDB.Save(&tracker).Error
}

func GetTracker(message *tgbotapi.Message, column string) (string, error) {
	var tracker Tracker
	result := TrackerDB.First(&tracker, "user_id = ?", message.Chat.ID)
	if result.Error != nil {
		return "", result.Error
	}
	var res string
	switch column {
	case "delete_olimps":
		res = tracker.DeleteOlimps
	case "olimps":
		res = tracker.Olimps
	case "name":
		res = tracker.Name
	case "stage":
		res = tracker.Stage
	case "filter":
		res = tracker.LastIdOlimps
	}
	return res, nil
}

func CreateNewTrackerUser(message *tgbotapi.Message, name, stage string) error {
	_, err := GetTracker(message, "name")
	if err != nil {
		newUser := Tracker{
			UserID: message.Chat.ID,
			Name:   name,
			Stage:  stage,
		}
		return TrackerDB.Create(&newUser).Error
	}
	return nil
}

func GetInfoAboutPersonTracker(userID int64) (Tracker, error) {
	var user Tracker
	result := TrackerDB.First(&user, "user_id = ?", userID)
	if result.Error != nil {
		return Tracker{}, result.Error
	}
	return user, nil
}
