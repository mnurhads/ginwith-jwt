package models

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name		string `json:"name" gorm:"type:varchar(50)"`
	Username	string `json:"username" gorm:"type:varchar(20);unique"`
	Email		string `json:"email" gorm:"type:varchar(50)"`
	Password	string `json:"password"`
	UserStatus  string `json:"userstatus" gorm:type:varchar(10)`
}

// declare hash
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))

	if err != nil {
		return err
	}

	return nil
}