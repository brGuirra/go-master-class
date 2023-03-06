package api

import (
	"github.com/brGuirra/simple-bank/utils"
	"github.com/go-playground/validator/v10"
)

var isValidCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportedCurrency(currency)
	}

	return false
}
