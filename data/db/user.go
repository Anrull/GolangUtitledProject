package db

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB, _ = gorm.Open(sqlite.Open("data/db/users.db"), &gorm.Config{})

type User struct {
	gorm.Model
	UserID      int64  `gorm:"column:user_id"`
	Name        string `gorm:"column:name"`
	Username    string `gorm:"column:username"`
	Classes     string `gorm:"column:classes;default=''"`
	Role        string `gorm:"column:role;default=''"`
	NameTeacher string `gorm:"column:name_teacher;default=''"`
	Newsletter  string `gorm:"column:newsletter;default='1'"`
	Temp        string `gorm:"column:temp;default=''"`
	Bot         string `gorm:"column:bot;default=''"`
	LastIDs     string `gorm:"column:last_ids;default=''"`
	Admin       string `gorm:"column:admin;default='0'"`
}

func init() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
	}
}

func NewUser(message tgbotapi.Message) error {
	userID := message.Chat.ID
	name := fmt.Sprintf("%s %s", message.From.FirstName, message.From.LastName)
	username := message.From.UserName

	var user User
	result := DB.First(&user, "user_id = ?", userID)

	if result.Error == nil { // Пользователь найден, обновляем данные
		user.Name = name
		user.Username = username
		return DB.Save(&user).Error
	} else if result.Error == gorm.ErrRecordNotFound { // Пользователь не найден, создаем нового
		newUser := User{
			UserID:     userID,
			Name:       name,
			Username:   username,
			Bot:        "bot-schedule",
			Newsletter: "1",
		}
		return DB.Create(&newUser).Error
	} else {
		return result.Error
	}
}

func Get(ChatID int64, column string) (string, error) {
	var user User
	result, res := DB.First(&user, "user_id = ?", ChatID), ""
	if result.Error != nil {
		return "", result.Error
	}

	switch column {
	case "name":
		res = user.Name
	case "classes":
		res = user.Classes
	case "role":
		res = user.Role
	case "newsletter":
		res = user.Newsletter
	case "bot":
		res = user.Bot
	case "last_ids":
		res = user.LastIDs
	case "username":
		res = user.Username
	case "name_teacher":
		res = user.NameTeacher
	case "temp":
		res = user.Temp
	}
	return res, nil
}

func GetUserCount() (int64, error) {
	var count int64
	result := DB.Model(&User{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetInfoAboutPerson(userID int64) (User, error) {
	var user User
	result := DB.First(&user, "user_id = ?", userID)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func Update(ChatID int64, column, value string) error {
	var user User
	result := DB.First(&user, "user_id = ?", ChatID)
	if result.Error != nil {
		return result.Error
	}

	switch column {
	case "name":
		user.Name = value
	case "classes":
		user.Classes = value
	case "role":
		user.Role = value
	case "newsletter":
		user.Newsletter = value
	case "bot":
		user.Bot = value
	case "last_ids":
		user.LastIDs = value
	case "username":
		user.Username = value
	case "name_teacher":
		user.NameTeacher = value
	case "temp":
		user.Temp = value
	}
	return DB.Save(&user).Error
}

//func main() {
//	fmt.Println(Get(566, "username"))
//}
