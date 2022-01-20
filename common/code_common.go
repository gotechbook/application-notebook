package common

const (
	OK                 = 200  // Success
	NotLoggedIn        = 1000 // 未登录
	ParameterIllegal   = 1001 // 参数不合法
	UnauthorizedUserId = 1002 // 非法的用户Id
	Unauthorized       = 1003 // 未授权
	OperationFailure   = 1004 // 操作失败
	ServerError        = 1005 // 系统错误
	RoutingNotExist    = 1010 // 路由不存在
)

// 根据code 获取信息
func GetCodeMessage(code uint32, msg string) (rst string) {
	codeMap := map[uint32]string{
		OK:                 "success",
		NotLoggedIn:        "NotLoggedIn",
		RoutingNotExist:    "routingNotExist",    // 路由不存在
		ParameterIllegal:   "parameterIllegal",   // 参数不合法
		UnauthorizedUserId: "unauthorizedUserId", // 非法的用户Id
		Unauthorized:       "unauthorized",       // 未授权
		OperationFailure:   "operationFailure",   // 操作失败
		ServerError:        "serverError",        // 系统错误
	}
	if msg == "" {
		if value, ok := codeMap[code]; ok {
			rst = value
		} else {
			rst = "未定义错误类型!"
		}
	}
	return
}
