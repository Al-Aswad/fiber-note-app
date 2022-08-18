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

	app.Listen(":3000")
}
