package model

import "time"

//
type User struct {
	Id            int64     `json:"id"`            // id
	Account       string    `json:"Account"`       // account
	Passwd        string    `json:"passwd"`        //
	WalletAddress string    `json:"walletaddress"` //
	CreateTime    time.Time `json:"createtime"`    //
	UpdateTime    time.Time `json:"updatetime"`    // 更新时间
}
