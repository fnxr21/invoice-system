package service

import (
	"fmt"
	"time"

	invoicedto "github.com/fnxr21/invoice-system/internal/dto/invoice"
	"github.com/fnxr21/invoice-system/internal/model"
	repositories "github.com/fnxr21/invoice-system/internal/repository"
)

type InvoiceService interface {
	CreateInvoice(request *invoicedto.InvoiceRequest) (*model.Invoice, error)
	IndexInvoice(req invoicedto.InvoiceIndexing) ([]*model.InvoiceIndexingNew, *invoicedto.PagingInvoiceIndexing, error)
	GetInvoiceByID(id uint) (*model.Invoice, error)
}

type invoiceService struct {
	InvoiceRepository repositories.Invoice
}

func ServiceInvoice(InvoiceRepository repositories.Invoice) *invoiceService {
	return &invoiceService{InvoiceRepository}
}

func (r *invoiceService) CreateInvoice(request *invoicedto.InvoiceRequest) (*model.Invoice, error) {

	issueDate, err := parseDate(request.IssueDate)
	if err != nil {
		return nil, err // Handle error for issueDate
	}

	dueDate, err := parseDate(request.DueDate)
	if err != nil {
		return nil, err // Handle error for dueDate
	}
	invoice := model.Invoice{
		IssueDate:  issueDate,
		DueDate:    dueDate,
		Subject:    request.Subject,
		CustomerID: request.CustomerID,
	}

	//append direct to invoice
	for _, item := range request.Items {
		invoice.InvoiceItem = append(invoice.InvoiceItem, model.InvoiceItem{
			ItemID:    item.ID,
			Name:      item.Name,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		})
	}
	// fmt.Println(invoice.InvoiceItem)
	fmt.Println()
	// fmt.Println(request)
	// fmt.Println()

	createInvoice, err := r.InvoiceRepository.CreateInvoice(invoice)
	if err != nil {
		return nil, err
	}

	return createInvoice, nil
}
func (r *invoiceService) IndexInvoice(req invoicedto.InvoiceIndexing) ([]*model.InvoiceIndexingNew, *invoicedto.PagingInvoiceIndexing, error) {

	issueDate, err := parseDate(req.IssueDate)
	if err != nil {
		return nil, nil, err // Handle error for issueDate
	}

	fmt.Println("check", issueDate)
	dueDate, err := parseDate(req.DueDate)
	if err != nil {
		return nil, nil, err // Handle error for dueDate
	}

	fmt.Println()
	invoceIndex := model.InvoiceIndexing{
		InvoiceID:    req.InvoiceID,
		IssueDate:    issueDate,
		DueDate:      dueDate,
		Subject:      req.Subject,
		CustomerName: req.CustomerName,
		TotalItems:   req.TotalItems,
		Status:       req.Status,
		Page:         req.Page,
		Size:         req.Size,
	}
	// user, total_page, err := r.InvoiceRepository.GetInvoceIndexing(invoceIndex)
	user, err := r.InvoiceRepository.GetInvoceIndexing(invoceIndex)
	if err != nil {
		return nil, nil, err
	}
	pagination := invoicedto.PagingInvoiceIndexing{
		Size:        req.Size,
		CurrentPage: req.Page,
	}

	return user, &pagination, nil
}
func (r *invoiceService) GetInvoiceByID(id uint) (*model.Invoice, error) {
	user, err := r.InvoiceRepository.GetInvoiceByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func parseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02" // Adjust the format if necessary

	// Parse the string into time.Time
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		parsedDate.IsZero()
		// return time.Time{}, fmt.Errorf("invalid date format: %v", err)
	}
	return parsedDate, nil
}
