package webcore_features

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/k23dev/tango/pkg/webcore"
)

func SetupRoutes(tapp *webcore.TangoApp) {

	// setup
	setup := tapp.Fiber.Group("/setup")

	setup.Get("/", func(c *fiber.Ctx) error {
		return Setup(c, tapp)
	})

	// app monitor
	setup.Get("/monitor", monitor.New(monitor.Config{Title: tapp.App.Config.App_name + " Monitor Page"}))

	//status
	setup.Get("/status", func(c *fiber.Ctx) error {
		return Status(c)
	})

	// seeder
	if tapp.App.Config.App_setup_enabled {
		setup.Get("/seed", func(c *fiber.Ctx) error {
			return Seed(c, tapp)
		})
		setup.Get("/seed/:table_name", func(c *fiber.Ctx) error {
			return Seed(c, tapp)
		})
	}

}
