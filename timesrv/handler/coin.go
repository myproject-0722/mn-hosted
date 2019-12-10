package handler

import (
	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	log "github.com/sirupsen/logrus"
)

//统计各币种托管数及收益
func CountCoins() {
	coinlist, err := dao.NodeDao.GetAllCoinList(db.Factoty.GetSession())
	if err != nil {
		log.Error("db GetAllCoinList", err.Error())
		return
	}

	for _, v := range coinlist {
		coinName := v.CoinName

		node, err := dao.NodeDao.GetMasternodeCountByCoin(db.Factoty.GetSession(), coinName)
		if err != nil {
			log.Error("db GetMasternodeCountByCoin", err.Error())
			continue
		}

		err = dao.NodeDao.UpdateCoinCount(db.Factoty.GetSession(), coinName, node)
		if err != nil {
			log.Error("db UpdateCoinCount", err.Error())
			continue
		}
	}
}
