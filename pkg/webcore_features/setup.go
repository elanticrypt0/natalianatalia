package webcore_features

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k23dev/natalianatalia/api/models"
	"github.com/k23dev/natalianatalia/pkg/webcore"
)

func Setup(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	migrateModels(gas)
	return c.SendString("Setup enabled. Models Migrated.")
}

func SetupOnStartup(gas *webcore.GasonlineApp) {
	migrateModels(gas)
}

func migrateModels(gas *webcore.GasonlineApp) {
	gas.App.DB.Primary.AutoMigrate(&models.Category{})
}
