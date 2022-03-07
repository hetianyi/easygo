package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/hetianyi/easygo/logger"
)

var (
	validators = map[string]func(fl validator.FieldLevel) bool{}
)

// RegisterValidator 注册一个gin字段校验器
func RegisterValidator(validator func(fl validator.FieldLevel) bool) {
	mu.Lock()
	defer mu.Unlock()

	name := getFuncName(validator)
	if validators[name] != nil {
		panic("Validator register failed due to: Validator \"" + name + "\" already registered.")
	}
	logger.Debug("register Validator :: ", name)
	validators[name] = validator
}

func getValidatorByName(name string) func(fl validator.FieldLevel) bool {
	return validators[name]
}
