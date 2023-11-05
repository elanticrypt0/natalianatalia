package models

import (
	"github.com/elanticrypt0/natalianatalia/pkg/webcore"
	"gorm.io/gorm"
)

type Scripts struct {
	gorm.Model
	Name string `json:"name"`
}

func NewScripts() *Scripts {
	return &Scripts{}
}

func (m *Scripts) FindOne(gas *webcore.GasonlineApp, id int) *Scripts {
	var this_model Scripts
	gas.App.DB.Primary.First(&this_model, id)
	return &this_model
}

func (m *Scripts) FindAll(gas *webcore.GasonlineApp) *[]Scripts {
	var this_model []Scripts
	gas.App.DB.Primary.Order("created_at ASC").Find(&this_model)
	return &this_model
}

func (m *Scripts) Create(gas *webcore.GasonlineApp) *Scripts {
	gas.App.DB.Primary.Create(&m)
	return m
}

func (m *Scripts) Update(gas *webcore.GasonlineApp) *Scripts {
	gas.App.DB.Primary.Save(&m)
	return m
}

func (m *Scripts) Delete(gas *webcore.GasonlineApp, id int) *Scripts {
	gas.App.DB.Primary.Delete(&m)
	return m
}
