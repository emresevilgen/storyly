package validator

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"storyly/pkg/errors"
)

type (
	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return errors.CreateError(http.StatusBadRequest, err.Error())
	}
	return nil
}
