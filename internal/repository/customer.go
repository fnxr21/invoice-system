package repositories

import "github.com/fnxr21/invoice-system/internal/model"

type Customer interface {
	CreateCustomer(cust model.Customer) (*model.Customer, error)
	GetCustomerByID(id uint) (*model.Customer, error)
	ListCustomer() ([]*model.Customer, error)
}

func (r *repository) CreateCustomer(customer model.Customer) (*model.Customer, error) {
	err := r.db.Create(&customer).Error
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
