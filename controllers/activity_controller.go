package controllers

import (
	"al-aswad/fiber-note-app/helpers"
	"al-aswad/fiber-note-app/requests"
	"al-aswad/fiber-note-app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ActivityController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
}

type activityController struct {
	activityServ services.ActivityService
}

func NewActivityController(activityServ services.ActivityService) ActivityController {
	return &activityController{
		activityServ: activityServ,
	}
}

func (a *activityController) Create(ctx *fiber.Ctx) error {
	var requestActivity requests.CreateActivity

	errReq := ctx.BodyParser(&requestActivity)
	if errReq != nil {
		res := helpers.BuildErrorResponse("Not Found", errReq.Error(), nil)
		ctx.JSON(res)
		ctx.Status(404)
		return errReq
	}

	createActivity, err := a.activityServ.CreateActivity(requestActivity)
	if err != nil {
		res := helpers.BuildErrorResponse("Not Found", errReq.Error(), nil)
		ctx.JSON(res)
		ctx.Status(404)
		return err
	}

	res := helpers.BuildResponse("Success", "Success", createActivity)
	ctx.JSON(res)
	ctx.Status(200)

	return nil
}

func (a *activityController) GetAll(ctx *fiber.Ctx) error {
	activity, err := a.activityServ.GetAllActivity()

	if err != nil {
		res := helpers.BuildErrorResponse("Not Found", err.Error(), nil)
		ctx.JSON(res)
		ctx.Status(404)
		return err
	}

	res := helpers.BuildResponse("Success", "Success", activity)
	ctx.JSON(res)
	ctx.Status(200)

	return nil
}

func (a *activityController) GetOne(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		res := helpers.BuildErrorResponse("Not Found", err.Error(), nil)
		ctx.JSON(res)
		ctx.Status(404)
		return err
	}

	activity, err := a.activityServ.GetOne(id)
	if err != nil {
		res := helpers.BuildErrorResponse("Not Found", err.Error(), nil)
		ctx.JSON(res)
		ctx.Status(404)
		return err
	}

	res := helpers.BuildResponse("Success", "Success", activity)
	ctx.JSON(res)
	ctx.Status(200)

	return nil

}
