package pay

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/smartwalle/alipay/v3"
)

var (
	// appId
	appId = "2016101600698927"
	// 应用公钥
	//aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArMR10yQY4067dFE4pbJWRcz2v1BYqVDNKb5T63cPMLkIX5s/Z5XvBn1sAvpoL74lJBzYCi6xcUJ1PWKKkkqOckWJ3urYe7MCaMjb1Hgko0PyShmqUuE79pqaoqGWx0p6KZjDfKHLgNPsBFYZ9YoKpz8Ef0oPE+prlwhDBSoNRCVcVNg+zrs9PRFiWCscD/qgo+j4cve4oGW6Vt7A8loKNk0nlRaV5v/haq8CfsSdQSrSr00tK8z7RZYGPd8490ay+QRNeo7bsxtalH0FZ5V+GOrkfED8Hx/9jo32TXZjhO+EMMgnxTSOQkVaefEm3HcHGFMVu+0GsNwr7MKSSA6fNQIDAQAB"
	// 应用私钥
	privateKey = "MIIEowIBAAKCAQEApX9WVY1cqWd39se64WhjEFRN9prSz+yu+3oN7blGrS8W98ZOtal93P77eTTbutc06mhjzPQNQCenceH66Ae2hkY6TBR02mq8nvN2QXeAqJ4shnMAppngRmN1m+T7cKf/HC7tkwwu4uVPyOtyplLsg0SUn19M/ClZ0/AOYzHs57ChNiVeH5EKyrujJlEb7Jqsnx8d1AuHXaBdhr44GaVW4qMdE3qCrPdYdXHhv4a9Bkiwi7OJNIdBTP7Tzcn8IyCIOoLIPVGVPcta6Ck/2PBixOrfI73JQKznRKUs40A48nfU+73opFtHMCMqbDzAJfdpqrhvMYOXz+JehRX6F0LPQwIDAQABAoIBAQCaRxiObFdjPKdikFKwaoVe5ZhAOZgoaLW+jMuLPtqZ+3nnxR/+zWAdsj1vgk0L4i7cDjBrEV+A3PaFfWpO/1Gx3qnd3nwIWNQ5QTCOWv6/MaTEOVTz+iJOu80ZZN7Y6GMzPLQQDp1uuuIjpQmd71O4EyiRYV/8+fdZUUG4SwRT/p3/LumWeT67i7UecCfF9dteeJgeftiA0BozszHuPyVYhI0qjcyZNtMFNACo/K4wW3cwRUdJnsGmeim4s/63pWIB1Qk4T/nYkGipvw0inmAHMmqFhTUzTP3vJ+f6m/lvL6ekBBu1viIf8/Bxi4UXw0BHgd2ePCv2Zxw4A56urxZhAoGBANo1aGD/7eLldxwIjszGBqxHEmShyetpepVkdLRW9m2I6raFJM6SLExCOzeRA8KlpwcXbUssSJ8wcq8sXzsPUvH1YRk+SjPtuKO8rUHGM3NIy7p9lvIfmBbOfYN07xhKYdycnGpJCyacXrmb4ghkVSmYIr1MKUAy9C/pyVlHFqzLAoGBAMIo5v0eJoYhnTDlkhQPIYaw2ZU9QYEjL1kcCa6HVhI9j+pk0m0VdQmNO55E0KKt5J2p9UWA/bqnJTiSk4ck7rVSvLgMxrlDoIAiYDuDIZ/0lwLTUus9iEr1h//oTdebZsnekxfBJ2G116FjvAeiWzENL+RUWpfDd8vNyu2ek9BpAoGAYojfe33WVDE+WgBbS4jYlo75dUvBvHZDDpbwREdIvCmpo4X4GvfS3RTDXNI1Gn5nMEKZ7eovWQMtpoCo+ChxUiV2FUoVg+GDER0wN5ViwlpK9QmlUeyGZzYTY3s4RIXCLzbhQvV8/ZB7DeGgbh2wfznd5hEwR3c64S/25kO9r4UCgYBQFoEZXYt0fn8RgVCdN5STs3U8yxSvCO1p61fPBwIo6f3oKIhn+JbbRseVxDrvL52Cr219qvR+Pp3q1QNHlqNkZel0XcjG+K9Gy2c4hSGkkkaMItEsOahziw37MD6TtgVTNZ0lCkaNVm3Io5QW7hCBjjf4DheETFuo1I1lMKk2KQKBgAiIAi6dc7Y9WMJ65BEfFx2cqvuueKKQP7YDhRWThCoplR/tVptdVl8YyWIoh4C60AcKR6F+56iVWlcDkmjGOA8gPO+bAsCOEoSEyR0lO2QrqIyt3VEjv+tSxA0dyIF+8kyMmQkYqjZxNWero4Z3+qhwduFVaYLHdgYU79vpi/dp"
	Client, _  = alipay.New(appId, privateKey, false)
)

func init() {
	Client.LoadAppPublicCertFromFile("/home/lixu/git/golang/src/mn-hosted/conf/alipay/appCertPublicKey_2016101600698927.crt")
	Client.LoadAliPayPublicCertFromFile("/home/lixu/git/golang/src/mn-hosted/conf/alipay/alipayCertPublicKey_RSA2.crt")
	Client.LoadAliPayRootCertFromFile("/home/lixu/git/golang/src/mn-hosted/conf/alipay/alipayRootCert.crt")
}

func VerifySignAlipay(req *http.Request) (bool, error) {
	return Client.VerifySign(req.Form)
}

//网站扫码支付
func WebPageAlipay(orderid int64, price int32) (string, error) {
	pay := alipay.TradePagePay{}
	// 支付宝回调地址（需要在支付宝后台配置）
	// 支付成功后，支付宝会发送一个POST消息到该地址
	pay.NotifyURL = "http://pay.vpubchain.cn:8088/alipay"
	// 支付成功之后，浏览器将会重定向到该 URL
	pay.ReturnURL = "http://pay.vpubchain.cn:8088/return"
	//支付标题
	pay.Subject = "支付宝支付测试"

	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = strconv.FormatInt(orderid, 10)
	fmt.Println("tradeNo: ", pay.OutTradeNo)
	//销售产品码，与支付宝签约的产品码名称,目前仅支持FAST_INSTANT_TRADE_PAY
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	//金额
	var amount float64 = float64(price / 100)
	pay.TotalAmount = strconv.FormatFloat(amount, 'E', -1, 64)

	url, err := Client.TradePagePay(pay)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	payURL := url.String()
	//这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	fmt.Println(payURL)

	return payURL, nil
	//打开默认浏览器
	//payURL = strings.Replace(payURL, "&", "^&", -1)
	//exec.Command("cmd", "/c", "start", payURL).Start()
}
