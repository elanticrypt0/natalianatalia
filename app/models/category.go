package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name" param:"name" query:"name" form:"name"`
}

type CategoryDTO struct {
	Name string `json:"name" param:"name" query:"name" form:"name"`
}

func NewCategory() *Category {
	return &Category{}
}

func (c *Category) FindOne(db *gorm.DB, id int) *Category {
	var category Category
	db.First(&category, id)
	return &category
}

func (c *Category) FindAll(db *gorm.DB) []Category {
	var categories []Category
	// TODO revisar esta parte
	// gasonline.App.DB.Primary.Order("created_at ASC").Find(&categories)
	db.Order("created_at ASC").Find(&categories)
	return categories
}

func (c *Category) Create(db *gorm.DB, name string) *Category {
	category := Category{
		Name: name,
	}
	db.Create(&category)
	return &category
}

func (c *Category) Update(db *gorm.DB, category Category) *Category {
	db.Save(&category)
	return &category
}

func (c *Category) Delete(db *gorm.DB, id int) *Category {
	var category Category
	db.First(&category, id)
	db.Delete(&category)
	return &category
}

func (c *Category) GetIDAsString(id uint) string {
	return fmt.Sprintf("%d", id)
}
