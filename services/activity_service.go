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
