package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/myproject-0722/mn-hosted/apisvr/node/handler"
	"github.com/myproject-0722/mn-hosted/lib/register"
	node "github.com/myproject-0722/mn-hosted/proto/node"
	order "github.com/myproject-0722/mn-hosted/proto/order"
	user "github.com/myproject-0722/mn-hosted/proto/user"
	wallet "github.com/myproject-0722/mn-hosted/proto/wallet"
)

func main() {
	service := register.NewMicroService("go.mnhosted.api.node")

	// optionally setup command line usage
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&handler.Masternode{
				Client:       node.NewMasternodeService("go.mnhosted.srv.node", service.Client()),
				CoinClient:   node.NewCoinService("go.mnhosted.srv.node", service.Client()),
				UserClient:   user.NewUserService("go.mnhosted.srv.user", service.Client()),
				OrderClient:  order.NewOrderService("go.mnhosted.srv.order", service.Client()),
				WalletClient: wallet.NewWalletService("go.mnhosted.srv.wallet", service.Client()),
			},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
