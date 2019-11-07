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

	"context"
)

type Coinlist struct {
	Client node.CoinlistService
}

type Masternode struct {
	Client node.MasternodeService
}

func (s *Masternode) New(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Masternode New API request")

	userid, ok := req.Get["userid"]
	if !ok || len(userid.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "userid cannot be blank")
	}

	token, ok := req.Get["token"]
	if !ok || len(token.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.node", "token cannot be blank")
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

	response, err := s.Client.IsExsit(ctx, &node.MasterNodeCheckRequest{
		CoinName: strings.Join(coinname.Values, " "),
		MNKey:    strings.Join(mnkey.Values, " "),
	})

	if err == nil {
		if response.IsExsit {
			rsp.StatusCode = 404
			b, _ := json.Marshal(map[string]string{
				"message": "Masternode is exsit",
			})
			rsp.Body = string(b)
			return nil
		}
	} else {
		rsp.StatusCode = 500
		return errors.BadRequest("go.mnhosted.api.node", "system err")
	}

	strUserid := strings.Join(userid.Values, " ")
	intUserid, err := strconv.ParseInt(strUserid, 10, 64)

	resp, err := s.Client.New(ctx, &node.MasterNodeNewRequest{
		UserId:     intUserid,
		CoinName:   strings.Join(coinname.Values, " "),
		MNKey:      strings.Join(mnkey.Values, " "),
		ExternalIp: strings.Join(externalip.Values, " "),
		TimeType:   strings.Join(timetype.Values, " "),
	})

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

func (s *Coinlist) Get(ctx context.Context, req *api.Request, rsp *api.Response) error {
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
	response, err := s.Client.Get(ctx, &node.CoinListRequest{
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
			&Coinlist{Client: node.NewCoinlistService("go.mnhosted.srv.node", service.Client())},
		),
	)

	service.Server().Handle(
		service.Server().NewHandler(
			&Masternode{Client: node.NewMasternodeService("go.mnhosted.srv.node", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
