package handler

import (
	//"strconv"

	"time"

	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/http"
	log "github.com/sirupsen/logrus"
)

func CheckMasterNodeExpired() {
	log.Info("CheckMasterNodeExpired")
	nodelist, err := dao.NodeDao.GetExpiredTimeMasternode(db.Factoty.GetSession(), time.Now())
	if err != nil {
		log.Error("GetExpiredTimeMasternode", err)
	}

	for i, v := range nodelist {
		nodeid := dao.NodeDao.GetNodeIDByOrderID(db.Factoty.GetSession(), v.OrderID)
		if nodeid == -1 {
			continue
		}

		if http.DelVpsNode(nodeid) == false {
			continue
		}

		//先备份数据到备份表
		err := dao.NodeDao.BackupMasternode(db.Factoty.GetSession(), *v)
		if err != nil {
			log.Error("BackupMasternode orderid=", v.OrderID, i, err)
			continue
		}
		//删除masternode表纪录
		delerr := dao.NodeDao.DelMasternodeByID(db.Factoty.GetSession(), v.Id)
		if delerr != nil {
			log.Fatal("DelMasternodeByID orderid=", v.OrderID, i, delerr)
		}
		//删除节点
		/*var vpsDelCmd string = "vpsdel " + strconv.FormatInt(v.OrderID, 10)
		log.Print("vpsDelCmd=", vpsDelCmd)
		vpsDelCmdStatus := cmd.ExecShell(vpsDelCmd)
		if vpsDelCmdStatus == "1" {
			log.Error("vps add err")
		}*/
		/*item := new(node.CoinItem)
		item.CoinName = v.CoinName
		item.MNRequired = v.MNRequired
		item.DPrice = float64(v.DPrice / 100)
		item.MPrice = float64(v.MPrice / 100)
		item.YPrice = float64(v.YPrice / 100)
		item.Volume = v.Volume
		item.Roi = v.Roi
		item.MonthlyIncome = v.MonthlyIncome
		item.MNHosted = v.MNHosted
		rsp.Coinlist = append(rsp.Coinlist, item)*/
	}
}
