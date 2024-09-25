package db

import (
	"errors"
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

func GetFBLessons(stage string) (*[]FeedbackLesson, error) {
	var records []FeedbackLesson
	query := DBFeedbackLesson.Model(&FeedbackLesson{}).Where("stage = ?", stage)

	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}

	return &records, nil
}

func GetFBLessonsByWeek(stage string, week []string) (*[]FeedbackLesson, error) {
	var records []FeedbackLesson
	if len(week) != 5 {
		return nil, errors.New("week length should be 5")
	}
	query := DBFeedbackLesson.Model(&FeedbackLesson{}).Where(
		"stage = ? AND (date_lesson = ? OR date_lesson = ? OR "+
			"date_lesson = ? OR date_lesson = ? OR date_lesson = ?)",
		stage, week[0], week[1], week[2], week[3], week[4])

	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}

	return &records, nil
}

func GetFBLessonsByWeekTest(stage string, week []string) (*[]FeedbackLesson, error) {
	var records []FeedbackLesson

	// Проверка на пустой массив week
	if len(week) == 0 {
		return nil, errors.New("week cannot be empty")
	}

	// Формирование запроса с использованием IN для работы с любым количеством дней
	query := DBFeedbackLesson.Model(&FeedbackLesson{}).Where(
		"stage = ? AND date_lesson IN ?", stage, week)

	// Выполнение запроса
	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}

	return &records, nil
}
