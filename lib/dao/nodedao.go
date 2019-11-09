package dao

import (
	"database/sql"
	"time"

	"github.com/myproject-0722/mn-hosted/lib/dbsession"
	"github.com/myproject-0722/mn-hosted/lib/model"
	log "github.com/sirupsen/logrus"
)

type nodeDao struct{}

var NodeDao = new(nodeDao)

// get
func (*nodeDao) GetCoinList(session *dbsession.DBSession, pageNo int32, perPagenum int32) ([]*model.Coin, error) {
	start := (pageNo - 1) * perPagenum
	rows, err := session.Query("select id, coinname, mnrequired, mnprice, volume, roi, monthlyincome, mnhosted, createtime, updatetime from t_coinlist limit ? , ? ", start, perPagenum)
	if err != nil {
		return nil, err
	}
	coinlist := make([]*model.Coin, 0)
	for rows.Next() {
		coin := new(model.Coin)
		err := rows.Scan(&coin.Id, &coin.CoinName, &coin.MNRequired, &coin.MNPrice, &coin.Volume, &coin.Roi, &coin.MonthlyIncome, &coin.MNHosted, &coin.CreateTime, &coin.UpdateTime)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		coinlist = append(coinlist, coin)
	}

	return coinlist, nil
}

// get
func (*nodeDao) GetMasternode(session *dbsession.DBSession, coinname string, mnkey string) (*model.Masternode, error) {
	strSql := "select id, coinname, mnkey, userid, syncstatus, createtime, expiretime from t_masternode where coinname = '" + coinname + "' and mnkey= '" + mnkey + "'"
	log.Println("sql=", strSql)
	row := session.QueryRow(strSql)
	//row := session.QueryRow("select id, coinname, mnkey, userid, vps, status, txid from t_masternode where coinname = '?' and mnkey= '?'", coinname, mnkey)
	node := new(model.Masternode)
	err := row.Scan(&node.Id, &node.CoinName, &node.MNKey, &node.UserID, &node.SyncStatus, &node.CreateTime, &node.ExpireTime)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Error(err)
		return nil, err
	}

	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return node, nil
}

// udpate
func (*nodeDao) UpdateMasternodeSyncStatus(session *dbsession.DBSession, coinname string, mnkey string, status int32) error {
	result, err := session.Exec("update t_masternode set syncstatus = ? where coinname = '?' and mnkey = '?'", status, coinname, mnkey)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Println(result)
	return nil
}

// udpate
func (*nodeDao) UpdateMasternodeExpireTime(session *dbsession.DBSession, coinname string, mnkey string, expiretime time.Time) error {
	//strSql := "update t_masternode set expiretime = '" + expiretime.Unix() + "' where coinname = '" + coinname + "' and mnkey= '" + mnkey + "'"
	//log.Println("sql=", strSql)
	//result, err := session.Exec(strSql)
	result, err := session.Exec("update t_masternode set expiretime = ? where coinname = ? and mnkey = ?", expiretime, coinname, mnkey)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Println(result)
	return nil
}

// get
func (*nodeDao) GetMasternodeByUserID(session *dbsession.DBSession, userid int64) ([]*model.Masternode, error) {
	rows, err := session.Query("select id, coinname, mnkey, syncstatus, createtime, expiretime from t_masternode where userid = ?", userid)
	if err != nil {
		return nil, err
	}
	nodelist := make([]*model.Masternode, 0)
	for rows.Next() {
		node := new(model.Masternode)
		err = rows.Scan(&node.Id, &node.CoinName, &node.MNKey, &node.SyncStatus, &node.CreateTime, &node.ExpireTime)
		if err != nil {
			return nil, err
		}
		nodelist = append(nodelist, node)
	}
	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return nodelist, nil
}

// insert
func (*nodeDao) AddMasternode(session *dbsession.DBSession, node model.Masternode) (int64, error) {
	result, err := session.Exec("insert ignore into t_masternode(coinname, mnkey, userid, vps, dockerid, status, expiretime) values(?,?,?,?,?,?,?)",
		node.CoinName, node.MNKey, node.UserID, node.Vps, node.DockerID, node.Status, node.ExpireTime)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return id, nil
}
