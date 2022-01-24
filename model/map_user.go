package model

import (
	"context"
	"time"
)

type MapUser struct {
	ID        string    `json:"id"`
	MapId     string    `json:"blockId"`   // 地图ID
	Status    int64     `json:"status"`    // 状态
	IsDel     bool      `json:"isDel"`     // 是否删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 修改时间
}

type MapUserDto struct {
	ID        string    `json:"id"`
	MapId     string    `json:"blockId"`   // 地图ID
	Metadata  Map       `json:"metadata"`  // 地图详情
	Status    int64     `json:"status"`    // 状态
	IsDel     bool      `json:"isDel"`     // 是否删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 修改时间
}

func (e *MapUserDto) To() *MapUser {
	return &MapUser{
		ID:        e.ID,
		MapId:     e.MapId,
		Status:    e.Status,
		IsDel:     e.IsDel,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (e *MapUser) ToDto(meta Map) *MapUserDto {
	return &MapUserDto{
		ID:        e.ID,
		MapId:     e.MapId,
		Metadata:  meta,
		Status:    e.Status,
		IsDel:     e.IsDel,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func MapUserGetOne(ctx context.Context, id string) (rst MapUserDto, err error) { return }

func MapUserGetOneToken(ctx context.Context, tokenId string) (rst MapUserDto, err error) { return }

func MapUserGetList(ctx context.Context, userId string) (rst []MapUserDto, err error) { return }

func MapUserGetPage(ctx context.Context, page uint64, pageSize uint64) (rst []MapUserDto, err error) {
	return
}
