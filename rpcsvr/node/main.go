package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/micro/go-micro"

	"github.com/myproject-0722/mn-hosted/lib/dao"
	db "github.com/myproject-0722/mn-hosted/lib/db"
	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	redisclient "github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/register"
	node "github.com/myproject-0722/mn-hosted/proto/node"
)

type Coinlist struct{}

func (s *Coinlist) Get(ctx context.Context, req *node.CoinListRequest, rsp *node.CoinListResponse) error {
	coin, err := dao.NodeDao.GetCoinList(db.Factoty.GetSession(), req.CurPage, req.PageSize)
	if err != nil {
		return err
	}

	rsp.Rescode = 200
	item := new(node.CoinItem)
	item.CoinName = coin.CoinName
	item.MNRequired = coin.MNRequired
	item.MNPrice = coin.MNPrice
	item.Volume = coin.Volume
	item.Roi = coin.Roi
	item.MonthlyIncome = coin.MonthlyIncome
	item.MNHosted = coin.MNHosted
	rsp.Coinlist = append(rsp.Coinlist, item)
	return nil
}

type Masternode struct{}

func (s *Masternode) New(ctx context.Context, req *node.MasterNodeNewRequest, rsp *node.MasterNodeNewResponse) error {
	log.Print("Received MasterNodeNewRequest:", req.UserId)
	/*log.Print("Received SignUpRequest Name: ", req.Account, " Passwd: ", req.Passwd)

	id, err := dao.UserDao.Add(db.Factoty.GetSession(), req.Account, req.Passwd)
	if err != nil {
		rsp.Rescode = 404
		rsp.Msg = " SignUp Error"
		return nil
	}
	rsp.Rescode = 200
	rsp.Msg = " SignUp OK!"
	rsp.Id = id*/
	return nil
}

func main() {

	liblog.InitLog("/var/log/mn-hosted/rpcsvr/node", "node.log")
	db.Init()
	redisclient.Init()
	reg := register.NewRegistry()

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.mnhosted.srv.node"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	node.RegisterMasternodeHandler(service.Server(), new(Masternode))
	node.RegisterCoinlistHandler(service.Server(), new(Coinlist))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
