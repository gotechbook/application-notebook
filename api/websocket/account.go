package websocket

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gotechbook/application-notebook/common"
	"github.com/gotechbook/application-notebook/lib/cache"
	"github.com/gotechbook/application-notebook/logger"
	"github.com/gotechbook/application-notebook/servers/ws"
)

func PingController(c *ws.Client, seq string, message []byte) (code uint32, msg string, data interface{}) {
	logger.Info("PingController", fmt.Sprintf("addr:%s - seq:%s - msg:%s", c.Addr, seq, string(message)))
	code = common.OK
	data = "pong"
	return
}

func LoginController(c *ws.Client, seq string, message []byte) (code uint32, msg string, data interface{}) {
	var (
		currentTime = uint64(time.Now().Unix())
		req         = &common.LoginReq{}
		userOnline  = &ws.UserOnline{}
		err         error
	)
	code = common.OK
	if err = json.Unmarshal(message, req); err != nil {
		code = common.ParameterIllegal
		logger.Error("LoginController json Unmarshal", err, seq)
		return
	}
	if req.UserId == "" {
		code = common.UnauthorizedUserId
		logger.Error("LoginController 非法的用户", err, seq)
		return
	}
	if !ws.InAppIds(req.AppId) {
		code = common.Unauthorized
		logger.Error("LoginController 不支持的平台", err, seq)
		return
	}
	if c.IsLogin() {
		code = common.OperationFailure
		logger.Error("LoginController 用户已经登录", err, seq)
		return
	}
	// 存储数据
	c.Login(req.AppId, strings.ToLower(req.UserId), currentTime)
	userOnline = ws.UserLogin(ws.ServerIp, ws.ServerPort, req.AppId, strings.ToLower(req.UserId), c.Addr, currentTime)
	err = cache.SetUserOnlineInfo(c.GetKey(), userOnline)
	if err != nil {
		code = common.ServerError
		logger.Error("LoginController 用户登录 SetUserOnlineInfo", err, seq)
		return
	}
	ws.ClientManager.Login <- &ws.Login{
		AppId:  req.AppId,
		UserId: strings.ToLower(req.UserId),
		Client: c,
	}
	logger.Info("LoginController 用户登录 成功", fmt.Sprintf("seq:%s - addr:%s - userId:%s", seq, c.Addr, req.UserId))
	return
}
