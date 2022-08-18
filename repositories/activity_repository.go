package repositories

import (
	"al-aswad/fiber-note-app/models"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	Create(activity models.Activity) (models.Activity, error)
	GetAll()
	GetOne()
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

func (a *activityRepository) GetAll() {

}

func (a *activityRepository) GetOne() {

}

func (a *activityRepository) Update() {

}
func (a *activityRepository) Delete() {

}
