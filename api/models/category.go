package models

import (
	"strings"

	"github.com/k23dev/natalianatalia/pkg/webcore"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

func NewCategory() *Category {
	return &Category{}
}

func (c *Category) FindOne(gas *webcore.GasonlineApp, id int) *Category {
	var category Category
	gas.App.DB.Primary.First(&category, id)
	return &category
}

func (c *Category) FindAll(gas *webcore.GasonlineApp) []Category {
	var categories []Category
	// TODO revisar esta parte
	// gasonline.App.DB.Primary.Order("created_at ASC").Find(&categories)
	gas.App.DB.Primary.Order("created_at ASC").Find(&categories)
	return categories
}

func (c *Category) Create(gas *webcore.GasonlineApp, name string) *Category {
	category := Category{
		Name: strings.ToLower(name),
	}
	gas.App.DB.Primary.Create(&category)
	return &category
}

func (c *Category) Update(gas *webcore.GasonlineApp, category Category) *Category {
	gas.App.DB.Primary.Save(&category)
	return &category
}

func (c *Category) Delete(gas *webcore.GasonlineApp, id int) *Category {
	var category Category
	gas.App.DB.Primary.First(&category, id)
	gas.App.DB.Primary.Delete(&category)
	return &category
}
