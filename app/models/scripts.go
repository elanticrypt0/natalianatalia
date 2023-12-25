package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Scripts struct {
	gorm.Model
	Name string `json:"name"`
}

func NewScripts() *Scripts {
	return &Scripts{}
}

func (m *Scripts) FindOne(db *gorm.DB, id int) *Scripts {
	var this_model Scripts
	db.First(&this_model, id)
	return &this_model
}

func (m *Scripts) FindAll(db *gorm.DB) *[]Scripts {
	var this_model []Scripts
	db.Order("created_at ASC").Find(&this_model)
	return &this_model
}

func (m *Scripts) Create(db *gorm.DB) *Scripts {
	db.Create(&m)
	return m
}

func (m *Scripts) Update(db *gorm.DB) *Scripts {
	db.Save(&m)
	return m
}

func (m *Scripts) Delete(db *gorm.DB, id int) *Scripts {
	db.Delete(&m)
	return m
}

func (s *Scripts) GetIDAsString() string {
	return fmt.Sprintf("%d", s.ID)
}
