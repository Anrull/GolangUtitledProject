package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBFeedbackLesson, _ = gorm.Open(sqlite.Open("data/db/feedback.db"), &gorm.Config{})

type FeedbackLesson struct {
	gorm.Model
	UserName    string
	Stage       string
	UserID      int64
	DateLesson  string
	TitleLesson string
	Status      string
}

func init() {
	DBFeedbackLesson.AutoMigrate(&FeedbackLesson{})
}

func CreateFBLessons(UserID int64, UserName, Stage, Title, Date, Status string) error {
	newFBLesson := FeedbackLesson{
		UserName:    UserName,
		UserID:      UserID,
		DateLesson:  Date,
		TitleLesson: Title,
		Status:      Status,
		Stage:       Stage,
	}
	return DBFeedbackLesson.Create(&newFBLesson).Error
}
