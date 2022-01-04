package repository

import (
	"gitea.programmerfamily.com/go/product/internal/db"
	"gitea.programmerfamily.com/go/product/internal/entity"
	"gitea.programmerfamily.com/go/product/internal/form"
	"gorm.io/gorm"
)

type AttributeInterface interface {
	Create(form form.AddAttributeForm) error
	Update(form form.EditAttributeForm) error
	Delete(attributeForm form.DeleteAttributeForm) error
	Get(form form.GetAttributeForm) (entity.Attribute, error)
	GetList(form form.ListAttributeForm) ([]entity.Attribute, error)
}

type attributeRepository struct {
	db *gorm.DB
}

func NewAttributeRepository() AttributeInterface {
	return &attributeRepository{
		db: db.GetDB(),
	}
}

// add
func (ar *attributeRepository) Create(attributeForm form.AddAttributeForm) error {
	attribute := &entity.Attribute{
		Name: attributeForm.Name,
	}
	return ar.db.Create(&attribute).Error
}

func (ar *attributeRepository) Update(attributeForm form.EditAttributeForm)  error {
	return ar.db.Model(&entity.Attribute{}).
		Where("id = ?", attributeForm.ID).
		Update("name", attributeForm.Name).Error
}

func (ar attributeRepository) Delete(attributeForm form.DeleteAttributeForm) error {
	return ar.db.Delete(&entity.Attribute{}, attributeForm.ID).Error
}

func (ar *attributeRepository) Get(attributeForm form.GetAttributeForm) (entity.Attribute, error) {
	return entity.Attribute{}, nil
}

func (ar attributeRepository) GetList(attributeForm form.ListAttributeForm) ([]entity.Attribute, error) {
	var attributes []entity.Attribute
	var err error

	if err := ar.db.
		//Where("id = ?", 1).
		Find(&attributes).Error;err != nil {
		return nil, err
	}

	return attributes, err
}
// update
// get
// list
// delete
