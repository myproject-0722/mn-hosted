package handler

import (
	"context"
	"time"

	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/model"
	node "github.com/myproject-0722/mn-hosted/proto/node"
	log "github.com/sirupsen/logrus"
)

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
	log.Debug("Received MasterNodeNewRequest:", req.UserId, " ", req.CoinName, " ", req.MNKey, " ", req.OrderID)
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
