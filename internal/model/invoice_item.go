package model

import "gorm.io/gorm"

type InvoiceItem struct {
	gorm.Model
	InvoiceID   uint `gorm:"index"`
	ItemID      uint `gorm:"index"`
	InvoiceItem Item `gorm:"foreignKey:ItemID" json:"-"`
	Name        string
	Quantity    float64 `gorm:"type:decimal(15,2);not null"`
	UnitPrice   float64 `gorm:"type:decimal(15,2);not null"`
}
