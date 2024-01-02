package models

import (
	"fmt"

	"github.com/k23dev/natalianatalia/pkg/tango_errors"
	"gorm.io/gorm"
)

type Tanga_field struct {
	gorm.Model
	Name string `json:"name" param:"name" query:"name" form:"name"`
}

type Tanga_fieldDTO struct {
	Name string `json:"name" param:"name" query:"name" form:"name"`
}

type Tanga_fieldCounter struct {
	Total int
}

func NewTanga_field() *Tanga_field {
	return &Tanga_field{}
}

func (t *Tanga_field) Count(db *gorm.DB) (int, error) {
	counter := &Tanga_fieldCounter{}
	db.Model(&Tanga_field{}).Select("count(ID) as total").Find(&counter)
	return counter.Total, nil
}

func (t *Tanga_field) FindOne(db *gorm.DB, id int) (*Tanga_field, error) {
	var tanga_field Tanga_field
	db.First(&tanga_field, id)
	if tanga_field.ID == 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tanga_field",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &tanga_field, nil
}

func (t *Tanga_field) FindAll(db *gorm.DB) ([]Tanga_field, error) {
	var tanga_fields []Tanga_field
	db.Order("created_at ASC").Find(&tanga_fields)
	if len(tanga_fields) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tanga_field",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return tanga_fields, nil
}

func (t *Tanga_field) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Tanga_field, error) {
	tanga_fields := []Tanga_field{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&tanga_fields)
	if len(tanga_fields) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tanga_field",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &tanga_fields, nil
}

func (t *Tanga_field) Create(db *gorm.DB, name string) (*Tanga_field, error) {
	tanga_field := Tanga_field{
		Name: name,
	}
	db.Create(&tanga_field)
	return &tanga_field, nil
}

func (t *Tanga_field) Update(db *gorm.DB, id int, name string) (*Tanga_field, error) {
	db.Model(&Tanga_field{}).Where("ID =?", id).Update("name", name)
	return t, nil
}

func (t *Tanga_field) Delete(db *gorm.DB, id int) (*Tanga_field, error) {
	tanga_field, err := t.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&tanga_field)
	return tanga_field, nil
}

func (t *Tanga_field) GetIDAsString() string {
	return fmt.Sprintf("%d", t.ID)
}
