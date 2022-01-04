package entity

type AttributeValue struct {
	ID          int64  `gorm:"id" json:"id"`
	ProductId   int64  `gorm:"product_id" json:"product_id"`
	AttributeId int64  `gorm:"attribute_id" json:"attribute_id"`
	Value       string `gorm:"value" json:"value"`
}

func (av *AttributeValue) TableName() string {
	return "pms_product_attribute_value"
}