package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/micro/go-micro"

	"encoding/json"

	"github.com/myproject-0722/mn-hosted/lib/cmd"
	"github.com/myproject-0722/mn-hosted/lib/dao"
	db "github.com/myproject-0722/mn-hosted/lib/db"
	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	"github.com/myproject-0722/mn-hosted/lib/model"
	redisclient "github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/register"
	node "github.com/myproject-0722/mn-hosted/proto/node"
	"github.com/robfig/cron"
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

type Coinlist struct{}

func (s *Coinlist) Get(ctx context.Context, req *node.CoinListRequest, rsp *node.CoinListResponse) error {
	coinlist, err := dao.NodeDao.GetCoinList(db.Factoty.GetSession(), req.CurPage, req.PageSize)
	if err != nil {
		return err
	}

	rsp.Rescode = 200
	for i, v := range coinlist {
		log.Print("list get", i)
		item := new(node.CoinItem)
		item.CoinName = v.CoinName
		item.MNRequired = v.MNRequired
		item.MNPrice = v.MNPrice
		item.Volume = v.Volume
		item.Roi = v.Roi
		item.MonthlyIncome = v.MonthlyIncome
		item.MNHosted = v.MNHosted
		rsp.Coinlist = append(rsp.Coinlist, item)
	}

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
		//新建
		var command string = "docker run -d -e mnkey=" + req.MNKey + " -e externalip=" + req.ExternalIp + " mnhosted/dashcore:v1.0"
		fmt.Println(command)
		dockerid := cmd.ExecShell(command)
		dockerid = strings.TrimRight(dockerid, "\n")

		var masternode model.Masternode
		masternode.CoinName = req.CoinName
		masternode.MNKey = req.MNKey
		masternode.UserID = req.UserId
		masternode.Vps = req.ExternalIp
		masternode.Status = 1
		masternode.ExpireTime = time.Now()

		masternode.DockerID = dockerid
		if req.TimeType == "1" {
			masternode.ExpireTime = masternode.ExpireTime.AddDate(0, 0, int(req.TimeNum))
		} else if req.TimeType == "2" {
			masternode.ExpireTime = masternode.ExpireTime.AddDate(0, int(req.TimeNum), 0)
		} else if req.TimeType == "3" {
			masternode.ExpireTime = masternode.ExpireTime.AddDate(int(req.TimeNum), 0, 0)
		}
	}
	return nil
}

func (s *Masternode) New(ctx context.Context, req *node.MasterNodeNewRequest, rsp *node.MasterNodeNewResponse) error {
	log.Print("Received MasterNodeNewRequest:", req.UserId)
	//var command string = "docker run -d -p 19999:9999 --name mnhosted-dashcore -e mnkey=" + req.MNKey + " -e externalip=" + req.ExternalIp + " mnhosted/dashcore:v1.0"
	var command string = "docker run -d -e mnkey=" + req.MNKey + " -e externalip=" + req.ExternalIp + " mnhosted/dashcore:v1.0"
	fmt.Println(command)
	dockerid := cmd.ExecShell(command)
	dockerid = strings.TrimRight(dockerid, "\n")

	var masternode model.Masternode
	masternode.CoinName = req.CoinName
	masternode.MNKey = req.MNKey
	masternode.UserID = req.UserId
	masternode.Vps = req.ExternalIp
	masternode.Status = 1
	masternode.ExpireTime = time.Now()
	/* local, err := time.LoadLocation("Asia/Chongqing") //服务器设置的时区
	if err != nil {
		fmt.Println(err)
	}
	masternode.ExpireTime.In(local) */

	masternode.DockerID = dockerid
	if req.TimeType == "1" {
		masternode.ExpireTime = masternode.ExpireTime.AddDate(0, 0, 1)
	} else if req.TimeType == "2" {
		masternode.ExpireTime = masternode.ExpireTime.AddDate(0, 1, 0)
	} else if req.TimeType == "3" {
		masternode.ExpireTime = masternode.ExpireTime.AddDate(1, 0, 0)
	}

	id, err := dao.NodeDao.AddMasternode(db.Factoty.GetSession(), masternode)
	if err != nil {
		log.Error("sql error, please check!")
		return err
	}
	log.Print("mn id=", id, " dockerid=", dockerid)

	c := cron.New()
	spec := "*/30 * * * * ?"
	c.AddFunc(spec, func() {
		ret := GetNodeSyncStatus(req.UserId, req.MNKey, dockerid)
		if ret == true {
			c.Stop()
			//更新状态
			dao.NodeDao.UpdateMasternodeSyncStatus(db.Factoty.GetSession(), "dashcore", req.MNKey, 1)
		}
	})
	c.Start()
	rsp.Rescode = 200
	//defer c.Stop()
	//GetNodeSyncStatus(req.UserId, req.MNKey, dockerid)
	/* var syncCmd string = "docker exec -it " + dockerid + "/bin/bash -c dash-cli -testnet mnsync status"
	syncStatus := cmd.ExecShell(syncCmd)
	log.Print("mnstatus=", syncStatus) */
	return nil
}

func (s *Masternode) Get(ctx context.Context, req *node.MasterNodeListRequest, rsp *node.MasterNodeListResponse) error {
	log.Print("Received MasterNodeListRequest:", req.UserId)
	nodelist, err := dao.NodeDao.GetMasternodeByUserID(db.Factoty.GetSession(), req.UserId)
	if err != nil {
		return err
	}

	for i, v := range nodelist {
		log.Print("list get", i, v.SyncStatus)
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

	liblog.InitLog("/var/log/mn-hosted/rpcsvr/node", "node.log")
	db.Init()
	redisclient.Init()
	reg := register.NewRegistry()

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.mnhosted.srv.node"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	node.RegisterMasternodeHandler(service.Server(), new(Masternode))
	node.RegisterCoinlistHandler(service.Server(), new(Coinlist))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
