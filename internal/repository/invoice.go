package repositories

import (
	"github.com/fnxr21/invoice-system/internal/model"
)

type Invoice interface {
	CreateInvoice(invoice model.Invoice) (*model.Invoice, error)
	GetInvoiceByID(id uint) (*model.Invoice, error)
	ListInvoice() ([]*model.Invoice, error)
	UpdateInvoice(invoice model.Invoice) (*model.Invoice, error)
}

func (r *repository) CreateInvoice(invoice model.Invoice) (*model.Invoice, error) {
	var err error

	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var lastInvoiceID uint
	// get last invoice ID safely and lock transaction in  database level with for update
	if err = tx.Raw("SELECT IFNULL(MAX(invoice_id), 0) FROM invoices FOR UPDATE").Scan(&lastInvoiceID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Assign new invoice ID
	invoice.InvoiceID = lastInvoiceID + 1

	if err = tx.Create(&invoice).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	//create invoice item
	for _, item := range invoice.InvoiceItem {
		invoiceitem := model.InvoiceItem{
			InvoiceID: invoice.InvoiceID, //add invoice id
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

	return &invoice, err

}
func (r *repository) GetInvoiceByID(id uint) (*model.Invoice, error) {
	var invoice *model.Invoice
	err := r.db.Where("id = ?", id).First(&invoice).Error
	return invoice, err
}
func (r *repository) ListInvoice() ([]*model.Invoice, error) {
	var invoice []*model.Invoice
	err := r.db.Find(&invoice).Error
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
	if err = tx.First(&existingInvoice, "invoice = ?", invoice.InvoiceID).Error; err != nil {
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
	
	//bulk delete invoiceitems base on invoiceID
	if err = tx.Where("invoice_id=?", invoice.InvoiceID).Delete(&model.InvoiceItem{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	//create invoice item
	for _, item := range invoice.InvoiceItem {
		invoiceitem := model.InvoiceItem{
			InvoiceID: invoice.InvoiceID, //add invoice id
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

	return &invoice, nil

}
