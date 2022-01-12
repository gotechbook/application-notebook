package ws

import (
	"fmt"
	"time"

	"github.com/gotechbook/application-notebook/logger"
)

const (
	heartbeatTimeout = 3 * 60 // 用户心跳超时时间
)

type Login struct {
	AppId  uint32
	UserId string
	Client *Client
}

// 读取客户端数据
func (l *Login) GetKey() (key string) {
	key = GetUserKey(l.AppId, l.UserId)
	return
}

// 用户在线状态
type UserOnline struct {
	AccIp         string `json:"accIp"`         // acc Ip
	AccPort       string `json:"accPort"`       // acc 端口
	AppId         uint32 `json:"appId"`         // appId
	UserId        string `json:"userId"`        // 用户Id
	ClientIp      string `json:"clientIp"`      // 客户端Ip
	ClientPort    string `json:"clientPort"`    // 客户端端口
	LoginTime     uint64 `json:"loginTime"`     // 用户上次登录时间
	HeartbeatTime uint64 `json:"heartbeatTime"` // 用户上次心跳时间
	LogOutTime    uint64 `json:"logOutTime"`    // 用户退出登录的时间
	Qua           string `json:"qua"`           // qua
	DeviceInfo    string `json:"deviceInfo"`    // 设备信息
	IsLogoff      bool   `json:"isLogoff"`      // 是否下线
}

/**********************  数据处理  *********************************/

// 用户登录
func UserLogin(accIp, accPort string, appId uint32, userId string, addr string, loginTime uint64) (userOnline *UserOnline) {
	userOnline = &UserOnline{
		AccIp:         accIp,
		AccPort:       accPort,
		AppId:         appId,
		UserId:        userId,
		ClientIp:      addr,
		LoginTime:     loginTime,
		HeartbeatTime: loginTime,
		IsLogoff:      false,
	}
	return
}

// 用户心跳
func (u *UserOnline) Heartbeat(currentTime uint64) {
	u.HeartbeatTime = currentTime
	u.IsLogoff = false
	return
}

// 用户退出登录
func (u *UserOnline) LogOut() {
	currentTime := uint64(time.Now().Unix())
	u.LogOutTime = currentTime
	u.IsLogoff = true
	return
}

/**********************  数据操作  *********************************/

// 用户是否在线
func (u *UserOnline) IsOnline() (online bool) {
	if u.IsLogoff {
		return
	}
	currentTime := uint64(time.Now().Unix())
	if u.HeartbeatTime < (currentTime - heartbeatTimeout) {
		logger.Error("IsOnline 心跳超时", nil, fmt.Sprintf("appId:%d - userId:%s - heartbeatTime:%d", u.AppId, u.UserId, u.HeartbeatTime))
		return
	}
	if u.IsLogoff {
		logger.Error("IsOnline 用户已经下线", nil, fmt.Sprintf("appId:%d - userId:%s - heartbeatTime:%d", u.AppId, u.UserId, u.HeartbeatTime))
		return
	}
	return true
}

// 用户是否在本台机器上
func (u *UserOnline) UserIsLocal(localIp, localPort string) (result bool) {
	if u.AccIp == localIp && u.AccPort == localPort {
		result = true
		return
	}
	return
}
