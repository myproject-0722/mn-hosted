package handler

import (
	"context"

	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/model"
	"github.com/myproject-0722/mn-hosted/lib/pay"
	order "github.com/myproject-0722/mn-hosted/proto/order"
	log "github.com/sirupsen/logrus"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) Alipay(ctx context.Context, req *order.AlipayRequest, rsp *order.AlipayResponse) error {
	//查询coinname及价格
	v, err := dao.NodeDao.GetCoinItem(db.Factoty.GetSession(), req.CoinName)
	if err != nil {
		rsp.Rescode = 404
		return err
	}

	var price int32
	if req.TimeType == 1 {
		price = v.DPrice
	} else if req.TimeType == 2 {
		price = v.MPrice
	} else if req.TimeType == 3 {
		price = v.YPrice
	}
	log.Debug("Alipay coinname: ", req.CoinName, " price: ", price)

	//检查主节点是否存在
	masternode, err := dao.NodeDao.GetMasternode(db.Factoty.GetSession(), req.CoinName, req.MNKey)
	if err != nil || masternode != nil {
		rsp.Rescode = 401
		return err
	}

	//插入order表生成订单id
	var o model.Order
	o.UserID = req.UserID
	o.Coinname = req.CoinName
	o.TimeType = req.TimeType
	o.Price = price
	o.IsRenew = 1
	//o.TxID = req.TxID
	o.Status = 0

	orderid, err := dao.OrderDao.Insert(db.Factoty.GetSession(), &o)
	if err != nil {
		rsp.Rescode = 500
		log.Error("insert error: ", req.UserID, " ", req.CoinName, " ", req.TimeType, " ", price)
		return err
	}

	//生成支付网页
	payUrl, err := pay.WebPageAlipay(orderid, price)
	if err != nil {
		rsp.Rescode = 500
		log.Error("WebPageAlipay: ", req.UserID, " ", req.CoinName, " ", req.TimeType, " ", price, " orderid:", orderid)
	}

	rsp.Rescode = 200
	rsp.PayUrl = payUrl
	return nil
}

func (s *OrderService) New(ctx context.Context, req *order.NewRequest, rsp *order.NewResponse) error {
	var o model.Order
	o.UserID = req.UserID
	o.Coinname = req.Coinname
	o.TimeType = req.Timetype
	o.Price = req.Price
	o.IsRenew = req.IsRenew
	o.TxID = req.TxID
	o.Status = 0

	id, err := dao.OrderDao.Insert(db.Factoty.GetSession(), &o)
	if err != nil {
		rsp.Rescode = 500
		log.Error("insert error: ", req.UserID, " ", req.Coinname, " ", req.Timetype, " ", req.Price)
		return nil
	}

	rsp.ID = id
	rsp.Rescode = 200
	return nil
}

func (s *OrderService) Update(ctx context.Context, req *order.UpdateRequest, rsp *order.UpdateResponse) error {
	err := dao.OrderDao.Update(db.Factoty.GetSession(), req.ID, req.MNKey, 1)
	if err != nil {
		rsp.Rescode = 500
		return nil
	}

	rsp.Rescode = 200
	return nil
}
