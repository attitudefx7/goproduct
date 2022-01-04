package entity

type Sku struct {
	ID             int64   `gorm:"id" json:"id"`
	ProductId      int64   `gorm:"product_id" json:"product_id"`
	SkuCode        string  `gorm:"sku_code" json:"sku_code"`
	Price          float32 `gorm:"price" json:"price"`
	Stock          int     `gorm:"stock" json:"stock"`
	LowStock       int     `gorm:"low_stock" json:"low_stock"`
	Pic            string  `gorm:"pic" json:"pic"`
	Sale           int     `gorm:"sale" json:"sale"`
	PromotionPrice float32 `gorm:"promotion_price" json:"promotion_price"`
	LockStock      int     `gorm:"lock_stock" json:"lock_stock"`
	SpData         string  `gorm:"sp_data" json:"sp_data"`
}

func (s *Sku) TableName() string {
	return "pms_sku_stock"
}