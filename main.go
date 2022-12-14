package main

import (
	"al-aswad/fiber-note-app/config"
	"al-aswad/fiber-note-app/controllers"
	"al-aswad/fiber-note-app/repositories"
	"al-aswad/fiber-note-app/services"
	"fmt"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB = config.DBConnect()
var activityRep repositories.ActivityRepository = repositories.NewActivityRepository(db)
var activityService services.ActivityService = services.NewActivitySerive(activityRep)
var activityController controllers.ActivityController = controllers.NewActivityController(activityService)

var todoRepo repositories.TodoRepository = repositories.NewTodoRepository(db)
var todoService services.TodoService = services.NewTodoService(todoRepo)
var todoController controllers.TodoController = controllers.NewTodoController(todoService)

func main() {
	err := db.Error
	if err != nil {
		fmt.Println("Error Database !")
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Post("/activity-groups", activityController.Create)
	app.Get("/activity-groups", activityController.GetAll)
	app.Get("/activity-groups/:id", activityController.GetOne)
	app.Delete("/activity-groups/:id", activityController.Delete)
	app.Patch("/activity-groups/:id", activityController.Update)

	app.Post("/todo-items", todoController.Create)
	app.Get("/todo-items", todoController.GetAll)
	app.Get("/todo-items/:id", todoController.GetOne)
	app.Patch("/todo-items/:id", todoController.Update)
	app.Delete("/todo-items/:id", todoController.Delete)

	app.Listen(":3000")
}
