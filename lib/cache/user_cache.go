package cache

import (
	"encoding/json"
	"fmt"

	"github.com/gotechbook/application-notebook/lib/redislib"
	"github.com/gotechbook/application-notebook/logger"
	"github.com/gotechbook/application-notebook/servers/ws"
)

const (
	userOnlinePrefix    = "ws:user:online:" // 用户在线状态
	userOnlineCacheTime = 24 * 60 * 60
)

func getUserOnlineKey(userKey string) string {
	return fmt.Sprintf("%s%s", userOnlinePrefix, userKey)
}

func GetUserOnlineInfo(userKey string) (userOnline *ws.UserOnline, err error) {
	redisClient := redislib.GetClient()
	key := getUserOnlineKey(userKey)
	data, err := redisClient.Get(key).Bytes()
	if err != nil {
		logger.Error("GetUserOnlineInfo", err, userKey)
		return
	}
	userOnline = &ws.UserOnline{}
	err = json.Unmarshal(data, userOnline)
	if err != nil {
		logger.Error("GetUserOnlineInfo json Unmarshal ", err, userKey)
		return
	}
	logger.Info("获取用户在线数据", fmt.Sprintf("userKey:%s - loginTime:%d - heartbeatTime:%d - ip:%s ", userKey, userOnline.LoginTime, userOnline.HeartbeatTime, userOnline.AccIp))
	return
}

// 设置用户在线数据
func SetUserOnlineInfo(userKey string, userOnline *ws.UserOnline) (err error) {
	redisClient := redislib.GetClient()
	key := getUserOnlineKey(userKey)
	valueByte, err := json.Marshal(userOnline)
	if err != nil {
		logger.Error("SetUserOnlineInfo json Marshal", err, userKey)
		return
	}
	_, err = redisClient.Do("setEx", key, userOnlineCacheTime, string(valueByte)).Result()
	if err != nil {
		logger.Error("SetUserOnlineInfo redisClient Do", err, userKey)
	}
	return
}
