package webcore_features

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/k23dev/tango/api/models"
	"github.com/k23dev/tango/pkg/webcore"
)

func Setup(c *fiber.Ctx, tapp *webcore.TangoApp) error {
	automigrateModels(tapp)
	return c.SendString("Setup enabled. Models Migrated.")
}

func SetupOnStartup(tapp *webcore.TangoApp) {
	fmt.Println("\nDatabase automigration...")
	automigrateModels(tapp)
}

func automigrateModels(tapp *webcore.TangoApp) {
	tapp.App.DB.Primary.AutoMigrate(&models.Category{})
}
