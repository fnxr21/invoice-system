package repositories

import (
	"github.com/fnxr21/invoice-system/internal/model"
	"github.com/stretchr/testify/mock"
)

// MockAdminRepository is a mock implementation of AdminAuth interface
type MockItemRepository struct {
	mock.Mock
}

func (m *MockItemRepository) CreateItem(item model.Item) (*model.Item, error) {
	args := m.Mock.Called(item)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Item), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockItemRepository) GetItemByID(id uint) (*model.Item, error) {
	args := m.Mock.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Item), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockItemRepository) ListItem() ([]*model.Item, error) {
	args := m.Mock.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*model.Item), args.Error(1)
	}
	return nil, args.Error(1)
}
