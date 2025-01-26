package validators

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

var ValidatePrice validator.Func = func (fl validator.FieldLevel) bool {
	pattern := "^\\d+\\.\\d{2}$"

	matched, _ := regexp.MatchString(pattern, fl.Field().String())
	return matched
}

var ValidateRetailer validator.Func = func (fl validator.FieldLevel) bool {
	pattern := "^[\\w\\s\\-&]+$"

	matched, _ := regexp.MatchString(pattern, fl.Field().String())
	return matched
}

var ValidateDate validator.Func = func (fl validator.FieldLevel) bool {
	const layout = "2006-01-02" // Date format: YYYY-mm-dd

	_, err := time.Parse(layout, fl.Field().String())
	return err == nil
}

var ValidateTime validator.Func = func (fl validator.FieldLevel) bool {
	const layout = "15:04" // Time format: HH:mm

	_, err := time.Parse(layout, fl.Field().String())
	return err == nil
}

var ValidateDescription validator.Func = func (fl validator.FieldLevel) bool {
	pattern := "^[\\w\\s\\-]+$"

	matched, _ := regexp.MatchString(pattern, fl.Field().String())
	return matched
}