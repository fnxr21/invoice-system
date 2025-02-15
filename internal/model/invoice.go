package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	//database
	Invoice struct {
		gorm.Model
		// InvoiceID     uint          `gorm:"primaryKey;index"`
		IssueDate     time.Time     `gorm:"index"`
		Subject       string        `gorm:"index"`
		DueDate       time.Time     `gorm:"index"`
		Status        string        `gorm:"type:enum('paid', 'unpaid');default:'unpaid';not null;index"`
		CustomerID    uint          `gorm:"index"`
		Customer      Customer      `gorm:"foreignKey:CustomerID" json:"customer"`
		InvoiceItem   []InvoiceItem `json:"invoice"`
	}
	// model for repository
	InvoiceIndexing struct {
		InvoiceID    uint      `json:"invoice_id"`
		IssueDate    time.Time `json:"issue_date"`
		DueDate      time.Time `json:"Due_date"`
		Subject      string    `json:"subject"`
		CustomerName string    `json:"customer_name"`
		TotalItems   int       `json:"total_items"`
		Status       string    `json:"status"`
		Page         int       `json:"page"`
		Size         int       `json:"size"`
	}
)
