package handler

import (
	"net/http"
	"strconv"

	customerdto "github.com/fnxr21/invoice-system/internal/dto/customer"
	resultdto "github.com/fnxr21/invoice-system/internal/dto/result"
	"github.com/fnxr21/invoice-system/internal/service"
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

	response, _ := h.CustomerService.CreateCustomer(&req)

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}

func (h *handlerCustomer) ListCustomer(c echo.Context) error {
	response, _ := h.CustomerService.ListCustomer()

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}
func (h *handlerCustomer) GetCustomerByID(c echo.Context) error {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	response, _ := h.CustomerService.GetCustomerByID(uint(id))

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Data: response})
}
