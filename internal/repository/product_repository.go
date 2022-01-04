package repository

import (
	"gitea.programmerfamily.com/go/product/internal/db"
	"gorm.io/gorm"
)

type ProductInterface interface {

}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository() ProductInterface {
	return &productRepository{
		db: db.GetDB(),
	}
}

// add
func (pr productRepository) Add()  {

}
// update
// get
// list
// delete
