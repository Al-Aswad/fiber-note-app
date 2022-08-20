package controllers

import (
	"al-aswad/fiber-note-app/helpers"
	"al-aswad/fiber-note-app/requests"
	"al-aswad/fiber-note-app/services"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TodoController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type todoController struct {
	todoService services.TodoService
}

func NewTodoController(todoService services.TodoService) TodoController {
	return &todoController{
		todoService: todoService,
	}
}

func (t *todoController) Create(ctx *fiber.Ctx) error {

	var todoRequest requests.CreateTodo

	if err := ctx.BodyParser(&todoRequest); err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errValidate := requests.ValidateCreateTodo(todoRequest)
	if errValidate != nil {
		res := helpers.BuildBadRequest("Bad Request", "title cannot be null", struct{}{})
		ctx.JSON(res)
		ctx.Status(400)
		return nil
	}

	result, err := t.todoService.Create(todoRequest)
	if err != nil {
		res := helpers.BuildErrorResponse("Not Found", err.Error(), nil)
		ctx.JSON(res)
		ctx.Status(404)
		return err
	}

	res := helpers.BuildResponse("Success", "Success", result)
	ctx.JSON(res)
	ctx.Status(200)

	return nil
}

func (t *todoController) GetAll(ctx *fiber.Ctx) error {
	todo, err := t.todoService.GetAll()

	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res := helpers.BuildResponse("Success", "Success", todo)
	ctx.JSON(res)
	ctx.Status(200)

	return nil
}

func (t *todoController) Update(ctx *fiber.Ctx) error {
	todoUpdate := requests.UpdateTodo{}

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		res := helpers.BuildBadRequest("Bad Request", "ID not Valid", struct{}{})
		ctx.JSON(res)
		ctx.Status(400)
		return nil
	}

	errBind := ctx.BodyParser(&todoUpdate)

	if errBind != nil {
		res := helpers.BuildBadRequest("Bad Request", "Data not Valid", struct{}{})
		ctx.JSON(res)
		ctx.Status(400)
		return nil
	}

	errValidate := requests.ValidateUpdateTodo(todoUpdate)
	if errValidate != nil {
		res := helpers.BuildBadRequest("Bad Request", "title cannot be null", struct{}{})
		ctx.JSON(res)
		ctx.Status(400)
		return nil
	}

	result, status := t.todoService.Update(id, todoUpdate)

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

func (t *todoController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		res := helpers.BuildBadRequest("Bad Request", "ID not Valid", struct{}{})
		ctx.JSON(res)
		ctx.Status(400)
		return nil
	}

	status, _ := t.todoService.Delete(id)
	log.Println("Status ", status)
	if !status {
		res := helpers.BuildBadRequest("Not Found", "Activity with ID "+ctx.Params("id")+" Not Found", struct{}{})
		ctx.JSON(res)
		ctx.Status(400)
		return nil
	}

	res := helpers.BuildResponse("Success", "Success", helpers.EmptyResponse{})
	ctx.JSON(res)
	ctx.Status(200)

	return nil

}
