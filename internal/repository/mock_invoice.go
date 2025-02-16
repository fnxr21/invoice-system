package repositories

import (
	"github.com/fnxr21/invoice-system/internal/model"
	"github.com/stretchr/testify/mock"
	invoicedto "github.com/fnxr21/invoice-system/internal/dto/invoice"
)

// MockInvoiceRepository is a mock implementation of Invoice interface
type MockInvoiceRepository struct {
	mock.Mock
}

func (m *MockInvoiceRepository) CreateInvoice(invoice model.Invoice, items []invoicedto.Item) (*model.Invoice, error) {
	args := m.Mock.Called(invoice, items)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Invoice), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockInvoiceRepository) GetInvoiceByID(id uint) (*model.Invoice, error) {
	args := m.Mock.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Invoice), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockInvoiceRepository) ListInvoice() ([]*model.Invoice, error) {
	args := m.Mock.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*model.Invoice), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockInvoiceRepository) UpdateInvoice(invoice model.Invoice) (*model.Invoice, error) {
	args := m.Mock.Called(invoice)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Invoice), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockInvoiceRepository) GetInvoceIndexing(filter model.InvoiceIndexing) ([]*model.InvoiceIndexingNew, error) {
	args := m.Mock.Called(filter)
	if args.Get(0) != nil {
		return args.Get(0).([]*model.InvoiceIndexingNew), args.Error(1)
	}
	return nil, args.Error(1)
}
