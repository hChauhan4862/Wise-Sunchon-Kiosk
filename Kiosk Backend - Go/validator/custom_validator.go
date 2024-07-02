package custom_validation

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func hhmmFormat(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^(((0|1)[0-9])|(2[0-3]))([0-5][0-9])$`).MatchString(fl.Field().String())
}
func hhmmssFormat(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^(((0|1)[0-9])|(2[0-3]))([0-5][0-9])([0-5][0-9])$`).MatchString(fl.Field().String())
}

func HcValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("HHMM", hhmmFormat)
		v.RegisterValidation("HHMMSS", hhmmssFormat)
	}
}
