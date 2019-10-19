package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/micro/go-micro"

	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	"github.com/myproject-0722/mn-hosted/lib/register"
	user "github.com/myproject-0722/mn-hosted/proto/user"
)

type User struct{}

func (s *User) SignUp(ctx context.Context, req *user.SignUpRequest, rsp *user.SignUpResponse) error {
	log.Print("Received SignUpRequest Name: ", req.Name, " Passwd: ", req.Passwd)
	rsp.Rescode = 200
	rsp.Msg = " SignUp OK!"
	return nil
}

func (s *User) SignIn(ctx context.Context, req *user.SignInRequest, rsp *user.SignInResponse) error {
	log.Print("Received SignInRequest Name: ", req.Name, " Passwd: ", req.Passwd)
	rsp.Rescode = 200
	rsp.Msg = " SignIn OK!"
	return nil
}

func main() {

	liblog.InitLog("/var/log/mn-hosted/rpcsvr/user", "user.log")

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
