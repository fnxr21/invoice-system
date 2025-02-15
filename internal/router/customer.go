package router

import (
	"github.com/fnxr21/invoice-system/internal/handler"
	repositories "github.com/fnxr21/invoice-system/internal/repository"
	"github.com/fnxr21/invoice-system/internal/service"
	"github.com/fnxr21/invoice-system/pkg/mysql"
	"github.com/labstack/echo/v4"
)

func Customer(e *echo.Group) {
	repo := repositories.Repository(mysql.DB)
	service := service.ServiceCustomer(repo)
	h := handler.HandlerCustomer(service)

	e.POST("/customer", h.CreateCustomer)
	e.GET("/customers", h.ListCustomer)
	e.GET("/customer/:id", h.GetCustomerByID)
}
