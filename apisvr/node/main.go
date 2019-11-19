package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/myproject-0722/mn-hosted/lib/db"
	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	"github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/register"
	node "github.com/myproject-0722/mn-hosted/proto/node"
	order "github.com/myproject-0722/mn-hosted/proto/order"
	user "github.com/myproject-0722/mn-hosted/proto/user"
	wallet "github.com/myproject-0722/mn-hosted/proto/wallet"

	"context"
)

/*
type Coin struct {
	Client node.CoinService
}*/

type Masternode struct {
	Client       node.MasternodeService
	CoinClient   node.CoinService
	WalletClient wallet.WalletService
	UserClient   user.UserService
	OrderClient  order.OrderService
}

func (s *Masternode) Renew(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Masternode Renew API request")

	userid, ok := req.Get["userid"]
	if !ok || len(userid.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "userid cannot be blank")
	}

	coinname, ok := req.Get["coinname"]
	if !ok || len(coinname.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "coinname cannot be blank")
	}

	mnkey, ok := req.Get["mnkey"]
	if !ok || len(mnkey.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "mnkey cannot be blank")
	}

	externalip, ok := req.Get["externalip"]
	if !ok || len(externalip.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "externalip cannot be blank")
	}

	timetype, ok := req.Get["timetype"]
	if !ok || len(timetype.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "timetype cannot be blank")
	}

	timenum, ok := req.Get["timenum"]
	if !ok || len(timenum.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "timenum cannot be blank")
	}

	strUserid := strings.Join(userid.Values, " ")
	intUserid, err := strconv.ParseInt(strUserid, 10, 64)
	if err != nil {
		//
		return errors.BadRequest("go.mnhosted.api.node", "userid cannot be error")
	}

	strTimeNum := strings.Join(timenum.Values, " ")
	intTimeNum, err := strconv.Atoi(strTimeNum)
	if err != nil {
		return errors.BadRequest("go.mnhosted.api.node", "timenum cannot be error")
	}

	coinItem, err := s.CoinClient.GetCoinItem(ctx, &node.CoinItemRequest{
		CoinName: strings.Join(coinname.Values, " "),
	})
	if err != nil || coinItem == nil {
		return errors.BadRequest("go.mnhosted.api.node", "coinname not support")
	}

	userInfo, err := s.UserClient.GetInfo(ctx, &user.GetInfoRequest{
		UserID: intUserid,
	})
	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", "account err")
	}
	account := userInfo.Account

	strTimetype := strings.Join(timetype.Values, " ")
	var price float64 = 0
	if strTimetype == "1" {
		price = coinItem.Coin.DPrice * float64(intTimeNum)
	} else if strTimetype == "2" {
		price = coinItem.Coin.MPrice * float64(intTimeNum)
	} else if strTimetype == "3" {
		price = coinItem.Coin.YPrice * float64(intTimeNum)
	}

	payResp, err := s.WalletClient.Pay(ctx, &wallet.PayRequest{
		Account: account,
		Balance: price,
	})
	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", err.Error())
	}

	if payResp.Rescode != 200 {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", "not enlough")
	}

	txid := payResp.TxID

	intTimeType, err := strconv.Atoi(strTimetype)
	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", "timetype err")
	}

	orderResp, err := s.OrderClient.New(ctx, &order.NewRequest{
		UserID:   intUserid,
		Coinname: strings.Join(coinname.Values, " "),
		Timetype: int32(intTimeType),
		Price:    int32(price),
		TxID:     txid,
		IsRenew:  1,
	})

	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", "order err")
	}

	log.Println("orderID=", orderResp.ID)

	resp, err := s.Client.Renew(ctx, &node.MasterNodeRenewRequest{
		UserId:     intUserid,
		CoinName:   strings.Join(coinname.Values, " "),
		MNKey:      strings.Join(mnkey.Values, " "),
		ExternalIp: strings.Join(externalip.Values, " "),
		TimeType:   strings.Join(timetype.Values, " "),
		TimeNum:    int32(intTimeNum),
	})

	orderUpdateResp, err := s.OrderClient.Update(ctx, &order.UpdateRequest{
		ID:    orderResp.ID,
		MNKey: strings.Join(mnkey.Values, " "),
	})
	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", "order err")
	}

	resp.Rescode = orderUpdateResp.Rescode

	b, _ := json.Marshal(resp)
	rsp.Body = string(b)
	return nil
}

func (s *Masternode) New(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Masternode New API request")

	userid, ok := req.Get["userid"]
	if !ok || len(userid.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "userid cannot be blank")
	}

	coinname, ok := req.Get["coinname"]
	if !ok || len(coinname.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "coinname cannot be blank")
	}

	mnkey, ok := req.Get["mnkey"]
	if !ok || len(mnkey.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "mnkey cannot be blank")
	}

	externalip, ok := req.Get["externalip"]
	if !ok || len(externalip.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "externalip cannot be blank")
	}

	timetype, ok := req.Get["timetype"]
	if !ok || len(timetype.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "timetype cannot be blank")
	}

	strUserid := strings.Join(userid.Values, " ")
	intUserid, err := strconv.ParseInt(strUserid, 10, 64)
	if err != nil {
		return errors.BadRequest("go.mnhosted.api.node", "userid err")
	}

	coinItem, err := s.CoinClient.GetCoinItem(ctx, &node.CoinItemRequest{
		CoinName: strings.Join(coinname.Values, " "),
	})
	if err != nil || coinItem == nil {
		return errors.BadRequest("go.mnhosted.api.node", "coinname not support")
	}

	response, err := s.Client.IsExsit(ctx, &node.MasterNodeCheckRequest{
		CoinName: strings.Join(coinname.Values, " "),
		MNKey:    strings.Join(mnkey.Values, " "),
	})

	if err == nil {
		if response.IsExsit {
			rsp.StatusCode = 404
			b, _ := json.Marshal(map[string]string{
				"message": "Masternode is exsit, please check",
			})
			rsp.Body = string(b)
			return nil
		}
	} else {
		rsp.StatusCode = 500
		return errors.BadRequest("go.mnhosted.api.node", "service err")
	}

	userInfo, err := s.UserClient.GetInfo(ctx, &user.GetInfoRequest{
		UserID: intUserid,
	})
	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", "account err")
	}
	account := userInfo.Account

	strTimetype := strings.Join(timetype.Values, " ")
	var price float64 = 0
	if strTimetype == "1" {
		price = coinItem.Coin.DPrice
	} else if strTimetype == "2" {
		price = coinItem.Coin.MPrice
	} else if strTimetype == "3" {
		price = coinItem.Coin.YPrice
	}

	payResp, err := s.WalletClient.Pay(ctx, &wallet.PayRequest{
		Account: account,
		Balance: price,
	})
	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", err.Error())
	}
	if payResp.Rescode != 200 {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", "not enlough")
	}

	intTimeType, err := strconv.Atoi(strTimetype)
	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", "timetype err")
	}

	txid := payResp.TxID
	orderResp, err := s.OrderClient.New(ctx, &order.NewRequest{
		UserID:   intUserid,
		Coinname: strings.Join(coinname.Values, " "),
		Timetype: int32(intTimeType),
		Price:    int32(price),
		TxID:     txid,
		IsRenew:  0,
	})
	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", "order err")
	}

	log.Println("orderID=", orderResp.ID)

	resp, err := s.Client.New(ctx, &node.MasterNodeNewRequest{
		UserId:     intUserid,
		CoinName:   strings.Join(coinname.Values, " "),
		MNKey:      strings.Join(mnkey.Values, " "),
		ExternalIp: strings.Join(externalip.Values, " "),
		TimeType:   strings.Join(timetype.Values, " "),
		OrderID:    orderResp.ID,
	})
	if err != nil {
		return errors.BadRequest("go.mnhosted.api.node", err.Error())
	}

	orderUpdateResp, err := s.OrderClient.Update(ctx, &order.UpdateRequest{
		ID:    orderResp.ID,
		MNKey: strings.Join(mnkey.Values, " "),
	})
	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", "order err")
	}

	resp.Rescode = orderUpdateResp.Rescode
	b, _ := json.Marshal(resp)
	rsp.Body = string(b)
	return nil
}

func (s *Masternode) Get(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Masternode Get request")

	userid, ok := req.Get["userid"]
	if !ok || len(userid.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "userid cannot be blank")
	}

	strUserid := strings.Join(userid.Values, " ")
	intUserid, err := strconv.ParseInt(strUserid, 10, 64)

	response, err := s.Client.Get(ctx, &node.MasterNodeListRequest{
		UserId: intUserid,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = response.Rescode
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	fmt.Println(rsp.Body)
	fmt.Println(response)
	return nil
}

func (s *Masternode) GetCount(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Masternode Get request")

	userid, ok := req.Get["userid"]
	if !ok || len(userid.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "userid cannot be blank")
	}

	strUserid := strings.Join(userid.Values, " ")
	intUserid, err := strconv.ParseInt(strUserid, 10, 64)

	response, err := s.Client.GetCount(ctx, &node.GetCountRequest{
		UserID: intUserid,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	fmt.Println(rsp.Body)
	fmt.Println(response)
	return nil
}

func (s *Masternode) GetCoinList(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Coinlist Get request")

	curPage, ok := req.Get["curpage"]
	if !ok || len(curPage.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "pageno cannot be blank")
	}

	strCurPage := strings.Join(curPage.Values, " ")
	intCurPage, err := strconv.Atoi(strCurPage)

	pageSize, ok := req.Get["pagesize"]
	if !ok || len(pageSize.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.node", "perpagenum cannot be blank")
	}

	strPageSize := strings.Join(pageSize.Values, " ")
	intPageSize, err := strconv.Atoi(strPageSize)
	response, err := s.CoinClient.Get(ctx, &node.CoinListRequest{
		CurPage:  int32(intCurPage),
		PageSize: int32(intPageSize),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = response.Rescode
	/*b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})*/
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	fmt.Println(rsp.Body)
	fmt.Println(response)

	return nil
}

func main() {
	liblog.InitLog("/var/log/mn-hosted/apisvr/node", "node.log")
	db.Init()
	redisclient.Init()
	reg := register.NewRegistry()

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.mnhosted.api.node"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Masternode{
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
