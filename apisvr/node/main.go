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

func (s *Coinlist) Get(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received SignIn API request")

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

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
