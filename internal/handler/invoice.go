package handler

import (
	"net/http"
	"strconv"

	invoicedto "github.com/fnxr21/invoice-system/internal/dto/invoice"
	resultdto "github.com/fnxr21/invoice-system/internal/dto/result"
	"github.com/fnxr21/invoice-system/internal/service"
	"github.com/labstack/echo/v4"
)

type handlerInvoice struct {
	InvoiceService service.InvoiceService
}

func HandlerInvoice(InvoiceService service.InvoiceService) *handlerInvoice {
	return &handlerInvoice{InvoiceService}
}

func (h *handlerInvoice) CreateInvoice(c echo.Context) error {

	var req invoicedto.InvoiceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":  "Invalid request payload",
			"detail": err.Error(),
		})
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	response, _ := h.InvoiceService.CreateInvoice(&req)

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}

func (h *handlerInvoice) IndexInvoice(c echo.Context) error {
	var req invoicedto.InvoiceIndexing
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":  "Invalid request payload",
			"detail": err.Error(),
		})
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	response, pagination, err := h.InvoiceService.IndexInvoice(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResultIndex{Data: response, Pagination: pagination})
}
func (h *handlerInvoice) GetInvoiceByID(c echo.Context) error {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	response, _ := h.InvoiceService.GetInvoiceByID(uint(id))

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}

func (h *handlerInvoice) UpdateInvoiceByID(c echo.Context) error {
	param_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	var req invoicedto.InvoiceRequestUpdate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":  "Invalid request payload",
			"detail": err.Error(),
		})
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	req.ID = uint(param_id)
	response, err := h.InvoiceService.UpdateInvoice(&req)

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}
