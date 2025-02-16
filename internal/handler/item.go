package handler

import (
	"net/http"
	"strconv"

	itemdto "github.com/fnxr21/invoice-system/internal/dto/item"
	resultdto "github.com/fnxr21/invoice-system/internal/dto/result"
	"github.com/fnxr21/invoice-system/internal/service"
	errorhandler "github.com/fnxr21/invoice-system/pkg/error"
	"github.com/labstack/echo/v4"
)

type handlerItem struct {
	ItemService service.ItemService
}

func HandlerItem(ItemService service.ItemService) *handlerItem {
	return &handlerItem{ItemService}
}

func (h *handlerItem) CreateItem(c echo.Context) error {

	var req itemdto.ItemRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	response, err := h.ItemService.CreateItem(&req)
	if err != nil {
		return errorhandler.ErrorHandler(err, http.StatusInternalServerError, c)
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}

func (h *handlerItem) ListItem(c echo.Context) error {
	response, _ := h.ItemService.ListItem()

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}
func (h *handlerItem) GetItemByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return errorhandler.ErrorHandler(err, http.StatusInternalServerError, c)
	}

	response, err := h.ItemService.GetItemByID(uint(id))
	if err != nil {
		return errorhandler.ErrorHandler(err, http.StatusInternalServerError, c)
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}
