package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k23dev/tango/api/features"
	"github.com/k23dev/tango/pkg/webcore"
)

func categoriesRoutes(tapp *webcore.TangoApp, api fiber.Router) {
	categories := api.Group("/categories")
	categories.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllCategories(c, tapp)
	})
	categories.Get("/:id", func(c *fiber.Ctx) error {
		return features.FindOneCategory(c, tapp)
	})

	categories.Post("/", func(c *fiber.Ctx) error {
		return features.CreateCategory(c, tapp)
	})

	categories.Put("/:id", func(c *fiber.Ctx) error {
		return features.UpdateCategory(c, tapp)
	})

	categories.Delete("/:id", func(c *fiber.Ctx) error {
		return features.DeleteCategory(c, tapp)
	})
}
