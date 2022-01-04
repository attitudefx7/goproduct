package entity

type Attribute struct {
	ID               int64  `gorm:"id" json:"id"`
	Name             string `gorm:"name" json:"name"`
}

func (a *Attribute) TableName() string {
	return "pms_product_attribute"
}