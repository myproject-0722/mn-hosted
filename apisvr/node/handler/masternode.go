package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/micro/go-micro/errors"
	log "github.com/sirupsen/logrus"

	api "github.com/micro/go-micro/api/proto"
	node "github.com/myproject-0722/mn-hosted/proto/node"
	order "github.com/myproject-0722/mn-hosted/proto/order"
	user "github.com/myproject-0722/mn-hosted/proto/user"
	wallet "github.com/myproject-0722/mn-hosted/proto/wallet"
)

type Masternode struct {
	Client       node.MasternodeService
	CoinClient   node.CoinService
	WalletClient wallet.WalletService
	UserClient   user.UserService
	OrderClient  order.OrderService
}

//只供使用数字货币续期使用
func (s *Masternode) Renew(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Debug("Received Masternode Renew API request")

	//参数解析
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

	//判断coinname是否能处理
	coinItem, err := s.CoinClient.GetCoinItem(ctx, &node.CoinItemRequest{
		CoinName: strings.Join(coinname.Values, " "),
	})
	if err != nil || coinItem == nil {
		return errors.BadRequest("go.mnhosted.api.node", "coinname not support")
	}

	//获取账号信息
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

	//数字货币支付
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

	//纪录订单
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

	log.Debug("orderID=", orderResp.ID)

	//续期
	resp, err := s.Client.Renew(ctx, &node.MasterNodeRenewRequest{
		UserId:   intUserid,
		CoinName: strings.Join(coinname.Values, " "),
		MNKey:    strings.Join(mnkey.Values, " "),
		TimeType: strings.Join(timetype.Values, " "),
		TimeNum:  int32(intTimeNum),
	})

	//更改订单状态
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

//供数字货币使用
func (s *Masternode) New(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Debug("Received Masternode New API request")

	//解析参数
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
	timetype, ok := req.Get["timetype"]
	if !ok || len(timetype.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "timetype cannot be blank")
	}
	strUserid := strings.Join(userid.Values, " ")
	intUserid, err := strconv.ParseInt(strUserid, 10, 64)
	if err != nil {
		return errors.BadRequest("go.mnhosted.api.node", "userid err")
	}

	//判断coinname是否可处理
	coinItem, err := s.CoinClient.GetCoinItem(ctx, &node.CoinItemRequest{
		CoinName: strings.Join(coinname.Values, " "),
	})
	if err != nil || coinItem == nil {
		return errors.BadRequest("go.mnhosted.api.node", "coinname not support")
	}

	//判断主节点是否存在
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

	//获取账号信息
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
	intTimeType, err := strconv.Atoi(strTimetype)
	if err != nil {
		rsp.StatusCode = 404
		return errors.BadRequest("go.mnhosted.api.node", "timetype err")
	}
	if intTimeType == 1 {
		price = coinItem.Coin.DPrice
	} else if intTimeType == 2 {
		price = coinItem.Coin.MPrice
	} else if intTimeType == 3 {
		price = coinItem.Coin.YPrice
	}

	//扣款
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

	//记录订单
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

	//创建主节点
	resp, err := s.Client.New(ctx, &node.MasterNodeNewRequest{
		UserId:   intUserid,
		CoinName: strings.Join(coinname.Values, " "),
		MNKey:    strings.Join(mnkey.Values, " "),
		TimeType: int32(intTimeType),
		OrderID:  orderResp.ID,
	})
	if err != nil {
		return errors.BadRequest("go.mnhosted.api.node", err.Error())
	}

	//修改订单状态为完成
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
	log.Debug("Received Masternode Get request")

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
	log.Debug("Received Masternode Get request")

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

func (s *Masternode) GetCoinRewards(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Debug("Received Masternode GetCoinRewards request")

	userid, ok := req.Get["userid"]
	if !ok || len(userid.Values) == 0 {
		return errors.BadRequest("go.mnhosted.api.node", "userid cannot be blank")
	}

	strUserid := strings.Join(userid.Values, " ")
	intUserid, err := strconv.ParseInt(strUserid, 10, 64)

	response, err := s.CoinClient.GetCoinRewards(ctx, &node.CoinRewardsRequest{
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
	log.Debug("Received Coinlist Get request")

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
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	fmt.Println(rsp.Body)
	fmt.Println(response)

	return nil
}

func (s *Masternode) Modify(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Debug("Received Modify request")

	mnid, ok := req.Get["mnid"]
	if !ok || len(mnid.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.node", "mnid cannot be blank")
	}
	strMNID := strings.Join(mnid.Values, " ")
	intMNID, err := strconv.ParseInt(strMNID, 10, 64)

	mnname, ok := req.Get["mnname"]
	strMNName := ""
	if ok && len(mnname.Values) > 0 {
		strMNName = strings.Join(mnname.Values, " ")
	}

	mnkey, ok := req.Get["mnkey"]
	if !ok || len(mnkey.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.node", "mnkey cannot be blank")
	}

	txid, ok := req.Get["txid"]
	strTxID := ""
	if ok && len(txid.Values) > 0 {
		strTxID = strings.Join(txid.Values, " ")
	}

	txindex, ok := req.Get["txindex"]
	intTxIndex := 0
	if ok && len(txindex.Values) > 0 {
		intTxIndex, _ = strconv.Atoi(strings.Join(txindex.Values, " "))
	}

	response, err := s.Client.Modify(ctx, &node.MasterNodeModifyRequest{
		MNID:    intMNID,
		MNName:  strMNName,
		MNKey:   strings.Join(mnkey.Values, " "),
		TxID:    strTxID,
		TxIndex: int32(intTxIndex),
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
