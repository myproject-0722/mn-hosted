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
func (*nodeDao) GetAllCoinList(session *dbsession.DBSession) ([]*model.Coin, error) {
	rows, err := session.Query("select id, coinname, mnrequired, dprice, mprice, yprice, volume, roi, monthlyincome, mnhosted, createtime, updatetime from t_coinlist")
	if err != nil {
		return nil, err
	}
	coinlist := make([]*model.Coin, 0)
	for rows.Next() {
		coin := new(model.Coin)
		err := rows.Scan(&coin.Id, &coin.CoinName, &coin.MNRequired, &coin.DPrice, &coin.MPrice, &coin.YPrice, &coin.Volume, &coin.Roi, &coin.MonthlyIncome, &coin.MNHosted, &coin.CreateTime, &coin.UpdateTime)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		coinlist = append(coinlist, coin)
	}

	return coinlist, nil
}

// get
func (*nodeDao) GetCoinList(session *dbsession.DBSession, pageNo int32, perPagenum int32) ([]*model.Coin, error) {
	start := (pageNo - 1) * perPagenum
	rows, err := session.Query("select id, coinname, mnrequired, dprice, mprice, yprice, volume, roi, monthlyincome, mnhosted, createtime, updatetime from t_coinlist limit ? , ? ", start, perPagenum)
	if err != nil {
		return nil, err
	}
	coinlist := make([]*model.Coin, 0)
	for rows.Next() {
		coin := new(model.Coin)
		err := rows.Scan(&coin.Id, &coin.CoinName, &coin.MNRequired, &coin.DPrice, &coin.MPrice, &coin.YPrice, &coin.Volume, &coin.Roi, &coin.MonthlyIncome, &coin.MNHosted, &coin.CreateTime, &coin.UpdateTime)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		coinlist = append(coinlist, coin)
	}

	return coinlist, nil
}

// udpate coins count
func (*nodeDao) UpdateCoinCount(session *dbsession.DBSession, coinname string, node *model.MasternodeCount) error {
	_, err := session.Exec("update t_coinlist set volume = ? , mnhosted = ? where coinname = ?", node.Earns, node.Count, coinname)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// get
func (*nodeDao) GetCoinItem(session *dbsession.DBSession, coinname string) (*model.Coin, error) {
	//strSql := "select id, coinname, mnrequired, dprice, mprice, yprice, volume, roi, monthlyincome, mnhosted, createtime, updatetime from t_coinlist where coinname = '" + coinname + "'"
	//log.Println("sql=", strSql)
	//row := session.QueryRow(strSql)
	row := session.QueryRow("select id, coinname, mnrequired, dprice, mprice, yprice, volume, roi, monthlyincome, mnhosted, createtime, updatetime from t_coinlist where coinname = ? ", coinname)
	coin := new(model.Coin)
	err := row.Scan(&coin.Id, &coin.CoinName, &coin.MNRequired, &coin.DPrice, &coin.MPrice, &coin.YPrice, &coin.Volume, &coin.Roi, &coin.MonthlyIncome, &coin.MNHosted, &coin.CreateTime, &coin.UpdateTime)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return coin, nil
}

// get
func (*nodeDao) GetMasternode(session *dbsession.DBSession, coinname string, mnkey string) (*model.Masternode, error) {
	//strSql := "select id, coinname, mnkey, userid, syncstatus, createtime, expiretime from t_masternode where coinname = '" + coinname + "' and mnkey= '" + mnkey + "'"
	//log.Println("sql=", strSql)
	//row := session.QueryRow(strSql)
	row := session.QueryRow("select id, coinname, mnkey, userid, syncstatus, createtime, expiretime from t_masternode where coinname = ? and mnkey = ? ", coinname, mnkey)
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

// 获取相关币种主节点的统计，正常分布式项目计算不应该放在数据库，不过此项目数据库压力较小，计算可放在数据库内统计
func (*nodeDao) GetMasternodeCountByCoin(session *dbsession.DBSession, coinname string) (*model.MasternodeCount, error) {
	row := session.QueryRow("select count(id), sum(earn) from t_masternode where coinname = ? ", coinname)
	node := new(model.MasternodeCount)
	err := row.Scan(&node.Count, &node.Earns)
	if err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		log.Error(err)
		return nil, err
	}

	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return node, nil
}

// get
func (*nodeDao) GetMasternodeCount(session *dbsession.DBSession, userID int64) int32 {
	row := session.QueryRow("select count(id) from t_masternode where userid = ?", userID)
	var count int32
	err := row.Scan(&count)

	if err != nil {
		log.Error(err)
		return 0
	}

	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return count
}

// get
func (*nodeDao) GetNodeIDByOrderID(session *dbsession.DBSession, orderid int64) int64 {
	row := session.QueryRow("select id from t_masternode where orderid = ?", orderid)
	var id int64
	err := row.Scan(&id)

	if err != nil {
		log.Error(err)
		return -1
	}

	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return id
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
func (*nodeDao) UpdateMasternodeMNStatus(session *dbsession.DBSession, vps string, key string, mnstatus int32, mnstatusex string) error {
	result, err := session.Exec("update t_masternode set mnpayee = ? ,mnstatus = ? , mnstatusex = ? where vps = ? ", key, mnstatus, mnstatusex, vps)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Println(result)
	return nil
}

// udpate
func (*nodeDao) UpdateMasternodeRewards(session *dbsession.DBSession, mnpayee string, rewards int64) error {
	result, err := session.Exec("update t_masternode set rewards = ? where mnpayee = ? ", rewards, mnpayee)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Println(result)
	return nil
}

// udpate
func (*nodeDao) UpdateMasternodeStatus(session *dbsession.DBSession, id int64, status int32) error {
	result, err := session.Exec("update t_masternode set status = ? where id = ? ", status, id)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Println(result)
	return nil
}

// udpate
func (*nodeDao) UpdateMasternodeVpsInfo(session *dbsession.DBSession, node *model.Masternode) error {
	result, err := session.Exec("update t_masternode set syncstatus = 1, vps = ? where id = ? ", node.Vps, node.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Println(result)
	return nil
}

// delete
func (*nodeDao) DelMasternodeByID(session *dbsession.DBSession, id int64) error {
	result, err := session.Exec("delete from t_masternode where id = ?", id)
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
	rows, err := session.Query("select id, coinname, mnkey, vps, earn, status, syncstatus, mnstatus, createtime, expiretime from t_masternode where userid = ?", userid)
	if err != nil {
		return nil, err
	}
	nodelist := make([]*model.Masternode, 0)
	for rows.Next() {
		node := new(model.Masternode)
		err = rows.Scan(&node.Id, &node.CoinName, &node.MNKey, &node.Vps, &node.Earn, &node.Status, &node.SyncStatus, &node.MNStatus, &node.CreateTime, &node.ExpireTime)
		if err != nil {
			return nil, err
		}
		nodelist = append(nodelist, node)
	}
	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return nodelist, nil
}

// get
func (*nodeDao) GetMasternodeByCoinName(session *dbsession.DBSession, coinname string) ([]*model.Masternode, error) {
	rows, err := session.Query("select id, coinname, mnkey, mnpayee, vps, earn, status, syncstatus, mnstatus, createtime, expiretime from t_masternode where coinname = ?", coinname)
	if err != nil {
		return nil, err
	}
	nodelist := make([]*model.Masternode, 0)
	for rows.Next() {
		node := new(model.Masternode)
		err = rows.Scan(&node.Id, &node.CoinName, &node.MNKey, &node.MNPayee, &node.Vps, &node.Earn, &node.Status, &node.SyncStatus, &node.MNStatus, &node.CreateTime, &node.ExpireTime)
		if err != nil {
			return nil, err
		}
		nodelist = append(nodelist, node)
	}
	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return nodelist, nil
}

// get
func (*nodeDao) GetExpiredTimeMasternode(session *dbsession.DBSession, expiretime time.Time) ([]*model.Masternode, error) {
	rows, err := session.Query("select id, coinname, mnkey, userid, orderid, status, syncstatus, mnstatus, createtime, expiretime, updatetime from t_masternode where status = 2 and expiretime <= ?", expiretime)
	if err != nil {
		return nil, err
	}
	nodelist := make([]*model.Masternode, 0)
	for rows.Next() {
		node := new(model.Masternode)
		err = rows.Scan(&node.Id, &node.CoinName, &node.MNKey, &node.UserID, &node.OrderID, &node.Status, &node.SyncStatus, &node.MNStatus, &node.CreateTime, &node.ExpireTime, &node.UpdateTime)
		if err != nil {
			return nil, err
		}
		nodelist = append(nodelist, node)
	}
	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return nodelist, nil
}

// get
func (*nodeDao) GetMasternodeBySyncStatus(session *dbsession.DBSession, syncstatus int32) ([]*model.Masternode, error) {
	rows, err := session.Query("select id, coinname, mnkey, userid, orderid, status, syncstatus, mnstatus, createtime, expiretime, updatetime from t_masternode where status = ? ", syncstatus)
	if err != nil {
		return nil, err
	}
	nodelist := make([]*model.Masternode, 0)
	for rows.Next() {
		node := new(model.Masternode)
		err = rows.Scan(&node.Id, &node.CoinName, &node.MNKey, &node.UserID, &node.OrderID, &node.Status, &node.SyncStatus, &node.MNStatus, &node.CreateTime, &node.ExpireTime, &node.UpdateTime)
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
	result, err := session.Exec("insert ignore into t_masternode(coinname, mnkey, userid, orderid, status, createtime, expiretime) values(?,?,?,?,?,?,?)",
		node.CoinName, node.MNKey, node.UserID, node.OrderID, node.Status, node.CreateTime, node.ExpireTime)
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

/*
func (*nodeDao) BackupMasternode(session *dbsession.DBSession, node model.Masternode) error {
	result, err := session.Exec("insert ignore into t_masternode_backup(id, coinname, mnkey, userid, orderid, status, createtime, expiretime) values(?,?,?,?,?,?,?,?)",
		node.Id, node.CoinName, node.MNKey, node.UserID, node.OrderID, node.Status, node.CreateTime, node.ExpireTime)
	if err != nil {
		log.Error(err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("backup:", id)
	return nil
}*/

// insert or update
func (*nodeDao) UpdateCoinsPrice(session *dbsession.DBSession, item model.CoinsPrice) error {

	_, err := session.Exec("insert into t_coinsprice(coinname, price) values(?,?) on duplicate key update coinname = ?",
		item.CoinName, item.Price, item.CoinName)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// get
func (*nodeDao) GetNodeByOrderID(session *dbsession.DBSession, orderid int64) (*model.Node, error) {
	row := session.QueryRow("select port from t_node where order_id = ? ", orderid)
	node := new(model.Node)
	err := row.Scan(&node.Port)
	if err != nil {
		return nil, err
	}

	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return node, nil
}
