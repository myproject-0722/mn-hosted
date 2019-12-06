package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/myproject-0722/mn-hosted/conf"
	db "github.com/myproject-0722/mn-hosted/lib/db"
	redisclient "github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/register"
	"github.com/myproject-0722/mn-hosted/lib/token"
	user "github.com/myproject-0722/mn-hosted/proto/user"
	wallet "github.com/myproject-0722/mn-hosted/proto/wallet"
	"github.com/myproject-0722/mn-hosted/rpcsvr/user/handler"
)

func main() {
	service := register.NewMicroService("go.mnhosted.srv.user")

	token := &token.Token{}
	token.InitConfig(conf.GetConsulHosts(), "micro", "config", "jwt-key", "key")
	db.Init()
	redisclient.Init()

	handler.Client = wallet.NewWalletService("go.mnhosted.srv.wallet", service.Client())
	handler.Token = token
	// Register Handlers
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
