package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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
	if err := RecordsDB.AutoMigrate(&Records{}); err != nil {
		log.Println("Не удалось создать БД")
	}
}

func AddRecord(name, class, olimp, sub, teacher, stage string) error {
	var record Records
	result := DB.First(&record, "name = ?", name)
	if result.Error == nil {
		return nil
	}

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

func GetRecords(name, sub, olimp, stage, teacher string) (*[]Records, error) {
	var records []Records
	query := getQuery(name, sub, olimp, stage, teacher)

	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}

	return &records, nil
}

func GetRecordsCount(name, sub, olimp, stage, teacher string) (int, error) {
	var count int64
	query := getQuery(name, sub, olimp, stage, teacher)

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func getQuery(name, sub, olimp, stage, teacher string) *gorm.DB {
	query := RecordsDB.Model(&Records{}).Where("name = ?", name)

	if sub != "nil" {
		query = query.Where("subjects LIKE ?", "%"+sub+"%")
	}
	if olimp != "nil" {
		query = query.Where("olimps LIKE ?", "%"+olimp+"%")
	}
	if stage != "nil" {
		query = query.Where("stage LIKE ?", "%"+stage+"%")
	}
	if teacher != "nil" {
		query = query.Where("teachers LIKE ?", "%"+teacher+"%")
	}

	return query
}
