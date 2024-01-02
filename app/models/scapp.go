package models

import (
	"fmt"

	"github.com/k23dev/natalianatalia/pkg/tango_errors"
	"gorm.io/gorm"
)

type Scapp struct {
	gorm.Model
	Name    string
	Path    string
	IsSudo  bool
	Comment string
}

type ScappDTO struct {
	Name    string `json:"name" param:"name" query:"name" form:"name"`
	Path    string `json:"path" param:"path" query:"path" form:"path"`
	IsSudo  bool   `json:"is_sudo" param:"is_sudo" query:"is_sudo" form:"is_sudo"`
	Comment string `json:"comment" param:"comment" query:"comment" form:"comment"`
}

type ScappCounter struct {
	Total int
}

func NewScapp() *Scapp {
	return &Scapp{}
}

func (c *Scapp) Count(db *gorm.DB) (int, error) {
	counter := &ScappCounter{}
	db.Model(&Scapp{}).Select("count(ID) as total").Find(&counter)
	fmt.Printf("%v", counter.Total)
	return counter.Total, nil
}

func (c *Scapp) FindOne(db *gorm.DB, id int) (*Scapp, error) {
	var scapp Scapp
	db.First(&scapp, id)
	if scapp.ID == 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Scapp",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &scapp, nil
}

func (c *Scapp) FindAll(db *gorm.DB) ([]Scapp, error) {
	var scapps []Scapp
	db.Order("created_at ASC").Find(&scapps)
	if len(scapps) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Scapp",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return scapps, nil
}

func (c *Scapp) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Scapp, error) {
	scapps := []Scapp{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&scapps)
	if len(scapps) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Scapp",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &scapps, nil
}

func (c *Scapp) Create(db *gorm.DB, cDTO ScappDTO) (*Scapp, error) {
	scapp := Scapp{
		Name:    cDTO.Name,
		Path:    cDTO.Path,
		IsSudo:  cDTO.IsSudo,
		Comment: cDTO.Comment,
	}
	db.Create(&scapp)
	return &scapp, nil
}

func (c *Scapp) Update(db *gorm.DB, id int, name string, path string, isSudo bool, comment string) (*Scapp, error) {
	db.Model(&Scapp{}).Where("ID =?", id).Update("name", name).Update("path", path).Update("is_sudo", isSudo).Update("comment", comment)
	return c, nil
}

func (c *Scapp) Delete(db *gorm.DB, id int) (*Scapp, error) {
	scapp, err := c.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&scapp)
	return scapp, nil
}

func (c *Scapp) GetIDAsString() string {
	return fmt.Sprintf("%d", c.ID)
}

func (c *Scapp) GetIsSudoAsString() string {
	return fmt.Sprintf("%v", c.IsSudo)
}
