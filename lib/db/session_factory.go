package db

import (
	"github.com/myproject-0722/mn-hosted/lib/dbsession"

	"github.com/myproject-0722/mn-hosted/conf"

	_ "github.com/go-sql-driver/mysql"
)

var Factoty *dbsession.DBSessionFactory

func Init() {
	var err error
	Factoty, err = dbsession.NewDBSessionFactory("mysql", conf.MySQL)
	if err != nil {
		panic(err)
	}
}
