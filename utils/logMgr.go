package utils

import (
	"github.com/spf13/viper"
)

var (
	LogMgr       *Logger
	GinLoggerMgr *Logger
)

func NewLogMgr() {
	LogMgr = NewLogger(viper.GetString("logDir"))
	GinLoggerMgr = NewLogger(viper.GetString("ginLogDir"))
}
