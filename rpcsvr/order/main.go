package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/micro/go-micro"

	"github.com/myproject-0722/mn-hosted/lib/dao"
	db "github.com/myproject-0722/mn-hosted/lib/db"
	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	"github.com/myproject-0722/mn-hosted/lib/model"
	redisclient "github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/register"
	order "github.com/myproject-0722/mn-hosted/proto/order"
)

type Order struct{}

func (s *Order) New(ctx context.Context, req *order.NewRequest, rsp *order.NewResponse) error {
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

func (s *Order) Update(ctx context.Context, req *order.UpdateRequest, rsp *order.UpdateResponse) error {
	err := dao.OrderDao.Update(db.Factoty.GetSession(), req.ID, req.MNKey, 1)
	if err != nil {
		rsp.Rescode = 500
		return nil
	}

	rsp.Rescode = 200
	return nil
}

func main() {

	liblog.InitLog("/var/log/mn-hosted/rpcsvr/order", "order.log")
	db.Init()
	redisclient.Init()
	reg := register.NewRegistry()

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.mnhosted.srv.order"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	order.RegisterOrderHandler(service.Server(), new(Order))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
