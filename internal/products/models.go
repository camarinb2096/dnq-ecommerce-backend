package products

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"column:name"`
	Description string  `gorm:"column:description"`
	Image       string  `gorm:"column:image"`
	Price       float32 `gorm:"column:price"`
	Quantity    int     `gorm:"column:quantity"`
}
