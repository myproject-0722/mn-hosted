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
type CoinsPrice struct {
	Id         int64     `json:"id"`         // id
	CoinName   string    `json:"coinname"`   // 货币名称
	Price      int32     `json:"price"`      //
	CreateTime time.Time `json:"createtime"` //
	UpdateTime time.Time `json:"updatetime"` // 更新时间
}

//
type CoinRewards struct {
	CoinName string  `json:"coinname"` // 货币名称
	Rewards  float64 `json:"rewards"`  //
	MNCount  int32   `json:"mncount"`
}

//
type Masternode struct {
	Id           int64     `json:"id"`           // id
	CoinName     string    `json:"coinname"`     // 货币名称
	MNKey        string    `json:"mnkey"`        // key
	MNPayee      string    `json:"mnpayee"`      // key
	UserID       int64     `json:"userid"`       //
	OrderID      int64     `json:"orderid"`      //
	Vps          string    `json:"vps"`          // vps
	DockerID     string    `json:"dockerid"`     //dockerid
	Status       int32     `json:"status"`       //
	SyncStatus   int32     `json:"syncstatus"`   //
	MNStatus     int32     `json:"mnstatus"`     //
	SyncStatusEx string    `json:"syncstatusex"` //
	MNStatusEx   string    `json:"mnstatusex"`   //
	IsNotify     bool      `json:"isnotify"`     //
	Earn         int64     `json:"earn"`         //
	CreateTime   time.Time `json:"createtime"`   //
	ExpireTime   time.Time `json:"expiretime"`   //
	UpdateTime   time.Time `json:"updatetime"`   // 更新时间
}

//
type MasternodeCount struct {
	Count int32 `json:"count"`
	Earns int64 `json:"earns"` // id
}

//
type Node struct {
	Id         int64     `json:"id"`         // id
	CoinName   string    `json:"coinname"`   // 货币名称
	UserID     int64     `json:"userid"`     //
	VpsID      int64     `json:"vpsid"`      //
	OrderID    int64     `json:"orderid"`    //
	PublicIP   string    `json:"publicip"`   //
	Port       int32     `json:"port"`       // vps
	State      string    `json:"state"`      // state
	Status     string    `json:"status"`     // status
	CreateTime time.Time `json:"createtime"` //
	UpdateTime time.Time `json:"updatetime"` // 更新时间
}
