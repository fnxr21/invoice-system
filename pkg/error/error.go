package errorhandler

import (
	// "net/http"

	// "encoding/json"
	"errors"
	// "fmt"
	// "fmt"

	// "net/http"

	resultdto "github.com/fnxr21/invoice-system/internal/dto/result"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var ErrCustomerExists = errors.New("customer already exists")

func ErrorHandler(err error, http int, c echo.Context) error {
	if errors.Is(err, ErrCustomerExists) {
		return c.JSON(http, resultdto.ErrorResult{Errors: "Customer already exists"})
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http, resultdto.ErrorResult{Errors: "Data not found"})
	}

	return c.JSON(http, resultdto.ErrorResult{Errors: err})
}

// default
///////
// func ErrorHandler(err error, http int, c echo.Context) error {
// 	// report, ok := err.(*echo.HTTPError)
// 	// if !ok {
// 	// 	report = echo.NewHTTPError(http, err.Error())
// 	// }
// 	fmt.Println(err,"chec")
// 	if errors.Is(err, errors.New("customer already exists check")) {
// 		return c.JSON(http, resultdto.ErrorResult{Errors: "Data not found"})
// 	}
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return c.JSON(http, resultdto.ErrorResult{Errors: "Data not found"})
// 	}

// 	// c.Logger().Error()
// 	// fmt.Println(report, "sss")
// 	return c.JSON(http, resultdto.ErrorResult{Errors: err})
// }
