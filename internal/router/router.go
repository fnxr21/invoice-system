package router

import (
	"github.com/labstack/echo/v4"
)

func RouterInit(r *echo.Group) {
	Customer(r)
	Item(r)
}
