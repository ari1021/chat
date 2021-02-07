package validation

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func ValidateEcho(e *echo.Echo) *echo.Echo {
	e.Validator = &customValidator{validator: validator.New()}
	return e
}
