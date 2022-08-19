package controllers

import (
	"al-aswad/fiber-note-app/helpers"
	"al-aswad/fiber-note-app/requests"
	"al-aswad/fiber-note-app/services"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ActivityController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
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

	errValidate := requests.ValidateCreateActivity(requestActivity)
	log.Println("validate ", errValidate)
	if errValidate != nil {
		res := helpers.BuildBadRequest("Bad Request", "title cannot be null", struct{}{})
		ctx.JSON(res)
		ctx.Status(400)
		return nil
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
		res := helpers.BuildBadRequest("Not Found", "Activity with ID "+ctx.Params("id")+" Not Found", struct{}{})
		ctx.JSON(res)
		ctx.Status(400)
		return nil
	}

	res := helpers.BuildResponse("Success", "Success", activity)
	ctx.JSON(res)
	ctx.Status(200)

	return nil

}

func (a *activityController) Update(ctx *fiber.Ctx) error {
	archiveUpdate := requests.CreateActivity{}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		res := helpers.BuildErrorResponse("Not Found", err.Error(), nil)
		ctx.JSON(res)
		ctx.Status(404)
		return err
	}

	errBind := ctx.BodyParser(&archiveUpdate)

	if errBind != nil {
		res := helpers.BuildErrorResponse("Not Found", errBind.Error(), nil)
		ctx.JSON(res)
		ctx.Status(404)
		return errBind
	}

	result, status := a.activityServ.UpdateActivity(id, archiveUpdate)

	if !status {
		res := helpers.BuildBadRequest("Not Found", "Activity with ID "+ctx.Params("id")+" Not Found", struct{}{})
		ctx.JSON(res)
		ctx.Status(400)
		return nil
	}

	res := helpers.BuildResponse("Success", "Success", result)
	ctx.JSON(res)
	ctx.Status(200)

	return nil
}

func (a *activityController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		res := helpers.BuildErrorResponse("Error", err.Error(), nil)
		ctx.JSON(res)
		ctx.Status(404)
		return err
	}

	_, hasil := a.activityServ.DeleteActivity(id)
	if !hasil {
		res := helpers.BuildErrorResponse("Activity with ID "+ctx.Params("id")+" Not Found", "not found", helpers.EmptyResponse{})
		ctx.JSON(res)
		ctx.Status(404)
		return err
	}

	res := helpers.BuildResponse("Success", "Success", helpers.EmptyResponse{})
	ctx.JSON(res)
	ctx.Status(200)

	return nil
}
