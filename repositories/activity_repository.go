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
	Update()
	Delete()
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
		return models.Activity{}, err
	}
	return activity, nil
}

func (a *activityRepository) Update() {

}
func (a *activityRepository) Delete() {

}
