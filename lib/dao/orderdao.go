package dao

import (
	"github.com/myproject-0722/mn-hosted/lib/dbsession"
	"github.com/myproject-0722/mn-hosted/lib/model"
	log "github.com/sirupsen/logrus"
)

type orderDao struct{}

var OrderDao = new(orderDao)

func (*orderDao) Insert(session *dbsession.DBSession, o *model.Order) (int64, error) {
	result, err := session.Exec("insert into t_order(userid, coinname, mnname, mnkey, timetype, price, isrenew, txid, txindex, status) value(?, ?, ?, ?, ?, ?, ?, ?)", o.UserID, o.CoinName, o.MNName, o.MNKey, o.TimeType, o.Price, o.IsRenew, o.TxID, o.TxIndex, o.Status)
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
func (*orderDao) GetOrderItem(session *dbsession.DBSession, orderID int64) (*model.Order, error) {
	row := session.QueryRow("select userid, coinname, mnkey, timetype, price, txid, isrenew, status from t_order where id = ? ", orderID)
	item := new(model.Order)
	err := row.Scan(&item.UserID, &item.CoinName, &item.MNKey, &item.TimeType, &item.Price, &item.TxID, &item.IsRenew, &item.Status)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return item, nil
}

// get
func (*orderDao) GetOrderListByUserID(session *dbsession.DBSession, userid int64) ([]*model.Order, error) {
	rows, err := session.Query("select id, coinname, mnkey, timetype, price, isrenew, status, createtime from t_order where userid = ?", userid)
	if err != nil {
		return nil, err
	}
	orderlist := make([]*model.Order, 0)
	for rows.Next() {
		item := new(model.Order)
		err = rows.Scan(&item.Id, &item.CoinName, &item.MNKey, &item.TimeType, &item.Price, &item.IsRenew, &item.Status, &item.CreateTime)
		if err != nil {
			return nil, err
		}
		orderlist = append(orderlist, item)
	}
	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return orderlist, nil
}

// get
func (*orderDao) GetInfoByUserID(session *dbsession.DBSession, userID int64) (*model.OrderInfo, error) {
	row := session.QueryRow("select count(id), sum(price) from t_order where userid = ? ", userID)
	item := new(model.OrderInfo)
	err := row.Scan(&item.Num, &item.Payout)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return item, nil
}

func (*orderDao) Update(session *dbsession.DBSession, id int64, mnkey string, status int32) error {
	result, err := session.Exec("update t_order set mnkey = ?, status = ? where id = ?", mnkey, status, id)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Println(result)
	return nil
}
