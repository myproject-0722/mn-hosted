package handler

import (
	"fmt"
	"net/http"

	"github.com/myproject-0722/mn-hosted/lib/pay"
	order "github.com/myproject-0722/mn-hosted/proto/order"
	"github.com/smartwalle/alipay/v3"
)

type Order struct {
	Client order.OrderService
}

func HttpNotifyServer() {
	//支付成功之后的返回URL页面
	http.HandleFunc("/return", func(rep http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		ok, err := pay.VerifySignAlipay(req)
		if err == nil && ok {
			rep.Write([]byte("支付成功"))
		}
	})
	//支付成功之后的通知页面
	http.HandleFunc("/alipay", func(rep http.ResponseWriter, req *http.Request) {
		var noti, _ = pay.Client.GetTradeNotification(req)
		if noti != nil {
			fmt.Println("支付成功")
			//修改订单状态。。。。
		} else {
			fmt.Println("支付失败")
		}
		alipay.AckNotification(rep) // 确认收到通知消息
	})

	fmt.Println("server start....")
	http.ListenAndServe(":8088", nil)
}
