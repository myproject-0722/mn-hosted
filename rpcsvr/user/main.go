package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/micro/go-micro"

	"github.com/myproject-0722/mn-hosted/lib/dao"
	db "github.com/myproject-0722/mn-hosted/lib/db"
	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	"github.com/myproject-0722/mn-hosted/lib/register"
	user "github.com/myproject-0722/mn-hosted/proto/user"
)

type User struct{}

func (s *User) SignUp(ctx context.Context, req *user.SignUpRequest, rsp *user.SignUpResponse) error {
	log.Print("Received SignUpRequest Name: ", req.Account, " Passwd: ", req.Passwd)

	id, err := dao.UserDao.Add(db.Factoty.GetSession(), req.Account, req.Passwd)
	if err != nil {
		rsp.Rescode = 404
		rsp.Msg = " SignUp Error"
		return nil
	}
	rsp.Rescode = 200
	rsp.Msg = " SignUp OK!"
	rsp.Id = id
	return nil
}

func (s *User) SignIn(ctx context.Context, req *user.SignInRequest, rsp *user.SignInResponse) error {
	log.Print("Received SignInRequest Name: ", req.Account, " Passwd: ", req.Passwd)
	isEquis, err := dao.UserDao.Check(db.Factoty.GetSession(), req.Account, req.Passwd)
	if isEquis != true || err != nil {
		rsp.Rescode = 404
		rsp.Msg = " SignIn Error!"
		return nil
	}
	rsp.Rescode = 200
	rsp.Msg = " SignIn OK!"
	return nil
}

func main() {

	liblog.InitLog("/var/log/mn-hosted/rpcsvr/user", "user.log")
	db.Init()
	reg := register.NewRegistry()

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.mnhosted.srv.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	user.RegisterUserHandler(service.Server(), new(User))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
