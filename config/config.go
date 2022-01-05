package config

import (
	"github.com/spf13/viper"
)

var (
	V *viper.Viper
	C Config
)

type Config struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
}
