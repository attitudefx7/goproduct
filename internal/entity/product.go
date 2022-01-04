package entity

type Product struct {
	ID               int64   `gorm:"id" json:"id"`
	BrandId          int64   `gorm:"brand_id" json:"brand_id"`
	FeightTemplateId int64   `gorm:"feight_template_id" json:"feight_template_id"`
	Name             string  `gorm:"name" json:"name"`
	Pic              string  `gorm:"pic" json:"pic"`
	ProductSn        string  `gorm:"product_sn" json:"product_sn"`
	DeleteStatus     int     `gorm:"delete_status" json:"delete_status"`
	PublishStatus    int     `gorm:"publish_status" json:"publish_status"`
	NewStatus        int     `gorm:"new_status" json:"new_status"`
	RecommandStatus  int     `gorm:"recommand_status" json:"recommand_status"`
	Sort             int     `gorm:"sort" json:"sort"`
	Sale             int     `gorm:"sale" json:"sale"`
	Price            float32 `gorm:"price" json:"price"`
	PromotionPrice   float32 `gorm:"promotion_price" json:"promotion_price"`
	SubTitle         string  `gorm:"sub_title" json:"sub_title"`
	Description      string  `gorm:"description" json:"description"`
	Stock            int     `gorm:"stock" json:"stock"`
	LowStock         int     `gorm:"low_stock" json:"low_stock"`
	Keywords         string  `gorm:"keywords" json:"keywords"`
	Note             string  `gorm:"note" json:"note"`
	AlbumPics        string  `gorm:"album_pics" json:"album_pics"`
	DetailTitle      string  `gorm:"detail_title" json:"detail_title"`
	DetailDesc       string  `gorm:"detail_desc" json:"detail_desc"`
	CategoryName     string  `gorm:"category_name" json:"category_name"`
	BrandName        string  `gorm:"brand_name" json:"brand_name"`
}

func (p *Product) TableName() string {
	return "pms_product"
}