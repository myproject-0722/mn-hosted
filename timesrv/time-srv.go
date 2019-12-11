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

	exitProgram := make(chan bool)

	c := cron.New()
	spec := "*/60 * * * * ?"
	//var i int = 0
	c.AddFunc(spec, func() {
		//i++
		handler.CheckMasterNodeExpired()

		handler.CountCoins()

		//暂时注释掉
		//handler.UpdateMasternodeInfo()

		//handler.UpdateCoinsPrice()
		//if i >= 5 {
		//exitProgram <- true
		//}
	})
	c.Start()

	<-exitProgram
	return
}
