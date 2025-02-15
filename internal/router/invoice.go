package router

import (
	"github.com/fnxr21/invoice-system/internal/handler"
	repositories "github.com/fnxr21/invoice-system/internal/repository"
	"github.com/fnxr21/invoice-system/internal/service"
	"github.com/fnxr21/invoice-system/pkg/mysql"
	"github.com/labstack/echo/v4"
)

func Invoice(e *echo.Group) {
	repo := repositories.Repository(mysql.DB)
	service := service.ServiceInvoice(repo)
	h := handler.HandlerInvoice(service)

	e.POST("/invoice", h.CreateInvoice)
	// e.GET("/invoice/list", h.ListInvoice)
	e.GET("/invoice/:id", h.GetInvoiceByID)
}
