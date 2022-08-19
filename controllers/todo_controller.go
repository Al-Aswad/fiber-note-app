package controllers

import (
	"al-aswad/fiber-note-app/helpers"
	"al-aswad/fiber-note-app/requests"
	"al-aswad/fiber-note-app/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type TodoController interface {
	Create(ctx *fiber.Ctx) error
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
	todoRequest := requests.CreateTodo{}

	errReq := ctx.BodyParser(&todoRequest)
	log.Println("Body Parser ", errReq)

	if errReq != nil {
		res := helpers.BuildErrorResponse("Not Found", errReq.Error(), nil)
		ctx.JSON(res)
		ctx.Status(404)
		return errReq
	}

	result, err := t.todoService.Create(todoRequest)
	if err != nil {
		res := helpers.BuildErrorResponse("Not Found", errReq.Error(), nil)
		ctx.JSON(res)
		ctx.Status(404)
		return err
	}

	res := helpers.BuildResponse("Success", "Success", result)
	ctx.JSON(res)
	ctx.Status(200)

	return nil
}
