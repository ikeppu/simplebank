package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/ikeppu/simplebank/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	currency, ok := fl.Field().Interface().(string)

	if ok {
		return util.IsSupportedCurrency(currency)
	}

	return false
}
