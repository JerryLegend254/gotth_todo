package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validate *validator.Validate
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validate: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validate.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
