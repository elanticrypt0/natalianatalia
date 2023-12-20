package webcore

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func MiddlewareSetup(gas *TangoApp) {
	// CORS
	gas.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: gas.App.Config.App_CORS_origins,
		AllowHeaders: gas.App.Config.App_CORS_headers,
	}))

	//  Recover from error
	gas.Fiber.Use(recover.New())

	// LoggerOnFile(gas.Fiber)
	LogOn(gas.Fiber)

}
