package seed

import (
	"log"

	"github.com/DLzer/icf/api/models"
	"github.com/jinzhu/gorm"
)

var workoutOne = models.Workout{
	Name:     "Workout A",
	Exercise: 1250,
}

var workoutTwo = models.Workout{
	Name:     "Workout B",
	Exercise: 1251,
}

var user = models.User{
	Name: "Dillon",
}

var exerciseOne = models.Exercise{
	Name:     "Squat",
	Sets:     5,
	Reps:     5,
	Category: 1250,
}

var exerciseTwo = models.Exercise{
	Name:     "Squat",
	Sets:     5,
	Reps:     5,
	Category: 1251,
}

var tracker = models.Tracker{
	Weight: 150,
}

func Load(db *gorm.DB) {

	// Drop Tracker based on FK Constraint
	err := db.Debug().DropTableIfExists(&models.Tracker{}, &models.Tracker{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	// Workouts
	err = db.Debug().DropTableIfExists(&models.Workout{}, &models.Workout{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Workout{}, &models.Workout{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	err = db.Debug().Model(&models.Workout{}).Create(&workoutOne).Error
	if err != nil {
		log.Fatalf("cannot seed scan table: %v", err)
	}
	err = db.Debug().Model(&models.Workout{}).Create(&workoutTwo).Error
	if err != nil {
		log.Fatalf("cannot seed scan table: %v", err)
	}

	// User
	err = db.Debug().DropTableIfExists(&models.User{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	err = db.Debug().Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed scan table: %v", err)
	}

	var userID = user.ID

	// Exercise
	err = db.Debug().DropTableIfExists(&models.Exercise{}, &models.Exercise{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Exercise{}, &models.Exercise{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	err = db.Debug().Model(&models.Exercise{}).Create(&exerciseOne).Error
	if err != nil {
		log.Fatalf("cannot seed scan table: %v", err)
	}
	err = db.Debug().Model(&models.Exercise{}).Create(&exerciseTwo).Error
	if err != nil {
		log.Fatalf("cannot seed scan table: %v", err)
	}

	var exerciseID = exerciseOne.ID

	// Tracker
	err = db.Debug().AutoMigrate(&models.Tracker{}, &models.Tracker{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	db.Model(&models.Tracker{}).AddForeignKey("exercise_id", "exercises(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Tracker{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	log.Printf("Exercise ID: %d, User ID: %d ...", exerciseID, userID)
	tracker.ExerciseID = exerciseID
	tracker.UserID = userID

	err = db.Debug().Model(&models.Tracker{}).Create(&tracker).Error
	if err != nil {
		log.Fatalf("cannot seed scan table: %v", err)
	}
}
