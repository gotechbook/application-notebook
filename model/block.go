package model

import (
	"context"
	"time"
)

type Block struct {
	ID             string    `json:"id"`
	Level          uint64    `json:"level"`
	Income         float64   `json:"income"`         // 收益
	OccupationRate float64   `json:"occupationRate"` // 占领成功率
	SynthesisRate  float64   `json:"synthesisRate"`  // 合成成功率
	SynthesisFee   float64   `json:"synthesisFee"`   // 合成手续费
	IsDel          bool      `json:"isDel"`          // 是否删除
	CreatedAt      time.Time `json:"createdAt"`      // 创建时间
	UpdatedAt      time.Time `json:"updatedAt"`      // 修改时间
}

func BlockCreate(ctx context.Context, data Block) (id string, err error) {
	return
}

func BlockDelete(ctx context.Context, id string) (b bool, err error) {
	return
}

func BlockGetOne(ctx context.Context, id string) (rst Block, err error) {
	return
}

func BlockGetList(ctx context.Context) (rst []Block, err error) {
	return
}

func BlockGetPage(ctx context.Context, page uint64, pageSize uint64) (rst []Block, err error) {
	return
}
