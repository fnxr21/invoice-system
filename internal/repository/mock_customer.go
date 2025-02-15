package repositories

import (
	"github.com/fnxr21/invoice-system/internal/model"
	"github.com/stretchr/testify/mock"
)

// MockAdminRepository is a mock implementation of AdminAuth interface
type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) CreateCustomer(customer model.Customer) (*model.Customer, error) {
	args := m.Mock.Called(customer)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCustomerRepository) GetCustomerByID(id uint) (*model.Customer, error) {
	args := m.Mock.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCustomerRepository) ListCustomer() ([]*model.Customer, error) {
	args := m.Mock.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*model.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}
