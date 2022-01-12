package common

type PageReq struct {
	Page     int64 `json:"page" form:"page"`         // 页码
	PageSize int64 `json:"pageSize" form:"pageSize"` // 每页大小
}

type IdReq struct {
	Id string `json:"id" form:"id"` // id
}

type HashReq struct {
	Hash string `json:"hash" form:"hash"` // hash
}

// 登录请求数据
type LoginReq struct {
	AppId  uint32 `json:"appId,omitempty"`
	UserId string `json:"userId,omitempty"`
}
