package main

import (
	"al-aswad/fiber-note-app/config"
	"fmt"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB = config.DBConnect()

func main() {
	err := db.Error
	if err != nil {
		fmt.Println("Error Database !")
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
