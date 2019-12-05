package main

import (
	log "github.com/sirupsen/logrus"

	db "github.com/myproject-0722/mn-hosted/lib/db"
	redisclient "github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/register"
	order "github.com/myproject-0722/mn-hosted/proto/order"
	"github.com/myproject-0722/mn-hosted/rpcsvr/order/handler"
)

func main() {
	service := register.NewMicroService("go.mnhosted.srv.order")
	db.Init()
	redisclient.Init()

	// Register Handlers
	order.RegisterOrderHandler(service.Server(), handler.NewOrderService())

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
