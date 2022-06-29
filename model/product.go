package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string      `gorm:"column:name; not null"`
	Customer    []*Customer `gorm:"many2many:customer_products"`
}
