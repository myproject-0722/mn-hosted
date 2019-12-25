package main

import (
	db "github.com/myproject-0722/mn-hosted/lib/db"
	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	"github.com/myproject-0722/mn-hosted/timesrv/handler"
	"github.com/robfig/cron"
)

func main() {
	liblog.InitLog("/home/lixu/git/golang/src/mn-hosted/log", "timesvr.log")
	db.Init()

	//status, _ := http.GetVdsMNStatus("VcmjrcQvxUSgdgphLRotF8kgHJmVfKNMLAL")
	//log.Info("GetVdsMNStatus", status)
	//rewards, _ := http.GetVdsRewards("VcZiR7G8rUFFuFm9HqdEb7M7J69rZhfqQQX")
	//log.Info("GetVdsRewards", rewards)

	exitProgram := make(chan bool)

	c := cron.New()
	spec := "*/60 * * * * ?"
	//var i int = 0
	c.AddFunc(spec, func() {
		//i++
		handler.CheckMasterNodeExpired()

		handler.CountCoins()

		handler.UpdateMasternodeInfo()

		//handler.UpdateCoinsPrice()
		//if i >= 5 {
		//exitProgram <- true
		//}
	})

	spec = "*/60 * * * * ?"
	c.AddFunc(spec, func() {

		handler.SyncDashMNStatus()

		handler.SyncDashMNRewards()

		handler.SyncVdsMNStatus()

		handler.SyncVdsMNRewards()

		handler.SyncSnowgemMNStatus()

		handler.SyncSnowgemMNRewards()

		handler.SyncDashMNStatus()

		handler.SyncDashMNRewards()

		//handler.UpdateDashMNBlockData()
	})
	c.Start()

	<-exitProgram
	return
}
