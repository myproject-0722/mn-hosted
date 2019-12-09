package handler

import (
	"context"

	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	node "github.com/myproject-0722/mn-hosted/proto/node"
)

type Coin struct{}

func (s *Coin) Get(ctx context.Context, req *node.CoinListRequest, rsp *node.CoinListResponse) error {
	coinlist, err := dao.NodeDao.GetCoinList(db.Factoty.GetSession(), req.CurPage, req.PageSize)
	if err != nil {
		return err
	}

	rsp.Rescode = 200
	coinlistLength := len(coinlist)
	for i := 0; i < coinlistLength; i++ {
		v := coinlist[i]
		item := new(node.CoinItem)
		item.CoinName = v.CoinName
		item.MNRequired = v.MNRequired
		item.DPrice = float64(v.DPrice) / 100
		item.MPrice = float64(v.MPrice) / 100
		item.YPrice = float64(v.YPrice) / 100
		item.Volume = v.Volume
		item.Roi = v.Roi
		item.MonthlyIncome = v.MonthlyIncome
		item.MNHosted = v.MNHosted
		rsp.Coinlist = append(rsp.Coinlist, item)
	}

	return nil
}

func (s *Coin) GetCoinItem(ctx context.Context, req *node.CoinItemRequest, rsp *node.CoinItemResponse) error {
	v, err := dao.NodeDao.GetCoinItem(db.Factoty.GetSession(), req.CoinName)
	if err != nil {
		return err
	}

	rsp.Rescode = 200
	rsp.Coin = new(node.CoinItem)
	rsp.Coin.CoinName = v.CoinName
	rsp.Coin.MNRequired = v.MNRequired
	rsp.Coin.DPrice = float64(v.DPrice / 100)
	rsp.Coin.MPrice = float64(v.MPrice / 100)
	rsp.Coin.YPrice = float64(v.YPrice / 100)
	rsp.Coin.Volume = v.Volume
	rsp.Coin.Roi = v.Roi
	rsp.Coin.MonthlyIncome = v.MonthlyIncome
	rsp.Coin.MNHosted = v.MNHosted

	return nil
}
