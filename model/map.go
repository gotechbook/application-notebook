package model

import "time"

type Map struct {
	ID        string    `json:"id"`
	TokenId   string    `json:"tokenId"`   // tokenId
	X         uint64    `json:"x"`         // 最大X轴
	Y         uint64    `json:"y"`         // 最大Y轴
	Remark    string    `json:"remark"`    // 别名
	Status    int64     `json:"status"`    // 状态
	IsDel     bool      `json:"isDel"`     // 是否删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 修改时间
}
