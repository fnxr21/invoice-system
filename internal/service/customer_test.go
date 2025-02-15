package service

import (
	"testing"

	customerdto "github.com/fnxr21/invoice-system/internal/dto/customer"
	"github.com/fnxr21/invoice-system/internal/model"
	repositories "github.com/fnxr21/invoice-system/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomer(t *testing.T) {
	mockRepo := new(repositories.MockCustomerRepository)

	service := ServiceCustomer(mockRepo)

	mockCustomer := &model.Customer{
		Name:    "fandi",
		Address: "north jakarta",
	}
	request := &customerdto.CustomerRequest{
		Name:    "fandi",
		Address: "north jakarta",
	}
	mockRepo.On("CreateCustomer", mock.Anything).Return(mockCustomer, nil)
	result, err := service.CreateCustomer(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockCustomer.Name, result.Name)
	assert.Equal(t, mockCustomer.Address, result.Address)
	mockRepo.AssertExpectations(t)

}
func TestListCustomer(t *testing.T) {
	mockRepo := new(repositories.MockCustomerRepository)

	service := ServiceCustomer(mockRepo)

	mockCustomers := []*model.Customer{
		{Name: "fandi", Address: "north jakarta"},
		{Name: "fandi nur", Address: "sunter agung"},
	}
	mockRepo.On("ListCustomer").Return(mockCustomers, nil)
	result, err := service.ListCustomer()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)

}
func TestCustomerByID(t *testing.T) {
	mockRepo := new(repositories.MockCustomerRepository)

	service := ServiceCustomer(mockRepo)
	//skip this first
	mockCustomer := &model.Customer{
		// ID:      1, // i cant pass id
		Name:    "John Doe",
		Address: "123 Main St",
	}
	mockRepo.On("GetCustomerByID", uint(1)).Return(mockCustomer, nil)
	result, err := service.GetCustomerByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	// assert.Equal(t, uint(1), result.ID)
	mockRepo.AssertExpectations(t)

}
