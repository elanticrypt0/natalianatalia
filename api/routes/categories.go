package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k23dev/natalianatalia/api/features"
	"github.com/k23dev/natalianatalia/pkg/webcore"
)

func categoriesRoutes(gas *webcore.GasonlineApp, api fiber.Router) {
	categories := api.Group("/categories")
	categories.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllCategories(c, gas)
	})
	categories.Get("/:id", func(c *fiber.Ctx) error {
		return features.FindOneCategory(c, gas)
	})

	categories.Post("/", func(c *fiber.Ctx) error {
		return features.CreateCategory(c, gas)
	})

	categories.Put("/:id", func(c *fiber.Ctx) error {
		return features.UpdateCategory(c, gas)
	})

	categories.Delete("/:id", func(c *fiber.Ctx) error {
		return features.DeleteCategory(c, gas)
	})
}
