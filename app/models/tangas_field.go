package models

import (
	"fmt"

	"github.com/k23dev/natalianatalia/pkg/tango_errors"
	"gorm.io/gorm"
)

type Tangas_field struct {
	gorm.Model
	Name    string `json:"name" param:"name" query:"name" form:"name"`
	FValue  string `json:"fvalue" param:"fvalue" query:"fvalue" form:"fvalue"`
	TangaID int
	Tanga   Tanga
}

type Tangas_fieldDTO struct {
	Name    string `json:"name" param:"name" query:"name" form:"name"`
	FValue  string `json:"fvalue" param:"fvalue" query:"fvalue" form:"fvalue"`
	TangaID int    `json:"tanga_id" param:"tanga_id" query:"tanga_id" form:"tanga_id"`
}

type Tangas_fieldCounter struct {
	Total int
}

func NewTangas_field() *Tangas_field {
	return &Tangas_field{}
}

func (t *Tangas_field) Count(db *gorm.DB) (int, error) {
	counter := &Tangas_fieldCounter{}
	db.Model(&Tangas_field{}).Select("count(ID) as total").Find(&counter)
	return counter.Total, nil
}

func (t *Tangas_field) FindOne(db *gorm.DB, id int) (*Tangas_field, error) {
	var tangas_field Tangas_field
	db.First(&tangas_field, id)
	if tangas_field.ID == 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tangas_field",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &tangas_field, nil
}

func (t *Tangas_field) FindAll(db *gorm.DB) ([]Tangas_field, error) {
	var tangas_fields []Tangas_field
	db.Order("created_at ASC").Find(&tangas_fields)
	if len(tangas_fields) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tangas_field",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return tangas_fields, nil
}

func (t *Tangas_field) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Tangas_field, error) {
	tangas_fields := []Tangas_field{}

	db.Order("created_at ASC").Preload("Tanga").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&tangas_fields)
	if len(tangas_fields) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tangas_field",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &tangas_fields, nil
}

func (t *Tangas_field) Create(db *gorm.DB, name, fvalue string, tangaID int) (*Tangas_field, error) {
	tangas_field := Tangas_field{
		Name:    name,
		FValue:  fvalue,
		TangaID: tangaID,
	}
	db.Create(&tangas_field)
	return &tangas_field, nil
}

func (t *Tangas_field) Update(db *gorm.DB, id int, name, fvalue string, tangaID int) (*Tangas_field, error) {
	db.Model(&Tangas_field{}).Where("ID =?", id).Update("name", name).Update("f_value", fvalue).Update("tanga_id", tangaID)
	return t, nil
}

func (t *Tangas_field) Delete(db *gorm.DB, id int) (*Tangas_field, error) {
	tangas_field, err := t.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&tangas_field)
	return tangas_field, nil
}

func (t *Tangas_field) GetIDAsString() string {
	return fmt.Sprintf("%d", t.ID)
}
