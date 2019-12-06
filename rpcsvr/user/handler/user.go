package handler

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/myproject-0722/mn-hosted/lib/dao"
	db "github.com/myproject-0722/mn-hosted/lib/db"
	"github.com/myproject-0722/mn-hosted/lib/mail"
	"github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/token"
	"github.com/myproject-0722/mn-hosted/lib/utils"
	user "github.com/myproject-0722/mn-hosted/proto/user"
	wallet "github.com/myproject-0722/mn-hosted/proto/wallet"
)

const issuer = "go.mnhosted.srv.auth"

type User struct {
	Client wallet.WalletService
	Token  *token.Token
}

func (s *User) SignUp(ctx context.Context, req *user.SignUpRequest, rsp *user.SignUpResponse) error {
	log.Info("Received SignUpRequest Name: ", req.Account, " Passwd: ", req.Passwd)

	accountid := dao.UserDao.Get(db.Factoty.GetSession(), req.Account)
	if accountid != -1 {
		rsp.Rescode = 406
		rsp.Msg = " Account Exsit"
		return nil
	}

	//改用协程执行,作测试
	/*
		address := make(chan string)
		success := make(chan bool)
		go func() {
			response, err := s.Client.New(ctx, &wallet.NewRequest{
				Account: req.Account,
			})
			if err != nil {
				success <- false
				return
			}
			success <- true
			address <- response.Address
		}()

		if <-success == false {
			rsp.Rescode = 500
			rsp.Msg = "Get wallet address  Error"
			return nil
		}
	*/
	id, err := dao.UserDao.Add(db.Factoty.GetSession(), req.Account, req.Passwd, "")
	if err != nil {
		rsp.Rescode = 500
		rsp.Msg = "Sql Server Error"
		return nil
	}
	rsp.Rescode = 200
	rsp.Msg = " SignUp OK!"
	rsp.Id = id
	return nil
}

func (s *User) SignIn(ctx context.Context, req *user.SignInRequest, rsp *user.SignInResponse) error {
	log.Info("Received SignInRequest Name: ", req.Account, " Passwd: ", req.Passwd)
	id, err := dao.UserDao.Check(db.Factoty.GetSession(), req.Account, req.Passwd)
	if err != nil {
		rsp.Rescode = 404
		rsp.Msg = " SignIn Error!"
		return nil
	}

	var tokenStr string
	expireTime := time.Now().Add(time.Hour * 24).Unix() // 1天后过期
	tokenStr, err = s.Token.Encode(issuer, req.Account, expireTime)
	if err != nil {
		return err
	}
	rsp.Token = tokenStr
	/*
		token := utils.GetRandomString(32)
		log.Print("token=", token)
		redisclient.Client.Set("userToken:"+fmt.Sprint(id), token, 0)*/

	rsp.Rescode = 200
	rsp.Msg = " SignIn OK!"
	rsp.Id = id
	//rsp.Token = token
	return nil
}

func (s *User) SignOut(ctx context.Context, req *user.SignOutRequest, rsp *user.SignOutResponse) error {
	log.Info("Received SignOut: ", req.UserID)

	user := dao.UserDao.GetUserByUserID(db.Factoty.GetSession(), req.UserID)
	if user == nil {
		rsp.Rescode = 404
		rsp.Msg = " Account Not Exsit"
		return nil
	}

	rsp.Rescode = 200
	return nil
}

func (s *User) GetInfo(ctx context.Context, req *user.GetInfoRequest, rsp *user.GetInfoResponse) error {
	log.Debug("Received GetInfo: ", req.UserID)

	user := dao.UserDao.GetUserByUserID(db.Factoty.GetSession(), req.UserID)
	if user == nil {
		rsp.Rescode = 404
		rsp.Msg = " Account Not Exsit"
		return nil
	}

	count := dao.NodeDao.GetMasternodeCount(db.Factoty.GetSession(), req.UserID)
	rsp.MNCount = count

	/*response, err := s.Client.GetBalance(ctx, &wallet.GetBalanceRequest{
		Account: user.Account,
	})
	if err != nil {
		rsp.Rescode = 500
		rsp.Msg = "Get wallet balance  Error"
		return nil
	}*/

	rsp.Rescode = 200
	rsp.Balance = 0
	//rsp.Balance = response.Balance
	rsp.Account = user.Account
	rsp.WalletAddress = ""
	//rsp.WalletAddress = user.WalletAddress
	return nil
}

func (s *User) MailCode(ctx context.Context, req *user.MailCodeRequest, rsp *user.MailCodeResponse) error {
	log.Info("Received MailCode: ", req.Account)

	user := dao.UserDao.GetUserByAccount(db.Factoty.GetSession(), req.Account)
	if user == nil {
		rsp.Rescode = 404
		return nil
	}

	authCode := utils.GetRandomString(6)
	log.Info("authCode=", authCode)
	redisclient.Client.Set("authCode:"+req.Account, authCode, 0)

	mailTo := []string{
		req.Account,
	}
	//邮件主题为"Hello"
	subject := "重置邮箱验证码"
	// 邮件正文
	body := authCode
	err := mail.SendMail(mailTo, subject, body)
	if err != nil {
		rsp.Rescode = 404
	}

	rsp.Rescode = 200
	return nil
}

func (s *User) Reset(ctx context.Context, req *user.ResetRequest, rsp *user.ResetResponse) error {
	log.Info("Received ResetPasswd Name: ", req.Account, " Passwd: ", req.Passwd, " Authcode:", req.Authcode)

	code := redisclient.Client.Get("authCode:" + req.Account)
	if code == nil || code.Val() != req.Authcode {
		rsp.Rescode = 404
		rsp.Msg = " Autocode err!"
		log.Error(" Autocode err")
	}

	err := dao.UserDao.UpdatePasswd(db.Factoty.GetSession(), req.Account, req.Passwd)
	if err != nil {
		rsp.Rescode = 404
		rsp.Msg = " UpdatePasswd Error!"
		return nil
	}

	rsp.Rescode = 200
	rsp.Msg = " Reset OK!"
	//rsp.Token = token
	return nil
}
