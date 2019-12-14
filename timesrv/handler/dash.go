package handler

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/http"
	log "github.com/sirupsen/logrus"
)

func ToStr(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

func SyncDashMNStatus() {
	//查询dash的主节点vps列表
	mnlist, err := dao.NodeDao.GetMasternodeByCoinName(db.Factoty.GetSession(), "dash")
	if err != nil {
		log.Error("GetMasternodeByCoinName", err.Error())
		return
	}

	var vpsArray []string
	for _, v := range mnlist {
		if v.Vps != "" {
			vpsArray = append(vpsArray, v.Vps)
		}
	}

	vpsList, _ := json.Marshal(vpsArray)
	log.Debug(string(vpsList))

	//获取主节点状态
	res, err := http.GetDashMNStatus(string(vpsList))
	if err != nil {
		log.Error("GetDashMNStatus", err.Error())
		return
	}

	dat, _ := simplejson.NewJson(res)
	dataArray, _ := dat.Get("data").Array()
	for _, v := range dataArray {
		item, _ := v.(map[string]interface{})
		ip := item["MasternodeIP"]
		port := item["MasternodePort"]
		key := item["MasternodePubkey"]
		activecount, _ := item["ActiveCount"].(json.Number).Int64()
		vps := ip.(string) + ":" + port.(string)
		if activecount != 0 {
			ex := item["ExStatus"]
			//exStatus := item["ExStatus"].(json.Number).String()
			exStatus := ex.([]interface{})
			//var ex []map[string]interface{}
			//err = json.Unmarshal([]byte(exStatus), &ex)
			if err == nil {
				item := exStatus[0].(map[string]interface{})
				status := item["StatusEx"]

				log.Info("ip:", ip, "port:", port, "key:", key, "activecount:", activecount, "status:", status)

				err = dao.NodeDao.UpdateMasternodeMNStatus(db.Factoty.GetSession(), vps, 1, status.(string))
				if err != nil {
					log.Error("UpdateMasternodeMNStatus:", err.Error())
				}
			}
			//ex, _ := simplejson.NewJson([]byte(exStatus))
		} else {
			err = dao.NodeDao.UpdateMasternodeMNStatus(db.Factoty.GetSession(), vps, 0, "")
			if err != nil {
				log.Error("UpdateMasternodeMNStatus:", err.Error())
			}
		}
	}
}
