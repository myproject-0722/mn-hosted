package dao

import (
	"github.com/myproject-0722/mn-hosted/lib/dbsession"
	"github.com/myproject-0722/mn-hosted/lib/model"
	log "github.com/sirupsen/logrus"
)

type orderDao struct{}

var OrderDao = new(orderDao)

func (*orderDao) Insert(session *dbsession.DBSession, o *model.Order) (int64, error) {
	result, err := session.Exec("insert into t_order(userid, coinname, timetype, price, txid, status) value(?, ?, ?, ?, ?, ?)", o.UserID, o.CoinName, o.TimeType, o.Price, o.TxID, o.Status)
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
	row := session.QueryRow("select userid, coinname, timetype, price, txid, isrenew, status from t_order where orderid = ? ", orderID)
	item := new(model.Order)
	err := row.Scan(&item.UserID, &item.CoinName, &item.TimeType, &item.Price, &item.TxID, &item.IsRenew, &item.Status)
	if err != nil {
		log.Error(err)
		return nil, err
	}

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
