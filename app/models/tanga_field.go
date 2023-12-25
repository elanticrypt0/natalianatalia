package models

import (
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

func (m *TangaField) FindOne(gas *webcore.GasonlineApp, id int) *TangaField {
	var this_model TangaField
	gas.App.DB.Primary.First(&this_model, id)
	return &this_model
}

func (m *TangaField) FindAll(gas *webcore.GasonlineApp) *[]TangaField {
	var this_model []TangaField
	gas.App.DB.Primary.Order("created_at ASC").Find(&this_model)
	return &this_model
}

func (m *TangaField) Create(gas *webcore.GasonlineApp) *TangaField {
	gas.App.DB.Primary.Create(&m)
	return m
}

func (m *TangaField) Update(gas *webcore.GasonlineApp) *TangaField {
	gas.App.DB.Primary.Save(&m)
	return m
}

func (m *TangaField) Delete(gas *webcore.GasonlineApp, id int) *TangaField {
	gas.App.DB.Primary.Delete(&m)
	return m
}

// Load commond tanga fields
func (m *TangaField) LoadCommodFields(gas *webcore.GasonlineApp) *[]TangaField {

	fields := &[]TangaField{}
	if gas.NNConfig.Tanga_fields_file != "" {
		fields_file := gas.NNConfig.Tanga_fields_file + ".toml"
		go4it.ReadAndParseToml(fields_file, &fields)
	}
	return fields

}
