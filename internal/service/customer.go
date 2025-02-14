package service

import (

	customerdto "github.com/fnxr21/invoice-system/internal/dto/customer"
	"github.com/fnxr21/invoice-system/internal/model"
	repositories "github.com/fnxr21/invoice-system/internal/repository"
)

type CustomerService interface {
	CreateCustomer(request *customerdto.CustomerRequest) (*model.Customer, error)
	ListCustomer() ([]*model.Customer, error)
	GetCustomerByID(id uint) (*model.Customer, error)
}

type customerService struct {
	CustomerRepository repositories.Customer
}

func ServiceCustomer(CustomerRepository repositories.Customer) *customerService {
	return &customerService{CustomerRepository}
}

func (r *customerService) CreateCustomer(request *customerdto.CustomerRequest) (*model.Customer, error) {
	customer := model.Customer{
		Name:    request.Name,
		Address: request.Address,
	}
	
	user, err := r.CustomerRepository.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *customerService) ListCustomer() ([]*model.Customer, error) {
	user, err := r.CustomerRepository.ListCustomer()
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *customerService) GetCustomerByID(id uint) (*model.Customer, error) {
	user, err := r.CustomerRepository.GetCustomerByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}


