package errorhandler

import (
	// "net/http"

	// "encoding/json"
	"errors"
	// "fmt"

	// "net/http"

	resultdto "github.com/fnxr21/invoice-system/internal/dto/result"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ErrorHandler(err error, http int, c echo.Context) error {
	// report, ok := err.(*echo.HTTPError)
	// if !ok {
	// 	report = echo.NewHTTPError(http, err.Error())
	// }
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http, resultdto.ErrorResult{Errors: "Data not found"})
	}

	// c.Logger().Error()
	// fmt.Println(report, "sss")
	return c.JSON(http, resultdto.ErrorResult{Errors: err})
}
