package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Workout struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null;" json:"workout_name"`
	Exercies uint32 `json:"excercises"`
}

func (w *Workout) GetWorkout(db *gorm.DB, uid uint32) (*Workout, error) {
	err := db.Debug().Model(Workout{}).Where("id = ?", uid).Take(&w).Error
	if err != nil {
		return &Workout{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Workout{}, errors.New("Workout not found")
	}

	return w, err
}
