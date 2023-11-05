package webcore_features

import (
	"github.com/elanticrypt0/natalianatalia/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(gas *webcore.GasonlineApp) {

	// setup
	setup := gas.Fiber.Group("/api/setup")
	if gas.App.Config.App_setup_enabled {
		setup.Get("/", func(c *fiber.Ctx) error {
			// return webcore_features.Setup(c)
			return Setup(c, gas)
		})
	}

	//status
	status := gas.Fiber.Group("/api/status")
	status.Get("/", func(c *fiber.Ctx) error {
		return Status(c)
	})

	// seeder
	seeder := gas.Fiber.Group("/api/seeder")
	if gas.App.Config.App_setup_enabled {
		seeder.Get("/", func(c *fiber.Ctx) error {
			SeedCMD(gas)
			return c.JSON("DB Seeded")
		})
		seeder.Get("/:table_name", func(c *fiber.Ctx) error {
			table := c.Params("table_name", "")
			SeedTable(gas, table)
			return c.JSON("DB Seeded")
		})
	}
}
