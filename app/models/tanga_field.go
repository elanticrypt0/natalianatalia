package models

import (
	"fmt"

	"github.com/k23dev/go4it"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"gorm.io/gorm"
)

type TangaField struct {
	gorm.Model
	Name      string `json:"f_name"`
	Data      string `json:"f_data"`
	TangaUUID string
	Tanga     Tanga `gorm:"foreignKey:TangaUUID;references:UUID"`
}

func NewTangaField() *TangaField {
	return &TangaField{}
}

func (m *TangaField) FindOne(db *gorm.DB, id int) *TangaField {
	var this_model TangaField
	db.First(&this_model, id)
	return &this_model
}

func (m *TangaField) FindAll(db *gorm.DB) *[]TangaField {
	var this_model []TangaField
	db.Order("created_at ASC").Find(&this_model)
	return &this_model
}

func (m *TangaField) Create(db *gorm.DB) *TangaField {
	db.Create(&m)
	return m
}

func (m *TangaField) Update(db *gorm.DB) *TangaField {
	db.Save(&m)
	return m
}

func (m *TangaField) Delete(db *gorm.DB, id int) *TangaField {
	db.Delete(&m)
	return m
}

// Load commond tanga fields
func (m *TangaField) LoadCommodFields(tapp *webcore.TangoApp) *[]TangaField {

	fields := &[]TangaField{}
	if tapp.NNConfig.Tanga_fields_file != "" {
		fields_file := tapp.NNConfig.Tanga_fields_file + ".toml"
		go4it.ReadAndParseToml(fields_file, &fields)
	}
	return fields

}

func (t *TangaField) GetIDAsString() string {
	return fmt.Sprintf("%d", t.ID)
}
