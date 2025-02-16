package repositories

import (
	"github.com/fnxr21/invoice-system/internal/model"
	errorhandler "github.com/fnxr21/invoice-system/pkg/error"
)

type Customer interface {
	CreateCustomer(cust model.Customer) (*model.Customer, error)
	GetCustomerByID(id uint) (*model.Customer, error)
	ListCustomer() ([]*model.Customer, error)
}

func (r *repository) CreateCustomer(customer model.Customer) (*model.Customer, error) {
	var err error

	// fmt.Println("invo", invoice.items)
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err = tx.Where("name = ? AND address = ?", customer.Name, customer.Address).Take(&customer).Error; err == nil {
		tx.Rollback()
		return nil, errorhandler.ErrCustomerExists
	}

	if err = tx.Create(&customer).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// last Commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &customer, err
}
func (r *repository) GetCustomerByID(id uint) (*model.Customer, error) {
	var customer *model.Customer
	err := r.db.Where("id = ?", id).First(&customer).Error
	return customer, err
}
func (r *repository) ListCustomer() ([]*model.Customer, error) {
	var customer []*model.Customer
	err := r.db.Find(&customer).Error
	return customer, err
}
