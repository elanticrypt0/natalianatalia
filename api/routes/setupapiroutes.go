package routes

import (
	"github.com/k23dev/tango/pkg/webcore"
)

func SetupApiRoutes(tapp *webcore.TangoApp) {
	api := tapp.Fiber.Group("/api")
	// categories
	categoriesRoutes(tapp, api)
}
