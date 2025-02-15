package repositories

import "github.com/fnxr21/invoice-system/internal/model"

type Item interface {
	CreateItem(item model.Item) (*model.Item, error)
	GetItemByID(id uint) (*model.Item, error)
	ListItem() ([]*model.Item, error)
}

func (r *repository) CreateItem(item model.Item) (*model.Item, error) {
	err := r.db.Create(&item).Error
	return &item, err
}
func (r *repository) GetItemByID(id uint) (*model.Item, error) {
	var customer *model.Item
	err := r.db.Where("id = ?", id).First(&customer).Error
	return customer, err
}
func (r *repository) ListItem() ([]*model.Item, error) {
	var customer []*model.Item
	err := r.db.Find(&customer).Error
	return customer, err
}
