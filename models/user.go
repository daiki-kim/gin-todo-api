package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

var UserDB *gorm.DB

func InitUserDB() {
	var err error
	UserDB, err = gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	UserDB.AutoMigrate(&User{})
}

func CreateUser(username, password string) (*User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := User{Username: username, Password: string(hashedPassword)}
	result := UserDB.Create(&user)
	return &user, result.Error
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	result := UserDB.Where("username = ?", username).First(&user)
	return &user, result.Error
}

func CheckPassword(user *User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
