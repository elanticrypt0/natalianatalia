package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tanga struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"type:uuid;primary_key;"`
	Codename string    `json:"codename"`
}

type TangaDTO struct {
	Codename string `json:"codename" param:"codename" query:"codename" form:"codename"`
}

func NewTanga() *Tanga {
	return &Tanga{}
}

func (m *Tanga) FindOne(db *gorm.DB, id int) *Tanga {
	var this_model Tanga
	db.First(&this_model, id)
	return &this_model
}

func (m *Tanga) FindAll(db *gorm.DB) *[]Tanga {
	var this_model []Tanga
	db.Order("created_at ASC").Find(&this_model)
	return &this_model
}

func (m *Tanga) Create(db *gorm.DB, codename string) *Tanga {
	m.Codename = codename
	db.Create(&m)
	return m
}

func (m *Tanga) Update(db *gorm.DB, tanga *Tanga) *Tanga {
	db.Save(&tanga)
	return m
}

func (m *Tanga) Delete(db *gorm.DB, id int) *Tanga {
	db.Delete(&m)
	return m
}

func (t *Tanga) GetIDAsString() string {
	return fmt.Sprintf("%d", t.ID)
}
