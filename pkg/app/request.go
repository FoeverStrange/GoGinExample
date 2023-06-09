package app

import (
	"GoGinExample/pkg/logging"

	"github.com/astaxie/beego/validation"
)

// 错误入log
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
