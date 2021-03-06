package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/errors"
	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/http"
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

func (s *Masternode) Modify(ctx context.Context, req *node.MasterNodeModifyRequest, rsp *node.MasterNodeModifyResponse) error {
	//获取主节点信息
	node, err := dao.NodeDao.GetMasternodeByMNID(db.Factoty.GetSession(), req.MNID)
	if err != nil {
		rsp.Rescode = 500
		log.Error("GetMasternode Error", err.Error())
		return err
	}

	//修改订单信息
	var order model.Order
	order.MNName = req.MNName
	order.Id = node.OrderID
	order.MNKey = req.MNKey
	order.TxID = req.TxID
	order.TxIndex = req.TxIndex

	err = dao.NodeDao.UpdateMasternodeMNKey(db.Factoty.GetSession(), req.MNID, req.MNKey)
	if err != nil {
		log.Error("UpdateMasternodeMNKey", err.Error())
		return err
	}

	err = dao.OrderDao.Modify(db.Factoty.GetSession(), &order)
	if err != nil {
		log.Error("Modify", err.Error())
		return err
	}

	nodeid := dao.NodeDao.GetNodeIDByOrderID(db.Factoty.GetSession(), node.OrderID)
	if nodeid == -1 {
		//调用vps服务
		if http.AddVpsNode(node.OrderID) == false {
			rsp.Rescode = 500
			return errors.BadRequest("AddVpsNode", "Vps add err")
		}
	}

	//调用vps服务
	if http.UpdateVpsNode(nodeid) == false {
		rsp.Rescode = 500
		return errors.BadRequest("UpdateVpsNode", "Vps add err")
	}

	return nil
}

func (s *Masternode) ChangeNotify(ctx context.Context, req *node.MasterNodeChangeNotifyRequest, rsp *node.MasterNodeChangeNotifyResponse) error {
	//获取主节点信息
	err := dao.NodeDao.UpdateMasternodeNotify(db.Factoty.GetSession(), req.MNID, req.IsNotify)
	if err != nil {
		rsp.Rescode = 500
		log.Error("UpdateMasternodeNotify Error", err.Error())
		return err
	}

	return nil
}

func (s *Masternode) New(ctx context.Context, req *node.MasterNodeNewRequest, rsp *node.MasterNodeNewResponse) error {
	log.Debug("Received MasterNodeNewRequest userid:", req.UserId, " coinname:", req.CoinName, " key:", req.MNKey, " orderid:", req.OrderID, " isrenew:", req.IsRenew)
	//将订单纪录写入主节点表
	var masternode model.Masternode
	//判断是否为续期
	if req.IsRenew == 1 {
		dbmasternode, err := dao.NodeDao.GetMasternode(db.Factoty.GetSession(), req.CoinName, req.MNKey)
		if err != nil {
			rsp.Rescode = 500
			log.Error("GetMasternode Error", err.Error())
			return err
		}

		log.Debug("renew coiname:", req.CoinName, " mnkey:", req.MNKey, " status:", dbmasternode.Status)

		//已失效
		if dbmasternode.Status == 2 {
			masternode.ExpireTime = time.Now()

			log.Debug("add vps node orderid:", dbmasternode.OrderID)
			//重新启动
			if http.AddVpsNode(dbmasternode.OrderID) == false {
				rsp.Rescode = 500
				return errors.BadRequest("AddVpsNode", "Vps add err")
			}

		} else {
			masternode.ExpireTime = dbmasternode.ExpireTime
		}

		if req.TimeType == 1 {
			masternode.ExpireTime = masternode.ExpireTime.AddDate(0, 0, 1)
		} else if req.TimeType == 2 {
			masternode.ExpireTime = masternode.ExpireTime.AddDate(0, 1, 0)
		} else if req.TimeType == 3 {
			masternode.ExpireTime = masternode.ExpireTime.AddDate(1, 0, 0)
		}
		err = dao.NodeDao.UpdateMasternodeExpireTime(db.Factoty.GetSession(), req.CoinName, req.MNKey, masternode.ExpireTime)
		if err != nil {
			rsp.Rescode = 500
			log.Error("UpdateMasternodeExpireTime Error", err.Error())
			return err
		}
	} else {
		masternode.CoinName = req.CoinName
		masternode.MNKey = req.MNKey
		masternode.UserID = req.UserId
		masternode.OrderID = req.OrderID
		masternode.Status = 1
		masternode.CreateTime = time.Now()
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

		if http.AddVpsNode(req.OrderID) == false {
			rsp.Rescode = 500
			return errors.BadRequest("AddVpsNode", "Vps add err")
		}
	}
	rsp.Rescode = 200
	return nil
}

func (s *Masternode) Get(ctx context.Context, req *node.MasterNodeListRequest, rsp *node.MasterNodeListResponse) error {
	log.Debug("Received MasterNodeListRequest:", req.UserId)
	nodelist, err := dao.NodeDao.GetMasternodeByUserID(db.Factoty.GetSession(), req.UserId, req.CurPage, req.PageSize)
	if err != nil {
		log.Error("db GetMasternodeByUserID", err.Error())
		return err
	}

	for _, v := range nodelist {
		//log.Debug("masternodelist get", i, v.SyncStatus)
		item := new(node.MasternodeItem)
		item.MNID = v.Id
		item.CoinName = v.CoinName
		item.MNKey = v.MNKey
		item.MNPayee = v.MNPayee
		item.Vps = v.Vps
		item.Status = v.Status
		item.Earn = float64(v.Earn) / 1000000
		item.SyncStatus = v.SyncStatus
		item.SyncStatusEx = v.SyncStatusEx
		item.MNStatus = v.MNStatusEx
		item.IsNotify = v.IsNotify

		item.CreateTime = v.CreateTime.Local().Format("2006-01-02 15:04:05")
		item.ExpireTime = v.ExpireTime.Local().Format("2006-01-02 15:04:05")
		rsp.Masternodelist = append(rsp.Masternodelist, item)
	}
	return nil
}
