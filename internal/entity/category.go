package entity

type Category struct {
	ID           int64  `gorm:"id" json:"id"`
	ParentId     int64  `gorm:"parent_id" json:"parent_id"`
	Name         string `gorm:"name" json:"name"`
	Level        int    `gorm:"level" json:"level"`
	ProductCount int    `gorm:"product_count" json:"product_count"`
	ProductUnit  int    `gorm:"product_unit" json:"product_unit"`
	NavStatus    int    `gorm:"nav_status" json:"nav_status"`
	ShowStatus   int    `gorm:"show_status" json:"show_status"`
	Sort         int    `gorm:"sort" json:"sort"`
	Icon         string `gorm:"icon" json:"icon"`
	Keywords     string `gorm:"keywords" json:"keywords"`
	Description  string `gorm:"description" json:"description"`
}

func (c *Category) TableName() string {
	return "pms_product_category"
}