package router

import (
	"github.com/fnxr21/invoice-system/internal/handler"
	repositories "github.com/fnxr21/invoice-system/internal/repository"
	"github.com/fnxr21/invoice-system/internal/service"
	"github.com/fnxr21/invoice-system/pkg/mysql"
	"github.com/labstack/echo/v4"
)

func Item(e *echo.Group) {
	repo := repositories.Repository(mysql.DB)
	service := service.ServiceItem(repo)
	h := handler.HandlerItem(service)

	e.POST("/item", h.CreateItem)
	e.GET("/items", h.ListItem)
	e.GET("/item/:id", h.GetItemByID)
}
