package handler

import (
	//"strconv"

	"strconv"
	"time"

	"github.com/myproject-0722/mn-hosted/lib/dao"
	"github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/http"

	//	"github.com/myproject-0722/mn-hosted/lib/http"
	log "github.com/sirupsen/logrus"
)

func getVpsIpByVpsID(vpsinfo []interface{}, vpsid int64) string {
	//b, _ := json.Marshal(dat)
	//strData := string(b)
	//fmt.Println("strData:", strData)
	//vpsJson, _ := simplejson.NewJson([]byte(res))
	//vpsinfo, _ := vpsJson.Get("Vps").Array()

	for _, v := range vpsinfo {
		item, _ := v.(map[string]interface{})
		if item["id"] == vpsid {
			return item["ipAddress"].(string)
		}
	}
	return ""
}

func UpdateMasternodeInfo() {
	log.Debug("UpdateMasternodeInfo")
	/* 不从vps获取信息了
	//获取vps信息
	vpsResp := http.GetAllVps()
	if vpsResp == nil {
		log.Error("GetAllVps nil")
		return
	}

	vpsJson, err := simplejson.NewJson([]byte(vpsResp))
	if err != nil {
		log.Error("GetAllVps ", err.Error())
		return
	}

	res, err := vpsJson.Get("Errno").String()
	if res != "0" {
		errmsg, _ := vpsJson.Get("Errmsg").String()
		log.Error("GetAllVps ", errmsg)
		return
	}
	vpsinfo, err := vpsJson.Get("Vps").Array()
	if err != nil {
		log.Error("GetAllVps:", err.Error())
		return
	}
	*/
	//获取同步未完成的主节点
	nodelist, err := dao.NodeDao.GetMasternodeBySyncStatus(db.Factoty.GetSession(), 0)
	if err != nil {
		log.Error("UpdateMasternodeInfo", err)
		return
	}

	for _, v := range nodelist {
		log.Debug("UpdateMasternodeVpsInfo orderid:", v.OrderID, "vps:", v.Vps)

		//查看node表中是否有纪录
		node, err := dao.NodeDao.GetNodeByOrderID(db.Factoty.GetSession(), v.OrderID)
		if err != nil {
			continue
		}

		//跟据vpsid获取vpsip
		//ip := getVpsIpByVpsID(vpsinfo, node.VpsID)

		v.Vps = node.PublicIP + ":" + strconv.Itoa(int(node.Port))

		log.Debug("UpdateMasternodeVpsInfo orderid:", v.OrderID, "vps:", v.Vps)
		err = dao.NodeDao.UpdateMasternodeVpsInfo(db.Factoty.GetSession(), v)
		if err != nil {
			log.Error("UpdateMasternodeVpsInfo:", err.Error())
			continue
		}
	}
}

func CheckMasterNodeExpired() {
	log.Debug("CheckMasterNodeExpired")
	nodelist, err := dao.NodeDao.GetExpiredTimeMasternode(db.Factoty.GetSession(), time.Now())
	if err != nil {
		log.Error("GetExpiredTimeMasternode", err)
		return
	}

	for i, v := range nodelist {
		nodeid := dao.NodeDao.GetNodeIDByOrderID(db.Factoty.GetSession(), v.OrderID)
		if nodeid == -1 {
			continue
		}

		if http.DelVpsNode(nodeid) == false {
			continue
		}

		//设置主节点已过期标志
		delerr := dao.NodeDao.UpdateMasternodeStatus(db.Factoty.GetSession(), v.Id, 2)
		if delerr != nil {
			log.Fatal("DelMasternodeByID orderid=", v.OrderID, i, delerr)
		}
	}
}
