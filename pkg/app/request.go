package app

import (
	"github.com/astaxie/beego/validation"
	"go-email/pkg/logging"
	"go.uber.org/zap"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.AppLogger.Error("validation error", zap.Error(err))
	}

	return
}
