package main

import (
	"log"

	"github.com/myproject-0722/mn-hosted/lib/register"
	"github.com/myproject-0722/mn-hosted/notifysrv/alipay/handler"
	order "github.com/myproject-0722/mn-hosted/proto/order"
)

func main() {
	service := register.NewMicroService("go.mnhosted.api.alipaynotify")

	// optionally setup command line usage
	service.Init()

	handler.Client = order.NewOrderService("go.mnhosted.srv.order", service.Client())
	/*
		service.Server().Handle(
			service.Server().NewHandler(
				&handler.Order{
					Client: order.NewOrderService("go.mnhosted.srv.order", service.Client()),
				},
			),
		)
	*/
	handler.HttpNotifyServer()

	//handler.Test()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
