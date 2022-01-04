package repository

import (
	"gitea.programmerfamily.com/go/product/internal/db"
	"gorm.io/gorm"
)

type AttributeValueInterface interface {

}

type attributeValueRepository struct {
	db *gorm.DB
}

func NewAttributeValueRepository() AttributeValueInterface {
	return &attributeValueRepository{
		db: db.GetDB(),
	}
}

// add
func (avr attributeValueRepository) Add()  {

}
// update
// get
// list
// delete
