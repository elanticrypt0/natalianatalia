package webcore

import (
	"github.com/gofiber/fiber/v2"
)

func SetupStaticRoutes(app *fiber.App) {

	app.Static("/", "./public")
	app.Static("/", "./assets")
	app.Static("/", "./js")
	app.Static("/", "./css")
	app.Static("/", "./images")

}
