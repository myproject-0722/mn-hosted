package dao

import (
	"github.com/myproject-0722/mn-hosted/lib/dbsession"
	log "github.com/sirupsen/logrus"
)

type userDao struct{}

var UserDao = new(userDao)

// Insert
func (*userDao) Add(session *dbsession.DBSession, account string, passwd string) (int64, error) {
	result, err := session.Exec("insert into t_account(account, passwd) value(?, ?)", account, passwd)
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

// Insert
func (*userDao) Check(session *dbsession.DBSession, account string, passwd string) (bool, error) {
	row := session.QueryRow("select passwd from t_account where account = ?", account)
	var dbPasswd string
	err := row.Scan(&dbPasswd)
	if err != nil {
		log.Error(err)
		return false, err
	}
	if passwd != dbPasswd {
		log.Error("password error")
		return false, err
	}
	return true, nil
}
