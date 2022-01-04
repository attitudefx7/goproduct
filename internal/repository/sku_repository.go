package repository

import (
	"gitea.programmerfamily.com/go/product/internal/db"
	"gorm.io/gorm"
)

type SkuInterface interface {

}

type skuRepository struct {
	db *gorm.DB
}

func NewSkuRepository() SkuInterface {
	return &skuRepository{
		db: db.GetDB(),
	}
}

// add
func (sr skuRepository) Add()  {

}
// update
// get
// list
// delete
