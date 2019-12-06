package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/myproject-0722/mn-hosted/lib/pay"
	order "github.com/myproject-0722/mn-hosted/proto/order"
	log "github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
)

/*type Order struct {
	Client order.OrderService
}*/
var Client order.OrderService

//var OrderServ Order
/*
func Test() {
	response, err := Client.ConfirmAlipay(context.Background(), &order.ConfirmAlipayRequest{
		OrderID: 20,
		Price:   1,
	})
	if err != nil || response.Rescode != 200 {
		log.Error("ConfirmAlipay:", err.Error())
		return
	}
}
*/
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
			log.Debug("支付成功 tradeno:", noti.TradeNo, " outtraceno:", noti.OutTradeNo, " amount:", noti.TotalAmount, " ", noti.BuyerPayAmount)
			//修改订单状态。。。。
			orderID, err := strconv.ParseInt(noti.OutTradeNo, 10, 64)
			if err != nil {
				log.Error("TradeNo parse:", orderID)
				return
			}

			amount, err := strconv.ParseFloat(noti.BuyerPayAmount, 64)
			if err != nil {
				log.Error("BuyerPayAmount parse:", amount)
				return
			}

			response, err := Client.ConfirmAlipay(context.Background(), &order.ConfirmAlipayRequest{
				OrderID: orderID,
				Price:   int32(amount * 100),
			})
			if err != nil || response.Rescode != 200 {
				log.Error("ConfirmAlipay:", err.Error())
				return
			}
		} else {
			log.Error("支付失败")
			return
		}

		alipay.AckNotification(rep) // 确认收到通知消息
	})

	fmt.Println("server start....")
	http.ListenAndServe(":8088", nil)
}
