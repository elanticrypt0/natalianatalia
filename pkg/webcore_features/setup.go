package webcore_features

import (
	"github.com/elanticrypt0/natalianatalia/api/models"
	"github.com/elanticrypt0/natalianatalia/pkg/webcore"
	"github.com/gofiber/fiber/v2"
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
