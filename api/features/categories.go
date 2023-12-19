package features

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/k23dev/natalianatalia/api/models"
	"github.com/k23dev/natalianatalia/pkg/webcore"
)

func FindOneCategory(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	cat := models.NewCategory()
	return c.JSON(cat.FindOne(gas, id))
}

func FindAllCategories(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	cat := models.NewCategory()
	categories := cat.FindAll(gas)
	return c.JSON(categories)
}

func CreateCategory(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	// name comes from json in body
	// cat := new(models.Category)
	cat := models.NewCategory()
	c.BodyParser(&cat)
	category := cat.Create(gas, cat.Name)
	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	cat := models.NewCategory()
	category := cat.FindOne(gas, id)
	c.BodyParser(&cat)
	category = cat.Update(gas, *category)
	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	cat := models.NewCategory()
	category := cat.Delete(gas, id)
	return c.JSON(category)
}
