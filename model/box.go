package model

import "time"

type Box struct {
	ID        string    `json:"id"`
	Type      int64     `json:"type"`      // 类型
	IsDel     bool      `json:"isDel"`     // 是否删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 修改时间
}
