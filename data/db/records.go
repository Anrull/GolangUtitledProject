package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Records struct {
	gorm.Model
	Date     string
	Name     string
	Class    string
	Olimps   string
	Stage    string
	Subjects string
	Teachers string
}

var RecordsDB, _ = gorm.Open(sqlite.Open("data/db/records.db"), &gorm.Config{})

func init() {
	RecordsDB.AutoMigrate(&Records{})
}

func AddRecord(name, class, olimp, sub, teacher, stage string) error {
	newUser := Records{
		Date:     time.Now().Format("2006-01-02"),
		Name:     name,
		Class:    class,
		Olimps:   olimp,
		Stage:    stage,
		Subjects: sub,
		Teachers: teacher,
	}
	return RecordsDB.Create(&newUser).Error
}
