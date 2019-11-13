package validator

import (
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

//注册Validators
func InitValidators() {
	if v,ok := binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("notZero",NotZero)
	}
}