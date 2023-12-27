package validations

import (
	"github.com/go-playground/validator/v10"
	"log"
	"regexp"
)

func IranianMobileNumberValidatopr(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)

	if !ok {
		return false
	}

	res, err := regexp.MatchString(`09(1[0-9]|3[1-9]|2[1-9])-?[0-9]{3}-?[0-9]{4}`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}
