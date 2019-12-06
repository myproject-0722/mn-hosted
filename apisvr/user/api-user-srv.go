package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/myproject-0722/mn-hosted/apisvr/user/handler"
	"github.com/myproject-0722/mn-hosted/lib/register"
	user "github.com/myproject-0722/mn-hosted/proto/user"
)

func main() {
	service := register.NewMicroService("go.mnhosted.api.user")

	// optionally setup command line usage
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&handler.User{Client: user.NewUserService("go.mnhosted.srv.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
