package handler

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/http"
	"github.com/myproject-0722/mn-hosted/lib/model"
	log "github.com/sirupsen/logrus"
)

func ToStr(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

func SyncDashMNRewards() {
	//查询dash的主节点vps列表
	mnlist, err := dao.NodeDao.GetMasternodeByCoinName(db.Factoty.GetSession(), "dash")
	if err != nil {
		log.Error("GetMasternodeByCoinName", err.Error())
		return
	}

	for _, v := range mnlist {
		if v.MNPayee != "" {
			//vpsArray = append(vpsArray, v.Vps)
			//获取rewards
			rewards, err := GetDashRewardsByAddress(v.MNPayee)
			if err != nil {
				log.Error("GetDashRewardsByAddress:", err.Error())
				continue
			}

			//更新rewards
			err = dao.NodeDao.UpdateMasternodeRewards(db.Factoty.GetSession(), v.MNPayee, rewards)
			if err != nil {
				log.Error("UpdateMasternodeRewards", err.Error())
				continue
			}
		}
	}
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
		key := item["MasternodePubkey"].(string)
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

				err = dao.NodeDao.UpdateMasternodeMNStatus(db.Factoty.GetSession(), vps, key, 1, status.(string))
				if err != nil {
					log.Error("UpdateMasternodeMNStatus:", err.Error())
				}
			}
			//ex, _ := simplejson.NewJson([]byte(exStatus))
		} else {
			err = dao.NodeDao.UpdateMasternodeMNStatus(db.Factoty.GetSession(), vps, key, 0, "")
			if err != nil {
				log.Error("UpdateMasternodeMNStatus:", err.Error())
			}
		}
	}
}

var maxBlockID int32 = 0

func UpdateDashMNBlockData() {
	if maxBlockID == 0 {
		//从数据库中查找
		id, _ := dao.DashBlockDao.GetMaxBlockID(db.Factoty.GetSession())
		maxBlockID = id
	}

	//获取最近的blockdata
	res, err := http.GetDashBlockData()
	if err != nil {
		log.Error("GetDashBlockData:", err.Error())
		return
	}

	dat, _ := simplejson.NewJson(res)
	dataArray, _ := dat.Get("data").Get("blocks").Array()

	for _, v := range dataArray {
		item, _ := v.(map[string]interface{})
		id, _ := item["BlockId"].(json.Number).Int64()
		blockID := int32(id)
		if blockID <= maxBlockID {
			break
		}

		block := new(model.DashBlock)
		value, _ := item["BlockId"].(json.Number).Float64()
		block.BlockID = int32(value)
		block.MNPayee = item["BlockMNPayee"].(string)
		block.PoolPubkey = item["BlockPoolPubKey"].(string)
		value, _ = item["BlockSupplyValue"].(json.Number).Float64()
		block.BlockValue = int32(value * 1000000)
		value, _ = item["BlockMNValue"].(json.Number).Float64()
		block.MNValue = int32(value * 1000000)
		block.PoolValue = block.BlockValue - block.MNValue
		block.IsSuperBlock = item["IsSuperBlock"].(bool)

		_, err := dao.DashBlockDao.Insert(db.Factoty.GetSession(), block)
		if err != nil {
			log.Error("DashBlockDao Insert:", err.Error())
			return
		}
	}
}

func GetDashRewardsByAddress(address string) (int64, error) {
	mnvalue, err := dao.DashBlockDao.GetMNRewardsByAddress(db.Factoty.GetSession(), address)
	if err != nil {
		log.Error("GetMNRewardsByAddress:", err.Error())
		return 0, err
	}

	poolvalue, err := dao.DashBlockDao.GetPoolRewardsByAddress(db.Factoty.GetSession(), address)
	if err != nil {
		log.Error("GetPoolRewardsByAddress:", err.Error())
		return 0, err
	}

	sum := mnvalue + poolvalue
	return sum, err
}
