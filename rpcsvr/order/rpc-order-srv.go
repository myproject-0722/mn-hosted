package main

import (
	log "github.com/sirupsen/logrus"

	db "github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/register"
	node "github.com/myproject-0722/mn-hosted/proto/node"
	order "github.com/myproject-0722/mn-hosted/proto/order"
	"github.com/myproject-0722/mn-hosted/rpcsvr/order/handler"
)

func main() {
	service := register.NewMicroService("go.mnhosted.srv.order")
	db.Init()
	//redisclient.Init()
	service.Server().Handle(
		service.Server().NewHandler(
			&handler.OrderService{
				Client: node.NewMasternodeService("go.mnhosted.srv.node", service.Client()),
			},
		),
	)

	// Register Handlers
	order.RegisterOrderHandler(service.Server(), new(handler.OrderService))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
