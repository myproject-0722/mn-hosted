package main

import (
	"log"

	"github.com/myproject-0722/mn-hosted/apisvr/order/handler"
	"github.com/myproject-0722/mn-hosted/lib/register"
	order "github.com/myproject-0722/mn-hosted/proto/order"
)

func main() {
	service := register.NewMicroService("go.mnhosted.api.order")

	// optionally setup command line usage
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&handler.Order{
				Client: order.NewOrderService("go.mnhosted.srv.order", service.Client()),
			},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
