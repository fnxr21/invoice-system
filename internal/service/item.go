package service

import (
	"fmt"

	itemdto "github.com/fnxr21/invoice-system/internal/dto/item"
	"github.com/fnxr21/invoice-system/internal/model"
	repositories "github.com/fnxr21/invoice-system/internal/repository"
)

type ItemService interface {
	CreateItem(request *itemdto.ItemRequest) (*model.Item, error)
	ListItem() ([]*model.Item, error)
	GetItemByID(id uint) (*model.Item, error)
}

type itemService struct {
	ItemRepository repositories.Item
}

func ServiceItem(ItemRepository repositories.Item) *itemService {
	return &itemService{ItemRepository}
}

func (r *itemService) CreateItem(request *itemdto.ItemRequest) (*model.Item, error) {
	item := model.Item{
		Name: request.Name,
		Type: request.Type,
	}
	fmt.Println(item)

	user, err := r.ItemRepository.CreateItem(item)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *itemService) ListItem() ([]*model.Item, error) {
	user, err := r.ItemRepository.ListItem()
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *itemService) GetItemByID(id uint) (*model.Item, error) {
	user, err := r.ItemRepository.GetItemByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
