package controller

import (
	"Crud/entity"
	"Crud/service"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	GetCustomers() []entity.Customer
	AddCustomer(ctx *gin.Context) entity.Customer
	UpdateCustomer(ctx *gin.Context) bool
}

type customerController struct {
	service service.CustomerService
}

func New(customerService service.CustomerService) CustomerController {
	return &customerController{
		service: customerService,
	}
}

func (c *customerController) GetCustomers() []entity.Customer {
	return c.service.GetCustomers()
}

func (c *customerController) AddCustomer(ctx *gin.Context) entity.Customer {
	var customer entity.Customer
	ctx.BindJSON(&customer)
	return c.service.AddCustomer(customer)
}

func (c *customerController) UpdateCustomer(ctx *gin.Context) bool {
	var customer entity.Customer
	ctx.BindJSON(&customer)
	return c.service.UpdateCustomer(customer)
}
