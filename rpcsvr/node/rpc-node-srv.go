package main

import (
	log "github.com/sirupsen/logrus"

	db "github.com/myproject-0722/mn-hosted/lib/db"
	redisclient "github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/register"
	node "github.com/myproject-0722/mn-hosted/proto/node"
	"github.com/myproject-0722/mn-hosted/rpcsvr/node/handler"
)

/*
func GetNodeSyncStatus(userid int64, mnkey string, dockerid string) bool {
	var statusCmd string = "docker exec " + dockerid + " /bin/bash -c 'dash-cli -testnet mnsync status'"
	log.Print("statuscmd=", statusCmd)
	syncStatus := cmd.ExecShell(statusCmd)
	log.Print("mnstatus=", syncStatus)

	if syncStatus != "" {
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(syncStatus), &data); err == nil {
			log.Println("==============json str 转map=======================")
			log.Println(data)
			log.Println(data["IsBlockchainSynced"])
			if data["IsBlockchainSynced"] == true {
				return true
			}
		}
	}

	return false
}*/

func main() {
	service := register.NewMicroService("go.mnhosted.srv.node")
	db.Init()
	redisclient.Init()

	// Register Handlers
	node.RegisterMasternodeHandler(service.Server(), new(handler.Masternode))
	node.RegisterCoinHandler(service.Server(), new(handler.Coin))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
