package mail

import (
	"strconv"

	"github.com/smallnest/rpcx/log"
	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "jiuling-monitor@qq.com",
		"pass": "123456",
		"host": "smtp.qq.com",
		"port": "25",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "90BlockChain"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                                 //发送给多个用户
	m.SetHeader("Subject", subject)                              //设置邮件主题
	m.SetBody("text/html", body)                                 //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	if err != nil {
		log.Error("send mail err:", err.Error())
	}
	return err

}
