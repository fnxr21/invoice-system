package repositories

import (
	"github.com/fnxr21/invoice-system/internal/model"
	errorhandler "github.com/fnxr21/invoice-system/pkg/error"
)

type Item interface {
	CreateItem(item model.Item) (*model.Item, error)
	GetItemByID(id uint) (*model.Item, error)
	ListItem() ([]*model.Item, error)
}

func (r *repository) CreateItem(item model.Item) (*model.Item, error) {

	var err error

	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err = tx.Where("name = ? AND type = ?", item.Name, item.Type).Take(&item).Error; err == nil {
		tx.Rollback()
		return nil, errorhandler.ErrCustomerExists
	}

	if err = tx.Create(&item).Error; err != nil {
		tx.Rollback()
		return nil, err

	}
	// last Commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

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
