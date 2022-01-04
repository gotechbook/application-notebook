package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	V *viper.Viper
	L *zap.Logger
)
