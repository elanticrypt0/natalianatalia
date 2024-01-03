package models

import (
	"fmt"
	"strconv"

	"github.com/k23dev/natalianatalia/pkg/tango_errors"
	"gorm.io/gorm"
)

type Scapp_param struct {
	gorm.Model
	ScappID    uint
	Scapp      Scapp
	CategoryID uint
	Category   Category
	Method     string
	InputType  string
	IsFlag     bool
	Comment    string
	Order      uint
}

type Scapp_paramDTO struct {
	ScappID    uint `json:"scapp_id" param:"scapp_id" query:"scapp_id" form:"scapp_id"`
	Scapp      Scapp
	CategoryID uint `json:"category_id" param:"category_id" query:"category_id" form:"category_id"`
	Category   Category
	Method     string `json:"method" param:"method" query:"method" form:"method"`
	InputType  string `json:"input_type" param:"input_type" query:"input_type" form:"input_type"`
	IsFlag     string `json:"is_flag" param:"is_flag" query:"is_flag" form:"is_flag"`
	IsFlagBool bool
	Comment    string `json:"comment" param:"comment" query:"comment" form:"comment"`
	Order      uint   `json:"order" param:"order" query:"order" form:"order"`
}

type Scapp_paramCounter struct {
	Total int
}

func NewScapp_param() *Scapp_param {
	return &Scapp_param{}
}

func (s *Scapp_param) Count(db *gorm.DB) (int, error) {
	counter := &Scapp_paramCounter{}
	db.Model(&Scapp_param{}).Select("count(ID) as total").Find(&counter)
	return counter.Total, nil
}

func (s *Scapp_param) FindOne(db *gorm.DB, id int) (*Scapp_param, error) {
	var scapp_param Scapp_param
	db.First(&scapp_param, id)
	if scapp_param.ID == 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Scapp_param",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &scapp_param, nil
}

func (s *Scapp_param) FindAll(db *gorm.DB) ([]Scapp_param, error) {
	var scapp_params []Scapp_param
	db.Order("created_at ASC").Find(&scapp_params)
	if len(scapp_params) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Scapp_param",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return scapp_params, nil
}

func (s *Scapp_param) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Scapp_param, error) {
	scapp_params := []Scapp_param{}

	db.Order("created_at ASC").Preload("Scapp").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&scapp_params)
	if len(scapp_params) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Scapp_param",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &scapp_params, nil
}

func (s *Scapp_param) Create(db *gorm.DB, dto Scapp_paramDTO) (*Scapp_param, error) {
	s.SatinizeDTOCreate(&dto)
	scapp_param := Scapp_param{
		ScappID:    dto.ScappID,
		CategoryID: dto.CategoryID,
		Method:     dto.Method,
		InputType:  dto.InputType,
		IsFlag:     dto.IsFlagBool,
		Comment:    dto.Comment,
		Order:      dto.Order,
	}

	db.Create(&scapp_param)
	return &scapp_param, nil
}

func (s *Scapp_param) Update(db *gorm.DB, id int, dto Scapp_paramDTO) (*Scapp_param, error) {
	s.SatinizeDTOUpdate(&dto)
	db.Model(&Scapp_param{}).Where("ID =?", id).Update("method", dto.Method).Update("category_id", dto.CategoryID).Update("scapp_id", dto.ScappID).Update("input_type", dto.InputType).Update("is_flag", dto.IsFlagBool).Update("comment", dto.Comment).Update("order", dto.Order)
	return s, nil
}

func (s *Scapp_param) Delete(db *gorm.DB, id int) (*Scapp_param, error) {
	scapp_param, err := s.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&scapp_param)
	return scapp_param, nil
}

func (s *Scapp_param) GetIDAsString() string {
	return fmt.Sprintf("%d", s.ID)
}

func (s *Scapp_param) SatinizeDTOCreate(dto *Scapp_paramDTO) error {
	flagb, _ := strconv.ParseBool(dto.IsFlag)
	dto.IsFlagBool = flagb
	return nil
}

func (s *Scapp_param) SatinizeDTOUpdate(dto *Scapp_paramDTO) error {
	flagb, _ := strconv.ParseBool(dto.IsFlag)
	dto.IsFlagBool = flagb
	return nil
}
