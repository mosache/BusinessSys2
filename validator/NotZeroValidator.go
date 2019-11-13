package validator

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

func NotZero (v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool{
		if val,ok :=field.Interface().(int);ok {
			if val == 0 {
				return false
			}
		}else {
			return false
		}
		return true
}