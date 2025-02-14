package repositories

import (
	"fmt"

	"github.com/fnxr21/invoice-system/internal/model"
)

type InvoiceIndexing interface {
	GetInvoceIndexing(filter model.InvoiceIndexing) ([]*model.InvoiceIndexing, error)
}

func (r *repository) GetInvoceIndexing(filter model.InvoiceIndexing) ([]*model.InvoiceIndexing, error) {

	pagesize := filter.Size
	offset := (filter.Page - 1) * pagesize

	var invoiceValue []*model.InvoiceIndexing

	query := r.db.
		Limit(pagesize).
		Offset(offset).
		Joins("LEFT JOIN customer ON  invoice.customer_id = customer.id ").
		Joins("LEFT JOIN invoice_item ON invoice.invoice_id = invoice_item.id")

	if filter.InvoiceID != 0 {
		query = query.Where("invoice.invoice_id LIKE ?", "%"+fmt.Sprint(filter.InvoiceID)+"%")
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

	if err := query.Find(&invoiceValue).Error; err != nil {
		return nil, err
	}

	return invoiceValue, nil

}
