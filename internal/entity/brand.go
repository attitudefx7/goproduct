package entity

type Brand struct {
	ID                  int64  `gorm:"id" json:"id"`
	Name                string `gorm:"name" json:"name"`
	FirstLetter         string `gorm:"first_letter" json:"first_letter"`
	Sort                int    `gorm:"sort" json:"sort"`
	FactoryStatus       int    `gorm:"factory_status" json:"factory_status"`
	ShowStatus          int    `gorm:"show_status" json:"show_status"`
	ProductCount        int    `gorm:"product_count" json:"product_count"`
	ProductCommentCount int    `gorm:"product_comment_count" json:"product_comment_count"`
	Logo                string `gorm:"logo" json:"logo"`
	BigPic              string `gorm:"big_pic" json:"big_pic"`
	BrandStory          string `gorm:"brand_story" json:"brand_story"`
}

func (b *Brand) TableName() string {
	return "pms_brand"
}