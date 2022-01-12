package ws

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gotechbook/application-notebook/common"
	"github.com/gotechbook/application-notebook/logger"
)

type DisposeFunc func(c *Client, seq string, message []byte) (code uint32, msg string, data interface{})

var (
	handlers        = make(map[string]DisposeFunc)
	handlersRWMutex sync.RWMutex
)

// 注册
func Register(key string, value DisposeFunc) {
	handlersRWMutex.Lock()
	defer handlersRWMutex.Unlock()
	handlers[key] = value

	return
}

func getHandlers(key string) (value DisposeFunc, ok bool) {
	handlersRWMutex.RLock()
	defer handlersRWMutex.RUnlock()
	value, ok = handlers[key]
	return
}

func ProcessData(c *Client, msg []byte) {

	var (
		req     = &WsRequest{}
		reqData []byte
		resCode uint32
		resMsg  string
		resData interface{}
		err     error
	)
	logger.Info("ProcessData 处理数据", fmt.Sprintf("addr:%s - msg:%s", c.Addr, string(msg)))

	defer func() {
		if r := recover(); r != nil {
			logger.Error("ProcessData recover", nil, r)
		}
	}()

	err = json.Unmarshal(msg, req)
	if err != nil {
		logger.Error("ProcessData json Unmarshal", err, "数据不合法")
		c.SendMsg([]byte("数据不合法"))
		return
	}

	reqData, err = json.Marshal(req.Data)
	if err != nil {
		logger.Error("ProcessData json Marshal", err, "处理数据失败")
		c.SendMsg([]byte("处理数据失败"))
		return
	}

	logger.Info("ProcessData req", fmt.Sprintf("cmd:%s - addr:%s", req.Cmd, c.Addr))

	// 采用 map 注册的方式
	if v, ok := getHandlers(req.Cmd); ok {
		resCode, resMsg, resData = v(c, req.Seq, reqData)
	} else {
		resCode = common.RoutingNotExist
		logger.Error("ProcessData getHandlers 路由不存在", nil, fmt.Sprintf("cmd:%s - addr:%s", req.Cmd, c.Addr))
	}

	responseHead := NewResponseHead(req.Seq, req.Cmd, resCode, common.GetCodeMessage(resCode, resMsg), resData)

	headByte, err := json.Marshal(responseHead)
	if err != nil {
		logger.Error("ProcessData json Marshal 处理数据", err, fmt.Sprintf("cmd:%s - addr:%s", req.Cmd, c.Addr))
		return
	}
	c.SendMsg(headByte)

	logger.Info("ProcessData send", fmt.Sprintf("cmd:%s - addr:%s--appId:%d - userId:%s - code:%d", req.Cmd, c.Addr, c.AppId, c.UserId, resCode))

	return
}
