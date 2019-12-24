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

func (s *Order) GetOrderList(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Debug("Received Order GetOrderList API request")

	userid, ok := req.Get["userid"]
	if !ok || len(userid.Values) == 0 {
		log.Error("userid cannot be blank")
		return errors.BadRequest("go.mnhosted.api.order", "userid cannot be blank")
	}

	strUserid := strings.Join(userid.Values, " ")
	intUserid, err := strconv.ParseInt(strUserid, 10, 64)
	if err != nil {
		log.Error("userid err")
		return errors.BadRequest("go.mnhosted.api.order", "userid err")
	}

	resp, err := s.Client.GetOrderList(ctx, &order.GetOrderListRequest{
		UserID: intUserid,
	})

	if err != nil {
		rsp.StatusCode = 404
		log.Error("Alipay error", err.Error())
		return errors.BadRequest("go.mnhosted.api.order", "userid err")
	}

	rsp.StatusCode = resp.Rescode

	b, _ := json.Marshal(resp)
	rsp.Body = string(b)
	return nil
}

func (s *Order) GetOrderInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Debug("Received Order GetInfo API request")

	userid, ok := req.Get["userid"]
	if !ok || len(userid.Values) == 0 {
		log.Error("userid cannot be blank")
		return errors.BadRequest("go.mnhosted.api.order", "userid cannot be blank")
	}

	strUserid := strings.Join(userid.Values, " ")
	intUserid, err := strconv.ParseInt(strUserid, 10, 64)
	if err != nil {
		log.Error("userid err")
		return errors.BadRequest("go.mnhosted.api.order", "userid err")
	}

	resp, err := s.Client.GetInfo(ctx, &order.GetInfoRequest{
		UserID: intUserid,
	})

	if err != nil {
		rsp.StatusCode = 404
		log.Error("Alipay error", err.Error())
		return errors.BadRequest("go.mnhosted.api.order", "userid err")
	}

	rsp.StatusCode = resp.Rescode

	b, _ := json.Marshal(resp)
	rsp.Body = string(b)
	return nil
}

func (s *Order) Alipay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Debug("Received Masternode alipay API request")

	userid, ok := req.Get["userid"]
	if !ok || len(userid.Values) == 0 {
		log.Error("userid cannot be blank")
		return errors.BadRequest("go.mnhosted.api.order", "userid cannot be blank")
	}

	coinname, ok := req.Get["coinname"]
	if !ok || len(coinname.Values) == 0 {
		log.Error("coinname cannot be blank")
		return errors.BadRequest("go.mnhosted.api.order", "coinname cannot be blank")
	}

	txid, ok := req.Get["txid"]
	strTxID := ""
	if ok && len(txid.Values) > 0 {
		strTxID = strings.Join(txid.Values, " ")
	}

	mnname, ok := req.Get["mnname"]
	strMNName := ""
	if ok && len(mnname.Values) > 0 {
		strMNName = strings.Join(mnname.Values, " ")
	}

	txindex, ok := req.Get["txindex"]
	intTxIndex := 0
	if ok && len(txindex.Values) > 0 {
		intTxIndex, _ = strconv.Atoi(strings.Join(txindex.Values, " "))
	}

	mnkey, ok := req.Get["mnkey"]
	if !ok || len(mnkey.Values) == 0 {
		log.Error("mnkey cannot be blank")
		return errors.BadRequest("go.mnhosted.api.order", "mnkey cannot be blank")
	}

	timetype, ok := req.Get["timetype"]
	if !ok || len(timetype.Values) == 0 {
		log.Error("timetype cannot be blank")
		return errors.BadRequest("go.mnhosted.api.order", "timetype cannot be blank")
	}

	var intIsRenew int = 0
	isRenew, ok := req.Get["isrenew"]
	if ok && len(isRenew.Values) != 0 {
		intIsRenew, _ = strconv.Atoi(strings.Join(isRenew.Values, " "))
	}

	strUserid := strings.Join(userid.Values, " ")
	intUserid, err := strconv.ParseInt(strUserid, 10, 64)
	if err != nil {
		log.Error("userid err")
		return errors.BadRequest("go.mnhosted.api.order", "userid err")
	}

	strTimeType := strings.Join(timetype.Values, " ")
	intTimeType, err := strconv.Atoi(strTimeType)
	if err != nil {
		log.Error("timetype err")
		return errors.BadRequest("go.mnhosted.api.order", "timetype err")
	}

	resp, err := s.Client.Alipay(ctx, &order.AlipayRequest{
		UserID:   intUserid,
		CoinName: strings.Join(coinname.Values, " "),
		MNKey:    strings.Join(mnkey.Values, " "),
		MNName:   strMNName,
		TxID:     strTxID,
		TxIndex:  int32(intTxIndex),
		TimeType: int32(intTimeType),
		IsRenew:  int32(intIsRenew),
	})

	if err != nil {
		rsp.StatusCode = 404
		log.Error("Alipay error", err.Error())
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
