package ws

import (
	"runtime/debug"

	"github.com/gorilla/websocket"
	"github.com/gotechbook/application-notebook/logger"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60
)

// 用户连接
type Client struct {
	Addr          string          // 客户端地址
	Socket        *websocket.Conn // 用户连接
	Send          chan []byte     // 待发送的数据
	AppId         uint32          // 登录的平台Id app/web/ios
	UserId        string          // 用户Id，用户登录以后才有
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
}

// 初始化
func NewClient(addr string, socket *websocket.Conn, firstTime uint64) *Client {
	return &Client{
		Addr:          addr,
		Socket:        socket,
		Send:          make(chan []byte, 100),
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
	}
}

// 读取客户端数据
func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("读取客户端数据 错误stop", nil, debug.Stack())
		}
	}()

	defer func() {
		logger.Error("读取客户端数据 关闭send", nil, debug.Stack())
		close(c.Send)
	}()
	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			logger.Error("读取客户端数据 错误", err, msg)
			return
		}
		logger.Info("读取客户端数据 处理", msg)
		ProcessData(c, msg)
	}
}

// 向客户端写数据
func (c *Client) write() {
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				logger.Info("Client发送数据 关闭连接", c.Addr)
				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

// 关闭
func (c *Client) close() {
	close(c.Send)
}

// 发送消息
func (c *Client) SendMsg(msg []byte) {
	if c == nil {
		return
	}
	c.Send <- msg
}

// 心跳检测
func (c *Client) Heartbeat(currentTime uint64) {
	c.HeartbeatTime = currentTime
}

// 心跳超时
func (c *Client) IsHeartbeatTimeout(currentTime uint64) bool {
	if c.HeartbeatTime+heartbeatExpirationTime >= currentTime {
		return true
	}
	return false
}

// 登录
func (c *Client) Login(appId uint32, userId string, loginTime uint64) {
	c.AppId = appId
	c.UserId = userId
	c.LoginTime = loginTime
	c.Heartbeat(loginTime)
}

