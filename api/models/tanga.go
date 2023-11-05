package models

import (
	"github.com/elanticrypt0/natalianatalia/pkg/webcore"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tanga struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"type:uuid;primary_key;"`
	Codename string    `json:"codename"`
}

func NewTanga() *Tanga {
	return &Tanga{}
}

func (m *Tanga) FindOne(gas *webcore.GasonlineApp, id int) *Tanga {
	var this_model Tanga
	gas.App.DB.Primary.First(&this_model, id)
	return &this_model
}

func (m *Tanga) FindAll(gas *webcore.GasonlineApp) *[]Tanga {
	var this_model []Tanga
	gas.App.DB.Primary.Order("created_at ASC").Find(&this_model)
	return &this_model
}

func (m *Tanga) Create(gas *webcore.GasonlineApp) *Tanga {
	gas.App.DB.Primary.Create(&m)
	return m
}

func (m *Tanga) Update(gas *webcore.GasonlineApp) *Tanga {
	gas.App.DB.Primary.Save(&m)
	return m
}

func (m *Tanga) Delete(gas *webcore.GasonlineApp, id int) *Tanga {
	gas.App.DB.Primary.Delete(&m)
	return m
}
