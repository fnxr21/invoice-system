package service

import (
	"testing"

	itemdto "github.com/fnxr21/invoice-system/internal/dto/item"
	"github.com/fnxr21/invoice-system/internal/model"
	repositories "github.com/fnxr21/invoice-system/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateItem(t *testing.T) {
	mockRepo := new(repositories.MockItemRepository)

	service := ServiceItem(mockRepo)

	mockItem := &model.Item{
		Name: "fandi",
		Type: "Service",
	}
	request := &itemdto.ItemRequest{
		Name:    "fandi",
		Type: "Service",
	}
	mockRepo.On("CreateItem", mock.Anything).Return(mockItem, nil)
	result, err := service.CreateItem(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockItem.Name, result.Name)
	assert.Equal(t, mockItem.Type, result.Type)
	mockRepo.AssertExpectations(t)

}
func TestListItem(t *testing.T) {
	mockRepo := new(repositories.MockItemRepository)

	service := ServiceItem(mockRepo)

	mockItems := []*model.Item{
		{Name: "fandi", Type: "Front end"},
		{Name: "fandi nur", Type : "backend"},
	}
	mockRepo.On("ListItem").Return(mockItems, nil)
	result, err := service.ListItem()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)

}
func TestItemByID(t *testing.T) {
	mockRepo := new(repositories.MockItemRepository)

	service := ServiceItem(mockRepo)
	//skip this first
	mockItem := &model.Item{
		// ID:      1, // i cant pass id
		Name:    "fandi nur",
		Type: "backend",
	}
	mockRepo.On("GetItemByID", uint(1)).Return(mockItem, nil)
	result, err := service.GetItemByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	// assert.Equal(t, uint(1), result.ID)
	mockRepo.AssertExpectations(t)

}
