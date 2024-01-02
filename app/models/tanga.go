package models

import (
	"fmt"

	"github.com/k23dev/natalianatalia/pkg/tango_errors"
	"gorm.io/gorm"
)

type Tanga struct {
	gorm.Model
	Name    string `json:"name" param:"name" query:"name" form:"name"`
	Comment string `json:"comment" param:"comment" query:"comment" form:"comment"`
}

type TangaDTO struct {
	Name    string `json:"name" param:"name" query:"name" form:"name"`
	Comment string `json:"comment" param:"comment" query:"comment" form:"comment"`
}

type TangaCounter struct {
	Total int
}

func NewTanga() *Tanga {
	return &Tanga{}
}

func (t *Tanga) Count(db *gorm.DB) (int, error) {
	counter := &TangaCounter{}
	db.Model(&Tanga{}).Select("count(ID) as total").Find(&counter)
	return counter.Total, nil
}

func (t *Tanga) FindOne(db *gorm.DB, id int) (*Tanga, error) {
	var tanga Tanga
	db.First(&tanga, id)
	if tanga.ID == 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tanga",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &tanga, nil
}

func (t *Tanga) FindAll(db *gorm.DB) ([]Tanga, error) {
	var tangas []Tanga
	db.Order("created_at ASC").Find(&tangas)
	if len(tangas) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tanga",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return tangas, nil
}

func (t *Tanga) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Tanga, error) {
	tangas := []Tanga{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&tangas)
	if len(tangas) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tanga",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &tangas, nil
}

func (t *Tanga) Create(db *gorm.DB, name string, comment string) (*Tanga, error) {
	tanga := Tanga{
		Name:    name,
		Comment: comment,
	}
	db.Create(&tanga)
	return &tanga, nil
}

func (t *Tanga) Update(db *gorm.DB, id int, name string, comment string) (*Tanga, error) {
	db.Model(&Tanga{}).Where("ID =?", id).Update("name", name).Update("comment", comment)
	return t, nil
}

func (t *Tanga) Delete(db *gorm.DB, id int) (*Tanga, error) {
	tanga, err := t.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&tanga)
	return tanga, nil
}

func (t *Tanga) GetIDAsString() string {
	return fmt.Sprintf("%d", t.ID)
}
