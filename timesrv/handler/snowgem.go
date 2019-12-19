package handler

import (
	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/http"
	log "github.com/sirupsen/logrus"
)

func SyncSnowgemMNRewards() {
	//查询Snowgem的主节点vps列表
	mnlist, err := dao.NodeDao.GetMasternodeByCoinName(db.Factoty.GetSession(), "snowgem")
	if err != nil {
		log.Error("GetMasternodeByCoinName", err.Error())
		return
	}

	for _, v := range mnlist {
		if v.MNPayee != "" {
			//vpsArray = append(vpsArray, v.Vps)
			//获取rewards
			res, err := http.GetSnowgemMNRewards(v.MNPayee)
			if err != nil {
				log.Error("GetSnowgemRewards:", err.Error())
				continue
			}

			rewards := int64(res * 1000000)
			//更新rewards
			err = dao.NodeDao.UpdateMasternodeRewards(db.Factoty.GetSession(), v.MNPayee, rewards, "dash")
			if err != nil {
				log.Error("UpdateMasternodeRewards", err.Error())
				continue
			}
		}
	}
}

func SyncSnowgemMNStatus() {
	//查询Snowgem的主节点vps列表
	mnlist, err := dao.NodeDao.GetMasternodeByCoinName(db.Factoty.GetSession(), "snowgem")
	if err != nil {
		log.Error("GetMasternodeByCoinName", err.Error())
		return
	}

	for _, v := range mnlist {
		if v.MNPayee != "" {
			//获取主节点状态
			status, err := http.GetSnowgemMNStatus(v.MNPayee)
			if err != nil {
				log.Error("GetSnowgemMNStatus", err.Error())
				return
			}

			err = dao.NodeDao.UpdateMasternodeMNStatus(db.Factoty.GetSession(), v.Vps, v.MNPayee, 1, status)
			if err != nil {
				log.Error("UpdateMasternodeMNStatus:", err.Error())
			}
		}
	}
}
