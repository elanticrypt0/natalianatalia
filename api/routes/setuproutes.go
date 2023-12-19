package routes

import (
	"github.com/k23dev/natalianatalia/pkg/webcore"
)

func SetupFeaturesRoutes(gas *webcore.GasonlineApp) {
	api := gas.Fiber.Group("/api")
	// categories
	categoriesRoutes(gas, api)
}
