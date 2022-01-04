package repository

import (
	"gitea.programmerfamily.com/go/product/internal/db"
	"gorm.io/gorm"
)

type BrandInterface interface {

}

type brandRepository struct {
	db *gorm.DB
}

func NewBrandRepository() BrandInterface {
	return &brandRepository{
		db: db.GetDB(),
	}
}

// add
func (br brandRepository) Add()  {

}
// update
// get
// list
// delete
