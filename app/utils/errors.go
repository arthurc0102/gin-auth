package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/mattn/go-sqlite3"

	validator "gopkg.in/go-playground/validator.v8"
)

// FieldValidatorError change validator.FieldError to string
func FieldValidatorError(e *validator.FieldError) string {
	switch e.Tag {
	case "required":
		return fmt.Sprint("This field is required")
	case "notspace":
		return fmt.Sprint("This field can't be space")
	case "max":
		return fmt.Sprintf("This field can't be bigger or longer than %s", e.Param)
	case "min":
		return fmt.Sprintf("This field must be bigger or longer than %s", e.Param)
	case "email":
		return fmt.Sprint("Invalid email format")
	case "len":
		return fmt.Sprintf("This field must be %s characters long", e.Param)
	case "choices":
		choiceList := strings.Split(e.Param, " ")
		choiceListString := fmt.Sprintf(
			"%s or %s",
			strings.Join(choiceList[:len(choiceList)-1], ", "),
			choiceList[len(choiceList)-1],
		)

		return fmt.Sprintf("This field should be one of %s", choiceListString)
	case "eqfield":
		return fmt.Sprintf("This field should be equal to field %s", e.Param)
	}

	return fmt.Sprint("This field is not valid")
}

// HandleError handle binding error
func HandleError(object interface{}, err error) map[string]string {
	errors := make(map[string]string)
	val := reflect.ValueOf(object)

	switch reflect.TypeOf(err) {
	case reflect.TypeOf(validator.ValidationErrors{}):
		for _, e := range err.(validator.ValidationErrors) {
			field, _ := val.Type().FieldByName(e.Field)
			jsonKey := field.Tag.Get("json")

			if jsonKey == "" {
				jsonKey = e.Field
			}

			errors[jsonKey] = FieldValidatorError(e)
		}
	case reflect.TypeOf(&json.SyntaxError{}):
		errors["nonField"] = "Parse json error"
	case reflect.TypeOf(&json.UnmarshalTypeError{}):
		err := err.(*json.UnmarshalTypeError)
		errors[err.Field] = fmt.Sprintf("This field required a %s type value, not %s", err.Type, err.Value)
	case reflect.TypeOf(sqlite3.Error{}):
		errorSlice := strings.Split(err.Error(), ":")
		fieldInfo := strings.Split(errorSlice[len(errorSlice)-1], ".")
		field := strings.TrimSpace(fieldInfo[len(fieldInfo)-1])
		message := strings.TrimSpace(strings.Split(errorSlice[0], "constraint failed")[0])
		errors[field] = fmt.Sprintf("This field need to be %s", strings.ToLower(message))
	default:
		errors["nonField"] = err.Error()
	}

	return errors
}
