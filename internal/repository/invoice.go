package repositories

import (
	// "fmt"

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

	// fmt.Println("invo", invoice.items)
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err = tx.Create(&invoice).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	//create invoice item
	for _, item := range invoice.InvoiceItem {

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
	if err = tx.First(&existingInvoice, "invoice = ?", invoice.ID).Error; err != nil {
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
	if err = tx.Where("id=?", invoice.ID).Delete(&model.InvoiceItem{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	//create invoice item
	for _, item := range invoice.InvoiceItem {
		invoiceitem := model.InvoiceItem{
			InvoiceID: invoice.ID, //add invoice id
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
