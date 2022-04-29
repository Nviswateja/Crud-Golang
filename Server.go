package main

import (
	"Crud/controller"
	"Crud/service"

	"github.com/gin-gonic/gin"
)

var (
	customerService    service.CustomerService       = service.New()
	customerController controller.CustomerController = controller.New(customerService)
)

func main() {
	server := gin.Default()

	server.GET("/GetCustomers", func(ctx *gin.Context) {
		ctx.JSON(200, customerController.GetCustomers())
	})

	server.POST("/AddCustomer", func(ctx *gin.Context) {
		ctx.JSON(200, customerController.AddCustomer(ctx))
	})

	server.PUT("/UpdateCustomer", func(ctx *gin.Context) {
		ctx.JSON(200, customerController.UpdateCustomer(ctx))
	})

	server.Run()
}
