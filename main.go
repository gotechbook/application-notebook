package main

import (
	"github.com/gotechbook/application-notebook/lib/log"
	"github.com/gotechbook/application-notebook/lib/viper"
)

func main() {
	initConfig()
}

func initConfig() {
	viper.ViperInit()
	log.ZapInit()
}
