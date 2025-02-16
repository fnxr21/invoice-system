package repositories

import (
	// "fmt"

	"fmt"

	invoicedto "github.com/fnxr21/invoice-system/internal/dto/invoice"
	"github.com/fnxr21/invoice-system/internal/model"
)

type Invoice interface {
	CreateInvoice(invoice model.Invoice, items []invoicedto.Item) (*model.Invoice, error)
	GetInvoiceByID(id uint) (*model.Invoice, error)
	ListInvoice() ([]*model.Invoice, error)
	UpdateInvoice(invoice model.Invoice) (*model.Invoice, error)
	//indexing
	GetInvoceIndexing(filter model.InvoiceIndexing) ([]*model.InvoiceIndexingNew, error)
}

func (r *repository) CreateInvoice(invoice model.Invoice, items []invoicedto.Item) (*model.Invoice, error) {
	var err error

	// fmt.Println("invo", invoice.items)
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err = tx.Create(&invoice).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, item := range items {

		invoiceitem := model.InvoiceItem{
			InvoiceID: invoice.ID, //add invoice id
			ItemID:    item.ID,
			Name:      item.Name,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}
		// if err
		if err = tx.Create(&invoiceitem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// last Commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Load Customer agar bisa dikembalikan dalam response
	if err := r.db.Preload("Customer").Preload("InvoiceItem").First(&invoice, invoice.ID).Error; err != nil {
		return nil, err
	}
	return &invoice, err

}

func (r *repository) GetInvoiceByID(id uint) (*model.Invoice, error) {
	var invoice *model.Invoice

	err := r.db.Preload("Customer").Preload("InvoiceItem").Where("id = ?", id).First(&invoice).Error
	return invoice, err
}
func (r *repository) ListInvoice() ([]*model.Invoice, error) {
	var invoice []*model.Invoice
	err := r.db.Preload("Customer").Preload("InvoiceItem").Find(&invoice).Error
	return invoice, err
}

func (r *repository) UpdateInvoice(invoice model.Invoice) (*model.Invoice, error) {
	var err error
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	//check base on invoiceid
	var existingInvoice model.Invoice
	if err = tx.First(&existingInvoice, "invoice.id = ?", invoice.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// check invoice is same or not,if yes update the customer
	if existingInvoice.CustomerID != invoice.CustomerID {
		existingInvoice.CustomerID = invoice.CustomerID
	}
	if invoice.Subject != "" && invoice.Subject != existingInvoice.Subject {
		existingInvoice.Subject = invoice.Subject
	}
	if !invoice.IssueDate.IsZero() && invoice.IssueDate != existingInvoice.IssueDate {
		existingInvoice.IssueDate = invoice.IssueDate
	}

	// Step 3: Update DueDate if it's different
	if !invoice.DueDate.IsZero() && invoice.DueDate != existingInvoice.DueDate {
		existingInvoice.DueDate = invoice.DueDate
	}

	if err = tx.Save(&existingInvoice).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// last Commit
	if err := tx.Commit().Error; err != nil {
		return &invoice, err
	}
	tx = r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	//bulk delete invoiceitems base on invoiceID
	if err = tx.Where("invoice_id=?", invoice.ID).Unscoped().Delete(&model.InvoiceItem{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// create invoice item
	uniqueItems := make(map[int]bool)
	for _, item := range invoice.InvoiceItem {
		if _, exists := uniqueItems[int(item.InvoiceID)]; exists {
			continue // Skip jika item sudah ada
		}

		uniqueItems[int(item.ItemID)] = true
		invoiceitem := model.InvoiceItem{
			InvoiceID: invoice.ID, //add invoice id
			ItemID:    item.ItemID,
			Name:      item.Name,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}
		if err = tx.Create(&invoiceitem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// last Commit
	if err := tx.Commit().Error; err != nil {
		return &invoice, err
	}

	// Load Customer agar bisa dikembalikan dalam response
	if err := r.db.Preload("Customer").Preload("InvoiceItem").First(&invoice, invoice.ID).Error; err != nil {
		return nil, err
	}

	return &invoice, nil

}

func (r *repository) GetInvoceIndexing(filter model.InvoiceIndexing) ([]*model.InvoiceIndexingNew, error) {

	pagesize := filter.Size
	offset := (filter.Page - 1) * pagesize

	var invoiceValue []*model.InvoiceIndexingNew

	query := r.db.
		Limit(pagesize).
		Offset(offset).
		Model(&model.Invoice{}).
		Select("invoice.id as invoice_id, invoice.issue_date, invoice.due_date, invoice.subject, customer.name as customer_name, invoice.status, COALESCE(COUNT(invoice_item.id), 0) AS total_items").
		Joins("LEFT JOIN customer ON  invoice.customer_id = customer.id ").
		Joins("LEFT JOIN invoice_item ON invoice.id = invoice_item.invoice_id").
		Group("invoice.id, customer.name, invoice.issue_date, invoice.due_date, invoice.subject, invoice.status")

	if filter.InvoiceID != 0 {
		query = query.Where("invoice.id LIKE ?", "%"+fmt.Sprint(filter.InvoiceID)+"%")
	}
	if filter.Subject != "" {
		query = query.Where("invoice.subject LIKE ?", "%"+filter.Subject+"%")
	}
	if filter.CustomerName != "" {
		query = query.Where("customer.name LIKE ?", "%"+filter.CustomerName+"%")
	}
	if filter.Status != "" {
		query = query.Where("invoice.status = ?", filter.Status)
	}
	if !filter.IssueDate.IsZero() {
		query = query.Where("DATE(invoice.issue_date) = ?", filter.IssueDate.Format("2006-01-02"))
	}
	if !filter.DueDate.IsZero() {
		query = query.Where("DATE(invoice.due_date) = ?", filter.DueDate.Format("2006-01-02"))
	}

	if filter.TotalItems > 0 {
		query = query.Having("total_items = ?", filter.TotalItems)
	}

	if err := query.
		Find(&invoiceValue).Error; err != nil {
		return nil, err
	}

	return invoiceValue, nil

}
