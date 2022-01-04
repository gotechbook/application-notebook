// +build !windows

package log

import (
	"os"
	"path"
	"time"

	"github.com/gotechbook/application-notebook/global"
	flog "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := flog.New(
		path.Join(global.V.GetString("zap.director"), "%Y-%m-%d.log"),
		flog.WithLinkName(global.V.GetString("zap.link-name")),
		flog.WithMaxAge(7*24*time.Hour),
		flog.WithRotationTime(24*time.Hour),
	)
	if global.V.GetBool("zap.log-in-console") {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
