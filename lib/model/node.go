package model

import "time"

// User 账户
type Coin struct {
	Id            int32     `json:"id"`            // 用户id
	CoinName      string    `json:"coinname"`      // 手机号
	MNRequired    int32     `json:"mnrequired"`    // 昵称
	MNPrice       int32     `json:"mnprice"`       // 性别，1:男；2:女
	Volume        int32     `json:"volume"`        // 用户头像
	Roi           int32     `json:"roi"`           // 密码
	MonthlyIncome int32     `json:"monthlyincome"` // 创建时间
	MNHosted      int32     `json:"mnhosted"`      // 密码
	CreateTime    time.Time `json:"createtime"`    //
	UpdateTime    time.Time `json:"updatetime"`    // 更新时间
}
