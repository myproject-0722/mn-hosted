package dao

import (
	"github.com/myproject-0722/mn-hosted/lib/dbsession"
	"github.com/myproject-0722/mn-hosted/lib/model"
	log "github.com/sirupsen/logrus"
)

type dashBlockDao struct{}

var DashBlockDao = new(dashBlockDao)

// get
func (*dashBlockDao) GetMaxBlockID(session *dbsession.DBSession) (int32, error) {
	row, err := session.Query("select max(blockid) from t_dashblock")
	if err != nil {
		log.Error("GetMaxBlockIndex:", err.Error())
		return 0, err
	}

	var maxBlockID int32
	err = row.Scan(&maxBlockID)
	if err != nil {
		log.Error("GetMaxBlockIndex:", err.Error())
		return 0, err
	}

	return maxBlockID, nil
}

func (*dashBlockDao) Insert(session *dbsession.DBSession, d *model.DashBlock) (int64, error) {
	//BlockID      int32     `json:"blockid"`      // 货币名称
	//MNPayee      string    `json:"mnpayee"`      // 货币名称
	//PoolPubkey   string    `json:"poolpubkey"`   //
	//BlockValue   int32     `json:"dprice"`       //
	//MNValue      int32     `json:"mnvalue"`      //
	//PoolValue    int32     `json:"poolvalue"`    //
	//IsSuperBlock bool      `json:"issuperblock"` //
	result, err := session.Exec("insert into t_dashblock(blockid, mnpayee, poolpubkey, blockvalue, mnvalue, poolvalue, issuperblock) value(?, ?, ?, ?, ?, ?, ?)", d.BlockID, d.MNPayee, d.PoolPubkey, d.BlockValue, d.MNValue, d.PoolValue, d.IsSuperBlock)
	if err != nil {
		log.Error(err)
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Error(err)
		return -1, err
	}
	return id, nil
}

// get
func (*dashBlockDao) GetMNRewardsByAddress(session *dbsession.DBSession, address string) (int64, error) {
	row, err := session.Query("select sum(mnvalue) from t_dashblock where mnpayee = ? ", address)
	if err != nil {
		log.Error("GetMNRewardsByAddress:", err.Error())
		return 0, err
	}

	var rewards int64
	err = row.Scan(&rewards)
	if err != nil {
		log.Error("GetMNRewardsByAddress:", err.Error())
		return 0, err
	}

	return rewards, nil
}

// get
func (*dashBlockDao) GetPoolRewardsByAddress(session *dbsession.DBSession, address string) (int64, error) {
	row, err := session.Query("select sum(poolvalue) from t_dashblock where poolpubkey = ? ", address)
	if err != nil {
		log.Error("GetPoolRewardsByAddress:", err.Error())
		return 0, err
	}

	var rewards int64
	err = row.Scan(&rewards)
	if err != nil {
		log.Error("GetPoolRewardsByAddress:", err.Error())
		return 0, err
	}

	return rewards, nil
}
