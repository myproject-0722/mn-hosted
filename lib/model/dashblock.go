package model

import "time"

//
type DashBlock struct {
	Id           int64     `json:"id"`           // id
	BlockID      int32     `json:"blockid"`      // 货币名称
	MNPayee      string    `json:"mnpayee"`      // 货币名称
	PoolPubkey   string    `json:"poolpubkey"`   //
	BlockValue   int32     `json:"dprice"`       //
	MNValue      int32     `json:"mnvalue"`      //
	PoolValue    int32     `json:"poolvalue"`    //
	IsSuperBlock bool      `json:"issuperblock"` //
	CreateTime   time.Time `json:"createtime"`   //
	UpdateTime   time.Time `json:"updatetime"`   // 更新时间
}
