package model

import "time"

//
type Coin struct {
	Id            int64     `json:"id"`            // id
	CoinName      string    `json:"coinname"`      // 货币名称
	MNRequired    int32     `json:"mnrequired"`    //
	DPrice        int32     `json:"dprice"`        //
	MPrice        int32     `json:"mprice"`        //
	YPrice        int32     `json:"yprice"`        //
	Volume        int32     `json:"volume"`        //
	Roi           int32     `json:"roi"`           //
	MonthlyIncome int32     `json:"monthlyincome"` //
	MNHosted      int32     `json:"mnhosted"`      //
	CreateTime    time.Time `json:"createtime"`    //
	UpdateTime    time.Time `json:"updatetime"`    // 更新时间
}

//
type Masternode struct {
	Id         int64     `json:"id"`         // id
	CoinName   string    `json:"coinname"`   // 货币名称
	MNKey      string    `json:"mnkey"`      // key
	UserID     int64     `json:"userid"`     //
	OrderID    int64     `json:"orderid"`    //
	Vps        string    `json:"vps"`        // vps
	DockerID   string    `json:"dockerid"`   //dockerid
	Status     int32     `json:"status"`     //
	SyncStatus int32     `json:"syncstatus"` //
	CreateTime time.Time `json:"createtime"` //
	ExpireTime time.Time `json:"expiretime"` //
	UpdateTime time.Time `json:"updatetime"` // 更新时间
}
