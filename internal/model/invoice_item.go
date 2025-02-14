package model

import "gorm.io/gorm"

type InvoiceItem struct {
	gorm.Model
	InvoiceID   uint
	ItemID      uint
	InvoiceItem Item    `gorm:"foreignKey:ItemID" json:"-"`
	Quantity    float64 `gorm:"type:decimal(15,2);not null"`
	UnitPrice   float64 `gorm:"type:decimal(15,2);not null"`
}
