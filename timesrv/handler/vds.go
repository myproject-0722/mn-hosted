package handler

import (
	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/http"
	log "github.com/sirupsen/logrus"
)

func SyncVdsMNRewards() {
	//查询vds的主节点vps列表
	mnlist, err := dao.NodeDao.GetMasternodeByCoinName(db.Factoty.GetSession(), "vds")
	if err != nil {
		log.Error("GetMasternodeByCoinName", err.Error())
		return
	}

	for _, v := range mnlist {
		if v.MNPayee != "" {
			//vpsArray = append(vpsArray, v.Vps)
			//获取rewards
			res, err := http.GetVdsMNRewards(v.MNPayee)
			if err != nil {
				log.Error("GetVdsRewards:", err.Error())
				continue
			}

			rewards := int64(res * 1000000)
			//更新rewards
			err = dao.NodeDao.UpdateMasternodeRewards(db.Factoty.GetSession(), v.MNPayee, rewards, "vds")
			if err != nil {
				log.Error("UpdateMasternodeRewards", err.Error())
				continue
			}
		}
	}
}

func SyncVdsMNStatus() {
	//查询vds的主节点vps列表
	mnlist, err := dao.NodeDao.GetMasternodeByCoinName(db.Factoty.GetSession(), "vds")
	if err != nil {
		log.Error("GetMasternodeByCoinName", err.Error())
		return
	}

	for _, v := range mnlist {
		if v.MNPayee != "" {
			//获取主节点状态
			status, err := http.GetVdsMNStatus(v.MNPayee)
			if err != nil {
				log.Error("GetVdsMNStatus", err.Error())
				return
			}

			err = dao.NodeDao.UpdateMasternodeMNStatus(db.Factoty.GetSession(), v.Vps, v.MNPayee, 1, status)
			if err != nil {
				log.Error("UpdateMasternodeMNStatus:", err.Error())
				return
			}
		} else {
			//获取txid
			o, err := dao.OrderDao.GetOrderItem(db.Factoty.GetSession(), v.OrderID)
			if err != nil {
				log.Error("SyncVdsMNStatus GetOrderItem Error: ", v.OrderID)
				return
			}

			//获取payee
			payee, err := http.GetVdsMNPayee(o.TxID)
			if err != nil {
				log.Error("GetVdsMNPayee", err.Error())
				return
			}

			status, err := http.GetVdsMNStatus(payee)
			if err != nil {
				log.Error("GetVdsMNStatus", err.Error())
				return
			}

			err = dao.NodeDao.UpdateMasternodeMNStatus(db.Factoty.GetSession(), v.Vps, payee, 1, status)
			if err != nil {
				log.Error("UpdateMasternodeMNStatus:", err.Error())
				return
			}
		}
	}
}
