package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"encoding/json"

	"github.com/myproject-0722/mn-hosted/lib/cmd"
	"github.com/myproject-0722/mn-hosted/lib/dao"
	db "github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/model"
	redisclient "github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/register"
	node "github.com/myproject-0722/mn-hosted/proto/node"
)

func GetNodeSyncStatus(userid int64, mnkey string, dockerid string) bool {
	var statusCmd string = "docker exec " + dockerid + " /bin/bash -c 'dash-cli -testnet mnsync status'"
	log.Print("statuscmd=", statusCmd)
	syncStatus := cmd.ExecShell(statusCmd)
	log.Print("mnstatus=", syncStatus)

	if syncStatus != "" {
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(syncStatus), &data); err == nil {
			log.Println("==============json str 转map=======================")
			log.Println(data)
			log.Println(data["IsBlockchainSynced"])
			if data["IsBlockchainSynced"] == true {
				return true
			}
		}
	}

	return false
}

type Coin struct{}

func (s *Coin) Get(ctx context.Context, req *node.CoinListRequest, rsp *node.CoinListResponse) error {
	coinlist, err := dao.NodeDao.GetCoinList(db.Factoty.GetSession(), req.CurPage, req.PageSize)
	if err != nil {
		return err
	}

	rsp.Rescode = 200
	coinlistLength := len(coinlist)
	for i := 0; i < coinlistLength; i++ {
		v := coinlist[i]
		item := new(node.CoinItem)
		item.CoinName = v.CoinName
		item.MNRequired = v.MNRequired
		item.DPrice = float64(v.DPrice / 100)
		item.MPrice = float64(v.MPrice / 100)
		item.YPrice = float64(v.YPrice / 100)
		item.Volume = v.Volume
		item.Roi = v.Roi
		item.MonthlyIncome = v.MonthlyIncome
		item.MNHosted = v.MNHosted
		rsp.Coinlist = append(rsp.Coinlist, item)
	}

	return nil
}

func (s *Coin) GetCoinItem(ctx context.Context, req *node.CoinItemRequest, rsp *node.CoinItemResponse) error {
	v, err := dao.NodeDao.GetCoinItem(db.Factoty.GetSession(), req.CoinName)
	if err != nil {
		return err
	}

	rsp.Rescode = 200
	rsp.Coin = new(node.CoinItem)
	rsp.Coin.CoinName = v.CoinName
	rsp.Coin.MNRequired = v.MNRequired
	rsp.Coin.DPrice = float64(v.DPrice / 100)
	rsp.Coin.MPrice = float64(v.MPrice / 100)
	rsp.Coin.YPrice = float64(v.YPrice / 100)
	rsp.Coin.Volume = v.Volume
	rsp.Coin.Roi = v.Roi
	rsp.Coin.MonthlyIncome = v.MonthlyIncome
	rsp.Coin.MNHosted = v.MNHosted

	return nil
}

type Masternode struct{}

func (s *Masternode) IsExsit(ctx context.Context, req *node.MasterNodeCheckRequest, rsp *node.MasterNodeCheckResponse) error {
	masternode, err := dao.NodeDao.GetMasternode(db.Factoty.GetSession(), req.CoinName, req.MNKey)
	if err != nil {
		rsp.IsExsit = false
		return err
	}

	if masternode != nil {
		rsp.IsExsit = true
	} else {
		rsp.IsExsit = false
	}

	return nil
}

func (s *Masternode) GetCount(ctx context.Context, req *node.GetCountRequest, rsp *node.GetCountResponse) error {
	count := dao.NodeDao.GetMasternodeCount(db.Factoty.GetSession(), req.UserID)
	rsp.Count = count

	return nil
}

func (s *Masternode) Renew(ctx context.Context, req *node.MasterNodeRenewRequest, rsp *node.MasterNodeRenewResponse) error {
	masternode, err := dao.NodeDao.GetMasternode(db.Factoty.GetSession(), req.CoinName, req.MNKey)
	if err != nil {
		return err
	}

	if masternode != nil && masternode.ExpireTime.After(time.Now()) {
		//更新有效期时间即可
		if req.TimeType == "1" {
			masternode.ExpireTime = masternode.ExpireTime.AddDate(0, 0, int(req.TimeNum))
		} else if req.TimeType == "2" {
			masternode.ExpireTime = masternode.ExpireTime.AddDate(0, int(req.TimeNum), 0)
		} else if req.TimeType == "3" {
			masternode.ExpireTime = masternode.ExpireTime.AddDate(int(req.TimeNum), 0, 0)
		} else {
			//error
			rsp.Rescode = 401
		}

		dao.NodeDao.UpdateMasternodeExpireTime(db.Factoty.GetSession(), req.CoinName, req.MNKey, masternode.ExpireTime)
	} else {
		rsp.Rescode = 404
	}
	return nil
}

func (s *Masternode) New(ctx context.Context, req *node.MasterNodeNewRequest, rsp *node.MasterNodeNewResponse) error {
	log.Debug("Received MasterNodeNewRequest:", req.UserId)

	//将订单纪录写入主节点表
	var masternode model.Masternode
	masternode.CoinName = req.CoinName
	masternode.MNKey = req.MNKey
	masternode.UserID = req.UserId
	masternode.OrderID = req.OrderID
	masternode.Status = 1
	masternode.ExpireTime = time.Now()
	if req.TimeType == 1 {
		masternode.ExpireTime = masternode.ExpireTime.AddDate(0, 0, 1)
	} else if req.TimeType == 2 {
		masternode.ExpireTime = masternode.ExpireTime.AddDate(0, 1, 0)
	} else if req.TimeType == 3 {
		masternode.ExpireTime = masternode.ExpireTime.AddDate(1, 0, 0)
	}

	id, err := dao.NodeDao.AddMasternode(db.Factoty.GetSession(), masternode)
	if err != nil {
		rsp.Rescode = 500
		log.Error("Sql Error, please check!", err.Error())
		return err
	}
	log.Debug("AddMasternode id=", id)
	/* 暂时注掉方便测试
	if http.AddVpsNode(req.OrderID) == false {
		rsp.Rescode = 500
		return errors.BadRequest("AddVpsNode", "Vps add err")
	}
	*/
	rsp.Rescode = 200
	return nil
}

func (s *Masternode) Get(ctx context.Context, req *node.MasterNodeListRequest, rsp *node.MasterNodeListResponse) error {
	log.Debug("Received MasterNodeListRequest:", req.UserId)
	nodelist, err := dao.NodeDao.GetMasternodeByUserID(db.Factoty.GetSession(), req.UserId)
	if err != nil {
		return err
	}

	for i, v := range nodelist {
		log.Debug("masternodelist get", i, v.SyncStatus)
		item := new(node.MasternodeItem)
		item.MNID = v.Id
		item.CoinName = v.CoinName
		item.MNKey = v.MNKey
		item.SyncStatus = v.SyncStatus
		item.CreateTime = v.CreateTime.Format("2006-01-02 15:04:05")
		item.ExpireTime = v.ExpireTime.Local().Format("2006-01-02 15:04:05")
		rsp.Masternodelist = append(rsp.Masternodelist, item)
	}
	return nil
}

func main() {
	service := register.NewMicroService("go.mnhosted.srv.node")
	db.Init()
	redisclient.Init()

	// Register Handlers
	node.RegisterMasternodeHandler(service.Server(), new(Masternode))
	node.RegisterCoinHandler(service.Server(), new(Coin))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
