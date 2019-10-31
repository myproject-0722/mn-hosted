package dao

import (
	"fmt"

	"github.com/myproject-0722/mn-hosted/lib/dbsession"
	"github.com/myproject-0722/mn-hosted/lib/model"
	log "github.com/sirupsen/logrus"
)

type nodeDao struct{}

var NodeDao = new(nodeDao)

// Insert
func (*nodeDao) GetCoinList(session *dbsession.DBSession, pageNo int32, perPagenum int32) (*model.Coin, error) {
	start := (pageNo - 1) * perPagenum
	row := session.QueryRow("select id, coinname, mnrequired, mnprice, volume, roi, monthlyincome, mnhosted, createtime, updatetime from t_coinlist limit ? , ? ", start, perPagenum)
	coin := new(model.Coin)
	err := row.Scan(&coin.Id, &coin.CoinName, &coin.MNRequired, &coin.MNPrice, &coin.Volume, &coin.Roi, &coin.MonthlyIncome, &coin.MNHosted, &coin.CreateTime, &coin.UpdateTime)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	fmt.Println(coin.Id, coin.MNPrice, coin.MNRequired, coin.Volume)
	return coin, nil
}
