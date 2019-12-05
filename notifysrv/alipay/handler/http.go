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

type Order struct {
	Client order.OrderService
}

var OrderServ Order

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
			log.Debug("支付成功 tradeno:", noti.TradeNo, " outtraceno:", noti.OutTradeNo, " amount:", noti.TotalAmount, "", noti.BuyerPayAmount)
			//修改订单状态。。。。
			orderID, err := strconv.ParseInt(noti.TradeNo, 10, 64)
			if err != nil {
				log.Error("TradeNo parse:", orderID)
			}

			amount, err := strconv.Atoi(noti.BuyerPayAmount)
			if err != nil {
				log.Error("TradeNo parse:", amount)
			}

			response, err := OrderServ.Client.ConfirmAlipay(context.Background(), &order.ConfirmAlipayRequest{
				OrderID: orderID,
				Price:   int32(amount),
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
