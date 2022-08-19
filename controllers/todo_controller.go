package controllers

import (
	"al-aswad/fiber-note-app/helpers"
	"al-aswad/fiber-note-app/requests"
	"al-aswad/fiber-note-app/services"

	"github.com/gofiber/fiber/v2"
)

type TodoController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
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

	todoRequest := new(requests.CreateTodo)

	if err := ctx.BodyParser(todoRequest); err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result, err := t.todoService.Create(*todoRequest)
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
