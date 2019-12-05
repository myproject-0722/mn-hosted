package handler

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	order "github.com/myproject-0722/mn-hosted/proto/order"
	"github.com/smallnest/rpcx/log"
)

type Order struct {
	Client order.OrderService
}

func (s *Order) Alipay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Debug("Received Masternode alipay API request")

	userid, ok := req.Get["userid"]
	if !ok || len(userid.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.order", "userid cannot be blank")
	}

	coinname, ok := req.Get["coinname"]
	if !ok || len(coinname.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.order", "coinname cannot be blank")
	}

	mnkey, ok := req.Get["mnkey"]
	if !ok || len(mnkey.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.order", "mnkey cannot be blank")
	}

	timetype, ok := req.Get["timetype"]
	if !ok || len(timetype.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.order", "timetype cannot be blank")
	}

	strUserid := strings.Join(userid.Values, " ")
	intUserid, err := strconv.ParseInt(strUserid, 10, 64)
	if err != nil {
		return errors.BadRequest("go.mnhosted.api.order", "userid err")
	}

	strTimeType := strings.Join(timetype.Values, " ")
	intTimeType, err := strconv.Atoi(strTimeType)
	if err != nil {
		return errors.BadRequest("go.mnhosted.api.order", "timetype err")
	}

	resp, err := s.Client.Alipay(ctx, &order.AlipayRequest{
		UserID:   intUserid,
		CoinName: strings.Join(coinname.Values, " "),
		MNKey:    strings.Join(mnkey.Values, " "),
		TimeType: int32(intTimeType),
	})

	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.order", "order err")
	}

	rsp.StatusCode = resp.Rescode

	b, _ := json.Marshal(resp)
	rsp.Body = string(b)
	return nil
}

func (s *Order) AlipayNotify(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Debug("Received Masternode alipay notify request")

	//var noti, _ = client.GetTradeNotification(req)
	return nil
}