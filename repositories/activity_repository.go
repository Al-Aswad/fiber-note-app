package repositories

import (
	"al-aswad/fiber-note-app/models"
	"log"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	Create(activity models.Activity) (models.Activity, error)
	GetAll() ([]models.Activity, error)
	GetOne(id int) (models.Activity, error)
	Update(id int, activity models.Activity) (models.Activity, error)
	Delete(id int) bool
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{
		db: db,
	}
}

func (a *activityRepository) Create(activity models.Activity) (models.Activity, error) {
	err := a.db.Save(&activity).Error

	if err != nil {
		return models.Activity{}, err
	}

	return activity, nil
}

func (a *activityRepository) GetAll() ([]models.Activity, error) {
	var activity []models.Activity

	err := a.db.Find(&activity).Error
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (a *activityRepository) GetOne(id int) (models.Activity, error) {
	activity := models.Activity{}

	err := a.db.Debug().First(&activity, "id", id).Error
	if err != nil {
		log.Println("error ", activity)
		log.Println("error ", &err)
		return models.Activity{}, err
	}
	return activity, nil
}

func (a *activityRepository) Update(id int, activity models.Activity) (models.Activity, error) {
	var activityUpdate models.Activity

	err := a.db.Debug().Model(&activityUpdate).Where("id = ?", id).Updates(&activity)
	if err.RowsAffected == 0 {
		log.Println("Update ", err.Error)
		log.Println("Update Effect ", err.RowsAffected)
		return models.Activity{}, err.Error
	}

	return activityUpdate, nil
}

// Refactor
func (a *activityRepository) Delete(id int) bool {
	activity := models.Activity{}

	errFound := a.db.Delete(&activity, id)

	if errFound.RowsAffected == 0 {
		log.Println("Activity cari ", errFound)
		log.Println("Activity cari ", errFound.Error)
		return false
	}

	err := a.db.Debug().Where("id", id).Delete(&activity).Error

	log.Println("Activity ", err)

	if err != nil {
		return false
	}

	return true
}
