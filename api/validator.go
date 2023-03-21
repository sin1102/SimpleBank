package api

import (
	"learn/banking/utils"

	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return utils.IsSupportCurrency(currency)
	}
	return false
}
