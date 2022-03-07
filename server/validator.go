package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/hetianyi/easygo/logger"
)

var (
	validators = map[string]func(fl validator.FieldLevel) bool{}
)

// RegisterValidator 注册一个gin字段校验器
func RegisterValidator(tagName string, validator func(fl validator.FieldLevel) bool) {
	mu.Lock()
	defer mu.Unlock()

	if validators[tagName] != nil {
		panic("Validator register failed due to: Validator \"" + tagName + "\" already registered.")
	}
	logger.Debug("register Validator :: ", tagName)
	validators[tagName] = validator
}

func getValidatorByName(name string) func(fl validator.FieldLevel) bool {
	return validators[name]
}
