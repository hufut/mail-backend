package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model

	username string `json:"username"`
	password string `json:"password"`
	phone    string `json:"phone"`
	email    string `json:"email"`
}

func getUser(id int) (*User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}
