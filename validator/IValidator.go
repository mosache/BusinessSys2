package validator

import "gopkg.in/go-playground/validator.v8"

type IValidator interface {
	GetError(err validator.ValidationErrors) string
}
