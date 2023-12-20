package webcore_features

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/k23dev/go4it"
	"github.com/k23dev/tango/api/models"
	"github.com/k23dev/tango/pkg/webcore"
)

const seedDir = "./seeds/"

func Seed(c *fiber.Ctx, gas *webcore.TangoApp) error {
	seedCategories(gas)
	return c.JSON("OK")
}

func seedCategories(gas *webcore.TangoApp) {
	cat_list := []models.Category{}
	go4it.ReadAndParseJson(seedDir+"categories", &cat_list)
	// for _, category := range cat_list {
	// 	models.CreateCategory(gas, category.Name)
	// }
	gas.App.DB.Primary.Save(&cat_list)
	log.Println("Categories seeded")
}
