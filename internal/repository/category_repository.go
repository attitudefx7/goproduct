package repository

import (
	"gitea.programmerfamily.com/go/product/internal/db"
	"gorm.io/gorm"
)

type CategoryInterface interface {

}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() CategoryInterface {
	return &categoryRepository{
		db: db.GetDB(),
	}
}

// add
func (cr categoryRepository) Add()  {

}
// update
// get
// list
// delete
