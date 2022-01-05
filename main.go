package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gotechbook/application-notebook/config"
	"github.com/gotechbook/application-notebook/lib/log"
	"github.com/gotechbook/application-notebook/lib/viper"
	"github.com/gotechbook/application-notebook/logger"
	"github.com/gotechbook/application-notebook/servers/ws"
)

func main() {
	initConfig()
	go ws.StartWebSocket()
	httpPort := strconv.FormatInt(int64(config.C.Server.HttpPort), 10)
	router := gin.Default()
	http.ListenAndServe(":"+httpPort, router)
}

func initConfig() {
	config.V = viper.Viper()
	logger.L = log.Zap()

}
