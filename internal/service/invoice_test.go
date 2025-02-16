package service

import (
	"testing"
	"time"

	invoicedto "github.com/fnxr21/invoice-system/internal/dto/invoice"
	"github.com/fnxr21/invoice-system/internal/model"
	repositories "github.com/fnxr21/invoice-system/internal/repository"
	// "github.com/fnxr21/invoice-system/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// "github.com/stretchr/testify/mock"
)

func TestCreateInvoice(t *testing.T) {
	mockRepo := new(repositories.MockInvoiceRepository)
	service := ServiceInvoice(mockRepo)

	request := &invoicedto.InvoiceRequest{
		IssueDate:  "2024-02-16",
		DueDate:    "2024-03-16",
		Subject:    "Invoice for Service",
		CustomerID: 1,
		Items: []invoicedto.Item{
			{Name: "Item 1", Quantity: 20.00, UnitPrice: 10.00},
			{Name: "Item 2", Quantity: 20.00, UnitPrice: 10.00},
		},
	}

	issueDate, _ := time.Parse("2006-01-02", request.IssueDate)
	dueDate, _ := time.Parse("2006-01-02", request.DueDate)

	mockInvoice := &model.Invoice{
		IssueDate:  issueDate,
		DueDate:    dueDate,
		Subject:    request.Subject,
		CustomerID: request.CustomerID,
		InvoiceItem: []model.InvoiceItem{
			{Name: "Item 1", Quantity: 20.00, UnitPrice: 10.00},
			{Name: "Item 2", Quantity: 20.00, UnitPrice: 10.00},
		},
	}

	mockRepo.On("CreateInvoice", mock.Anything, request.Items).Return(mockInvoice, nil)

	result, err := service.CreateInvoice(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, request.Subject, result.Subject)
	mockRepo.AssertExpectations(t)
}

func TestGetInvoiceByID(t *testing.T) {
	mockRepo := new(repositories.MockInvoiceRepository)

	service := ServiceInvoice(mockRepo)

	mockInvoice := &model.Invoice{

		Subject:   "Invoice Test",
		IssueDate: time.Now(),
		DueDate:   time.Now().Add(30 * 24 * time.Hour),
	}
	mockInvoice.ID = 1

	mockRepo.On("GetInvoiceByID", uint(1)).Return(mockInvoice, nil)

	result, err := service.GetInvoiceByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, uint(1), result.ID)
	mockRepo.AssertExpectations(t)
}

func TestUpdateInvoice(t *testing.T) {
	mockRepo := new(repositories.MockInvoiceRepository)

	service := ServiceInvoice(mockRepo)

	request := &invoicedto.InvoiceRequestUpdate{
		ID:         1,
		IssueDate:  "2024-02-16",
		DueDate:    "2024-03-16",
		Subject:    "Updated Invoice",
		CustomerID: 1,
		Items: []invoicedto.Item{
			{Name: "Updated Item", Quantity: 5, UnitPrice: 2000},
		},
	}

	issueDate, _ := time.Parse("2006-01-02", request.IssueDate)
	dueDate, _ := time.Parse("2006-01-02", request.DueDate)

	mockInvoice := &model.Invoice{
		IssueDate:  issueDate,
		DueDate:    dueDate,
		Subject:    request.Subject,
		CustomerID: request.CustomerID,
		InvoiceItem: []model.InvoiceItem{
			{Name: "Updated Item", Quantity: 5, UnitPrice: 2000},
		},
	}
	mockInvoice.ID = request.ID

	mockRepo.On("UpdateInvoice", mock.Anything).Return(mockInvoice, nil)

	result, err := service.UpdateInvoice(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, request.Subject, result.Subject)
	mockRepo.AssertExpectations(t)
}

func TestIndexInvoice(t *testing.T) {
	mockRepo := new(repositories.MockInvoiceRepository)
	service := ServiceInvoice(mockRepo)

	// request := &invoicedto.InvoiceRequestUpdate{
	// 	ID:         1,
	// 	IssueDate:  "2024-02-16",
	// 	DueDate:    "2024-03-16",
	// 	Subject:    "Updated Invoice",
	// 	CustomerID: 1,
	// 	Items: []invoicedto.Item{
	// 		{Name: "Updated Item", Quantity: 5, UnitPrice: 2000},
	// 	},
	// }
	// issueDate, _ := time.Parse("2006-01-02", request.IssueDate)
	// dueDate, _ := time.Parse("2006-01-02", request.DueDate)

	req := invoicedto.InvoiceIndexing{
		InvoiceID:    1,
		IssueDate:    "2024-02-16",
		DueDate:      "2024-03-16",
		Subject:      "Test Invoice",
		CustomerName: "John Doe",
		TotalItems:   5,
		Status:       "paid",
		Page:         1,
		Size:         10,
	}


	issueDate, _ := time.Parse("2006-01-02", req.IssueDate)
	dueDate, _ := time.Parse("2006-01-02", req.DueDate)

	expectedInvoices := []*model.InvoiceIndexingNew{
		{
			InvoiceID:    1,
			IssueDate:    issueDate,
			DueDate:      dueDate,
			Subject:      "Test Invoice",
			CustomerName: "John Doe",
			TotalItems:   5,
			Status:       "paid",
		},
	}

	mockRepo.On("GetInvoceIndexing", mock.Anything).Return(expectedInvoices, nil)

	// invoices, pagination, err := service.GetInvoceIndexing(req)
	invoices, pagination,err := service.IndexInvoice(req)


	assert.Nil(t, err)
	assert.NotNil(t, invoices)
	assert.Equal(t, 1, pagination.CurrentPage)
	assert.Equal(t, 10, pagination.Size)
	assert.Equal(t, expectedInvoices, invoices)

	mockRepo.AssertExpectations(t)
}
