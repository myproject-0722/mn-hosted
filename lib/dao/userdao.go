package dao

import (
	"github.com/myproject-0722/mn-hosted/lib/dbsession"
	"github.com/myproject-0722/mn-hosted/lib/model"
	log "github.com/sirupsen/logrus"
)

type userDao struct{}

var UserDao = new(userDao)

// Insert
func (*userDao) Add(session *dbsession.DBSession, account string, passwd string, walletaddress string) (int64, error) {
	result, err := session.Exec("insert into t_account(account, passwd, walletaddress) value(?, ?, ?)", account, passwd, walletaddress)
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

// Get
func (*userDao) Get(session *dbsession.DBSession, account string) int64 {
	row := session.QueryRow("select id from t_account where account = ?", account)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		//log.Error(err)
		return -1
	}

	return id
}

func (*userDao) GetUserByUserID(session *dbsession.DBSession, id int64) *model.User {
	row := session.QueryRow("select account,walletaddress from t_account where id = ?", id)
	user := new(model.User)
	err := row.Scan(&user.Account, &user.WalletAddress)
	if err != nil {
		//log.Error(err)
		return nil
	}

	return user
}

// Insert
func (*userDao) Check(session *dbsession.DBSession, account string, passwd string) (int64, error) {
	row := session.QueryRow("select id, passwd from t_account where account = ?", account)
	var dbPasswd string
	var id int64
	err := row.Scan(&id, &dbPasswd)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	if passwd != dbPasswd {
		log.Error("password error")
		return 0, err
	}
	return id, nil
}
