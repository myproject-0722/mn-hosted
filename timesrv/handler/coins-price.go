package handler

import (
	"fmt"

	"github.com/bitly/go-simplejson"

	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/http"
	"github.com/myproject-0722/mn-hosted/lib/model"
	log "github.com/sirupsen/logrus"
)

func UpdateCoinsPrice() {
	log.Debug("UpdateCoinPrice")

	res, err := http.GetCoinsPrice()
	if err != nil {
		log.Error(err.Error())
		return
	}
	/*
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(res), &dat); err == nil {
			//res := dat["Errno"]
			//return res, nil
			log.Error(err.Error())
		}*/

	dat, _ := simplejson.NewJson([]byte(res))
	dataArray, _ := dat.Get("data").Array()
	for _, v := range dataArray {
		//就在这里对di进行类型判断
		item, _ := v.(map[string]interface{})
		id := item["id"]
		priceUsd := item["priceUsd"]
		fmt.Println(id, priceUsd)

		coin := new(model.CoinsPrice)
		coin.Id = id.(int64)
		coin.Price = priceUsd.(int32)
		err := dao.NodeDao.UpdateCoinsPrice(db.Factoty.GetSession(), *coin)
		if err != nil {
			log.Error("UpdateCoinsPrice", err)
		}
	}

	fmt.Println(res)
}
