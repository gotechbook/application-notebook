package log

import (
	"os"
	"path"
	"time"

	"github.com/gotechbook/application-notebook/config"
	log "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := log.New(
		path.Join(config.V.GetString("zap.director"), "%Y-%m-%d.log"),
		log.WithMaxAge(7*24*time.Hour),
		log.WithRotationTime(24*time.Hour),
	)
	if config.V.GetBool("zap.log-in-console") {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
