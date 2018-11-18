package validators

import (
	"reflect"
	"strings"

	validator "gopkg.in/go-playground/validator.v8"
)

// Choices validator
func Choices(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	args := strings.Split(param, " ")
	value := field.Interface().(string)

	for _, arg := range args {
		if arg == value {
			return true
		}
	}

	return false
}
