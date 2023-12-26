package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/k23dev/natalianatalia/pkg/tango_errors"
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

type TangaCounter struct {
	Total int
}

func NewTanga() *Tanga {
	return &Tanga{}
}

func (t *Tanga) Count(db *gorm.DB) (int, error) {
	counter := &TangaCounter{}
	db.Model(&Tanga{}).Select("count(ID) as total").Where("delete = ? ", "").Find(&counter)
	return counter.Total, nil
}

func (m *Tanga) FindOne(db *gorm.DB, id int) (*Tanga, error) {
	var this_model Tanga
	db.First(&this_model, id)
	if this_model.ID == 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tanga",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &this_model, nil
}

func (m *Tanga) FindAll(db *gorm.DB) (*[]Tanga, error) {
	var this_model []Tanga
	db.Order("created_at ASC").Find(&this_model)
	if len(this_model) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tanga",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &this_model, nil
}

func (t *Tanga) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Tanga, error) {
	tangas := []Tanga{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&tangas)
	if len(tangas) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tangas",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &tangas, nil
}

func (m *Tanga) Create(db *gorm.DB, codename string) (*Tanga, error) {
	m.Codename = codename
	db.Create(&m)
	return m, nil
}

func (m *Tanga) Update(db *gorm.DB, id int, codename string) (*Tanga, error) {
	db.Model(&Tanga{}).Where("ID =?", id).Update("codename", codename)
	return m, nil
}

func (m *Tanga) Delete(db *gorm.DB, id int) (*Tanga, error) {
	tanga, err := m.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&tanga)
	return tanga, nil
}

func (t *Tanga) GetIDAsString() string {
	return fmt.Sprintf("%d", t.ID)
}
