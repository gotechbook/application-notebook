package model

import (
	"context"
	"time"
)

type BlockUser struct {
	ID        string    `json:"id"`
	UserId    string    `json:"userId"`    // 用户ID
	TokenId   string    `json:"tokenId"`   // tokenId
	Color     string    `json:"color"`     // 颜色
	BlockId   string    `json:"configId"`  // 块ID
	Remark    string    `json:"remark"`    // 别名
	Status    int64     `json:"status"`    // 状态
	IsDel     bool      `json:"isDel"`     // 是否删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 修改时间
}

type BlockUserDto struct {
	ID        string    `json:"id"`
	UserId    string    `json:"userId"`    // 用户ID
	TokenId   string    `json:"tokenId"`   // tokenId
	Color     string    `json:"color"`     // 颜色
	BlockId   string    `json:"configId"`  // 块ID
	Remark    string    `json:"remark"`    // 别名
	Status    int64     `json:"status"`    // 状态
	IsDel     bool      `json:"isDel"`     // 是否删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 修改时间
	Metadata  Block     `json:"metadata"`  // 元数据
}

func (e *BlockUserDto) To() *BlockUser {
	return &BlockUser{
		ID:        e.ID,
		UserId:    e.UserId,
		TokenId:   e.TokenId,
		Color:     e.Color,
		BlockId:   e.BlockId,
		Remark:    e.Remark,
		Status:    e.Status,
		IsDel:     e.IsDel,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (e *BlockUser) ToDto(meta Block) *BlockUserDto {
	return &BlockUserDto{
		ID:        e.ID,
		UserId:    e.UserId,
		TokenId:   e.TokenId,
		Color:     e.Color,
		BlockId:   e.BlockId,
		Remark:    e.Remark,
		Status:    e.Status,
		IsDel:     e.IsDel,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Metadata:  meta,
	}
}

func BlockUserCreate(ctx context.Context, data BlockUser) (id string, err error) { return }

func BlockUserCreateTokenId(ctx context.Context, data BlockUser) (id string, err error) { return }

func BlockUserDelete(ctx context.Context, id string) (b bool, err error) { return }

func BlockUserDeleteTokenId(ctx context.Context, tokenId string) (b bool, err error) { return }

func BlockUserGetOne(ctx context.Context, id string) (rst BlockUserDto, err error) { return }

func BlockUserGetOneToken(ctx context.Context, tokenId string) (rst BlockUserDto, err error) { return }

func BlockUserGetList(ctx context.Context, userId string) (rst []BlockUserDto, err error) { return }

func BlockUserGetPage(ctx context.Context, page uint64, pageSize uint64) (rst []BlockUserDto, err error) {
	return
}
