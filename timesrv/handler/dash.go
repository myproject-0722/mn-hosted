package handler

import (
	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/http"
	log "github.com/sirupsen/logrus"
)

func SyncDashMNRewards() {
	//查询vds的主节点vps列表
	mnlist, err := dao.NodeDao.GetMasternodeByCoinName(db.Factoty.GetSession(), "dash")
	if err != nil {
		log.Error("GetMasternodeByCoinName", err.Error())
		return
	}

	for _, v := range mnlist {
		if v.MNPayee != "" {
			//vpsArray = append(vpsArray, v.Vps)
			//获取rewards
			res, err := http.GetDashMNRewards(v.MNPayee)
			if err != nil {
				log.Error("GetDashRewards:", err.Error())
				continue
			}

			log.Debug(v.MNPayee, " rewards: ", res)
			rewards := int64(res * 1000000)
			//更新rewards
			err = dao.NodeDao.UpdateMasternodeRewards(db.Factoty.GetSession(), v.MNPayee, rewards)
			if err != nil {
				log.Error("UpdateMasternodeRewards", err.Error())
				continue
			}
		}
	}
}

func SyncDashMNStatus() {
	//查询vds的主节点vps列表
	mnlist, err := dao.NodeDao.GetMasternodeByCoinName(db.Factoty.GetSession(), "dash")
	if err != nil {
		log.Error("GetMasternodeByCoinName", err.Error())
		return
	}

	for _, v := range mnlist {
		if v.Vps != "" {
			//获取主节点payee
			mnpayee, err := http.GetDashMNPayee(v.Vps)
			if err != nil {
				log.Error("GetDashMNStatus", err.Error())
				return
			}

			//获取主节点状态
			status, err := http.GetDashMNStatus(mnpayee)
			if err != nil {
				log.Error("GetDashMNStatus", err.Error())
				return
			}

			err = dao.NodeDao.UpdateMasternodeMNStatus(db.Factoty.GetSession(), v.Vps, mnpayee, 1, status)
			if err != nil {
				log.Error("UpdateMasternodeMNStatus:", err.Error())
			}
		}
	}
}
