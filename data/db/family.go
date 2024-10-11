package db

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var FamilyDB, _ = gorm.Open(sqlite.Open("data/db/student.db"), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Silent),
})

type Students struct {
	gorm.Model
	Name  string
	Stage string
	Snils string
}

func Init() {
	FamilyDB.AutoMigrate(&Students{})
}

func CheckSnils(value string) (bool, string, string) {
	var student Students
	result := FamilyDB.First(&student, "snils = ?", hashValue(value))
	if result.Error != nil {
		return false, "", ""
	}
	return true, student.Name, student.Stage
}

func hashValue(value string) string {
	hash := sha256.Sum256([]byte(value))
	return hex.EncodeToString(hash[:])
}

func AddStudent(name, snils, stage string) error {
	f, _, _ := CheckSnils(snils)
	if !f {
		return FamilyDB.Create(&Students{Name: name, Snils: hashValue(snils), Stage: stage}).Error
	}
	return fmt.Errorf("student with snils %s already exists", snils)
}