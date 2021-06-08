package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Tracker struct {
	ID         uint32    `gorm:"primary_key; auto_increment" json:"id"`
	Weight     int       `json:"weight"`
	ExerciseID uint32    `gorm:"ForeignKey:id" json:"exerciseID"`
	Exercise   Exercise  `json:"exercise"`
	UserID     uint32    `gorm:"ForeignKey:id" json:"userID"`
	User       User      `json:"user"`
	Created    time.Time `gorm:"index"`
}

func (t *Tracker) GetTracker(db *gorm.DB, uid uint32) (*Tracker, error) {
	err := db.Debug().Model(Tracker{}).Where("id = ?", uid).Preload("User").Preload("Exercise").Take(&t).Error
	if err != nil {
		return &Tracker{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Tracker{}, errors.New("Tracker Not Found")
	}

	return t, err
}

func (t *Tracker) CreateTracker(db *gorm.DB) (*Tracker, error) {
	err := db.Debug().Create(&t).Error
	if err != nil {
		return &Tracker{}, err
	}

	return t, nil
}
