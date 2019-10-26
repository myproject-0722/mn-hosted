package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/micro/go-micro"

	db "github.com/myproject-0722/mn-hosted/lib/db"
	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	redisclient "github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/register"
	node "github.com/myproject-0722/mn-hosted/proto/node"
)

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

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
