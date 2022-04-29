package service

import "Crud/entity"

type CustomerService interface {
	GetCustomers() []entity.Customer
	AddCustomer(entity.Customer) entity.Customer
	UpdateCustomer(entity.Customer) bool
}

type customerService struct {
	customers []entity.Customer
}

func New() CustomerService {
	return &customerService{}
}

func (service *customerService) GetCustomers() []entity.Customer {
	return service.customers
}

func (service *customerService) AddCustomer(customer entity.Customer) entity.Customer {
	service.customers = append(service.customers, customer)
	return customer
}

func (service *customerService) UpdateCustomer(customer entity.Customer) bool {
	for _, c := range service.customers {
		if c.Id == customer.Id {
			c.Name = customer.Name
			return true
		}
	}
	return false
}
