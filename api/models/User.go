package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID      uint32 `gorm:"primary_key; auto_increment" json:"id"`
	Name    string `gorm:"not_null;" json:"name"`
	Created time.Time
}

func (u *User) GetUser(db *gorm.DB, uid uint32) (*User, error) {
	err := db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}

	return u, err
}

func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
