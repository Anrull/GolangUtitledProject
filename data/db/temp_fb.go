package db

import (
	"log"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var tempDB, _ = gorm.Open(sqlite.Open("data/db/temp.db"), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Silent),
})

type TempFb struct {
	gorm.Model
	Name string `gorm:"column:name"`
}

func init() {
	err := tempDB.AutoMigrate(&TempFb{})
	if err != nil {
		log.Println(err)
	}
}

func AddTempFb(name string) string {
	var tempFb TempFb
 	// Ищем существующую запись с таким же именем
	result := tempDB.Where("name = ?", name).First(&tempFb)
	if result.Error == nil {
		// Если запись найдена, возвращаем ее ID
		return strconv.Itoa(int(tempFb.ID))
	}

	// Если запись не найдена, создаем новую
	newTempFb := TempFb{Name: name}
	result = tempDB.Create(&newTempFb)
	if result.Error != nil {
		log.Println("Ошибка при добавлении в базу данных:", result.Error)
		return ""
	}

	// Возвращаем ID новой записи
	return strconv.Itoa(int(newTempFb.ID))
}

func GetTempFbNameByID(id string) (string, error) {
	var tempFb TempFb
	idInt, err := strconv.Atoi(id)
	if err != nil {
	  	return "", err 
	}
  
	// Ищем запись по ID
	result := tempDB.First(&tempFb, idInt)
	if result.Error != nil {
		return "", result.Error
	}
  
	return tempFb.Name, nil
}