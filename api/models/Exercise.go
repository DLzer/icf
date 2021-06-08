package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Exercise struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not_null;" json:"name"`
	Sets     uint32 `gorm:"not_null" json:"sets"`
	Reps     uint32 `gorm:"not_null" json:"reps"`
	Category int    `gorm:"not_null" json:"category"`
	Created  time.Time
}

func (e *Exercise) GetExercise(db *gorm.DB, uid uint32) (*Exercise, error) {
	err := db.Debug().Model(Exercise{}).Where("id = ?", uid).Take(&e).Error
	if err != nil {
		return &Exercise{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Exercise{}, errors.New("excercise not found")
	}

	return e, err
}

func (e *Exercise) CreateExercise(db *gorm.DB) (*Exercise, error) {
	err := db.Debug().Create(&e).Error
	if err != nil {
		return &Exercise{}, err
	}
	return e, nil
}
