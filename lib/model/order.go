package model

import "time"

//
type Order struct {
	Id         int64     `json:"id"` // id
	UserID     int64     `json:"userid"`
	CoinName   string    `json:"Coinname"` // account
	MNName     string    `json:"mnname"`   //
	MNKey      string    `json:"mnkey"`    //
	TimeType   int32     `json:"timetype"` //
	Price      int32     `json:"price"`
	TxID       string    `json:"txid"`
	TxIndex    int32     `json:"txindex"` //
	IsRenew    int32     `json:"isrenew"`
	Status     int32     `json:"status"`
	CreateTime time.Time `json:"createtime"` //
	UpdateTime time.Time `json:"updatetime"` // 更新时间
}

//
type OrderInfo struct {
	Num    int32   `json:"num"` //
	Payout float64 `json:"payout"`
}
