package dao

import (
	"github.com/myproject-0722/mn-hosted/lib/dbsession"
	"github.com/myproject-0722/mn-hosted/lib/model"
	log "github.com/sirupsen/logrus"
)

type orderDao struct{}

var OrderDao = new(orderDao)

func (*orderDao) Insert(session *dbsession.DBSession, o *model.Order) (int64, error) {
	result, err := session.Exec("insert into t_order(userid, coinname, timetype, price, txid, status) value(?, ?, ?, ?, ?, ?)", o.UserID, o.Coinname, o.TimeType, o.Price, o.TxID, o.Status)
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

func (*orderDao) Update(session *dbsession.DBSession, id int64, mnkey string, status int32) error {
	result, err := session.Exec("update t_order set mnkey = ?, status = ? where id = ?", mnkey, status, id)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Println(result)
	return nil
}
