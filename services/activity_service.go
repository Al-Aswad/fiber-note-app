package services

import (
	"al-aswad/fiber-note-app/models"
	"al-aswad/fiber-note-app/repositories"
	"al-aswad/fiber-note-app/requests"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type ActivityService interface {
	CreateActivity(activity requests.CreateActivity) (interface{}, error)
	GetAllActivity() ([]models.Activity, error)
	GetOne(id int) (models.Activity, error)
	DeleteActivity(id int) (models.Activity, bool)
	UpdateActivity(id int, activity requests.CreateActivity) (models.Activity, bool)
}

type ActivityServiceImpl struct {
	// db *gorm.DB
	activityRepo repositories.ActivityRepository
}

func NewActivitySerive(activityRepo repositories.ActivityRepository) ActivityService {
	return &ActivityServiceImpl{
		activityRepo: activityRepo,
	}
}

func (a *ActivityServiceImpl) CreateActivity(activity requests.CreateActivity) (interface{}, error) {
	activityCreate := models.Activity{}

	err := smapping.FillStruct(&activityCreate, smapping.MapFields(&activity))

	if err != nil {
		fmt.Println("[ActivityServiceImpl.Create] error fill struct", err)
		return activityCreate, err
	}

	activityCreate, err = a.activityRepo.Create(activityCreate)

	if err != nil {
		log.Println("[ActivityServiceImpl.Create] error execute query ", err)
	}

	return activityCreate, nil
}

func (a *ActivityServiceImpl) GetAllActivity() ([]models.Activity, error) {
	activity, err := a.activityRepo.GetAll()
	fmt.Println("activity ", activity)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (a *ActivityServiceImpl) GetOne(id int) (models.Activity, error) {
	activity, err := a.activityRepo.GetOne(id)
	if err != nil {
		log.Println("Activity Service ", err)
		return models.Activity{}, err
	}

	return activity, nil
}

func (a *ActivityServiceImpl) UpdateActivity(id int, activity requests.CreateActivity) (models.Activity, bool) {
	activityUpdate := models.Activity{}

	err := smapping.FillStruct(&activityUpdate, smapping.MapFields(&activity))

	if err != nil {
		log.Println("[ActivityServiceImpl.Create] error fill struct", err)
		return activityUpdate, false
	}

	activityUpdate, status := a.activityRepo.Update(id, activityUpdate)

	if !status {
		log.Println("Service Update Activity ", err)
		return models.Activity{}, status
	}

	return activityUpdate, true

}

func (a *ActivityServiceImpl) DeleteActivity(id int) (models.Activity, bool) {
	var activity models.Activity
	hasil := a.activityRepo.Delete(id)

	if !hasil {
		return activity, false

	}

	return activity, true
}
