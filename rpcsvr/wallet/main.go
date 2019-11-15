package main

import (
	"context"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/micro/go-micro"

	"github.com/myproject-0722/mn-hosted/conf"
	db "github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/http"
	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	redisclient "github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/register"
	wallet "github.com/myproject-0722/mn-hosted/proto/wallet"
)

type Wallet struct{}

func (s *Wallet) New(ctx context.Context, req *wallet.NewRequest, rsp *wallet.NewResponse) error {
	log.Print("Received NewRequest Account: ", req.Account)

	cmd := "getnewaddress " + req.Account
	res, err := http.GetRpcCallResult(cmd)
	if err != nil || res == nil {
		return err
	}

	rsp.Address = res.(string)
	return nil
}

func (s *Wallet) GetBalance(ctx context.Context, req *wallet.GetBalanceRequest, rsp *wallet.GetBalanceResponse) error {
	log.Print("Received GetBalanceRequest Account: ", req.Account)

	cmd := "getbalance " + req.Account
	res, err := http.GetRpcCallResult(cmd)
	if err != nil || res == nil {
		return err
	}

	rsp.Balance = res.(float64)
	return nil
}

func (s *Wallet) Pay(ctx context.Context, req *wallet.PayRequest, rsp *wallet.PayResponse) error {
	log.Print("Received payRequest Account: ", req.Account)

	cmd := "sendfrom " + req.Account + " " + conf.MNHostedAddress + " " + strconv.FormatFloat(req.Balance, 'f', 6, 64)
	res, err := http.GetRpcCallResult(cmd)
	if err != nil || res == nil {
		rsp.Rescode = 404
		return err
	}

	if res == nil {
		rsp.Rescode = 404
		return nil
	}

	rsp.Rescode = 200
	rsp.TxID = res.(string)
	log.Println("txid=", rsp.TxID)
	//rsp.re = res.(float64)
	return nil
}

func main() {

	liblog.InitLog("/var/log/mn-hosted/rpcsvr/wallet", "wallet.log")
	db.Init()
	redisclient.Init()
	reg := register.NewRegistry()

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.mnhosted.srv.wallet"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	wallet.RegisterWalletHandler(service.Server(), new(Wallet))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
