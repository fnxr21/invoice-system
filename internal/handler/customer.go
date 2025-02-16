package handler

import (
	"net/http"
	"strconv"

	customerdto "github.com/fnxr21/invoice-system/internal/dto/customer"
	resultdto "github.com/fnxr21/invoice-system/internal/dto/result"
	"github.com/fnxr21/invoice-system/internal/service"
	errorhandler "github.com/fnxr21/invoice-system/pkg/error"
	"github.com/labstack/echo/v4"
)

type handlerCustomer struct {
	CustomerService service.CustomerService
}

func HandlerCustomer(CustomerService service.CustomerService) *handlerCustomer {
	return &handlerCustomer{CustomerService}
}

func (h *handlerCustomer) CreateCustomer(c echo.Context) error {

	var req customerdto.CustomerRequest
	// Bind JSON payload to struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	response, err := h.CustomerService.CreateCustomer(&req)
	if err != nil {
		return errorhandler.ErrorHandler(err, http.StatusInternalServerError, c)
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}

func (h *handlerCustomer) ListCustomer(c echo.Context) error {
	response, err := h.CustomerService.ListCustomer()
	if err != nil {
		return errorhandler.ErrorHandler(err, http.StatusInternalServerError, c)
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}
func (h *handlerCustomer) GetCustomerByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return errorhandler.ErrorHandler(err, http.StatusInternalServerError, c)
	}
	response, err := h.CustomerService.GetCustomerByID(uint(id))
	if err != nil {
		return errorhandler.ErrorHandler(err, http.StatusInternalServerError, c)
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}
