package config

import (
	"log"

	"github.com/arthurc0102/gin-auth/app/validators"
	"github.com/gin-gonic/gin/binding"
	validator "gopkg.in/go-playground/validator.v8"
)

// RegisterValidators register validators
func RegisterValidators() {
	v, ok := binding.Validator.Engine().(*validator.Validate)

	if !ok {
		log.Fatalln("Validators can't register!")
		return
	}

	v.RegisterValidation("notspace", validators.NotSpace)
	v.RegisterValidation("choices", validators.Choices)
}
