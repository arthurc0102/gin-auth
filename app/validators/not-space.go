package validators

import (
	"reflect"
	"strings"

	validator "gopkg.in/go-playground/validator.v8"
)

// NotSpace check value can't be space
func NotSpace(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	value := field.Interface().(string)
	return len(strings.TrimSpace(value)) > 0
}
