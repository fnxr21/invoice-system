package model

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	InvoiceID     uint
	IssueDate     time.Time
	DueDate       time.Time
	status        string `gorm:"type:enum('paid', 'unpaid');default:'unpaid';not null"`
	CustomerID    uint
	Customer      Customer `gorm:"foreignKey:CustomerID" json:"-"`
	InvoiceItemID uint
	InvoiceItem   InvoiceItem `gorm:"foreignKey:InvoiceItemID" json:"-"`
}
