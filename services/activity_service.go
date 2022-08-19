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
	UpdateActivity(id int, activity requests.CreateActivity) (models.Activity, error)
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
	errValidate := requests.ValidateCreateActivity(activity)
	if errValidate != nil {
		log.Println("[ActivityServiceImpl.Create] Validate error ", errValidate)
		return errValidate, nil
	}

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
		return models.Activity{}, err
	}

	return activity, nil
}

func (a *ActivityServiceImpl) UpdateActivity(id int, activity requests.CreateActivity) (models.Activity, error) {
	activityUpdate := models.Activity{}

	err := smapping.FillStruct(&activityUpdate, smapping.MapFields(&activity))

	if err != nil {
		log.Println("[ActivityServiceImpl.Create] error fill struct", err)
		return activityUpdate, err
	}

	activityUpdate, err = a.activityRepo.Update(id, activityUpdate)

	if err != nil {
		return activityUpdate, err
	}

	return activityUpdate, nil

}

func (a *ActivityServiceImpl) DeleteActivity(id int) (models.Activity, bool) {
	var activity models.Activity
	hasil := a.activityRepo.Delete(id)

	if !hasil {
		return activity, false

	}

	return activity, true
}
