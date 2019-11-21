package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/myproject-0722/mn-hosted/lib/db"
	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	"github.com/myproject-0722/mn-hosted/lib/redisclient"
	"github.com/myproject-0722/mn-hosted/lib/register"
	user "github.com/myproject-0722/mn-hosted/proto/user"

	"context"
)

type User struct {
	Client user.UserService
}

func (s *User) SignIn(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received SignIn API request")

	account, ok := req.Get["account"]
	if !ok || len(account.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.user", "account cannot be blank")
	}

	passwd, ok := req.Get["passwd"]
	if !ok || len(passwd.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.user", "passwd cannot be blank")
	}

	response, err := s.Client.SignIn(ctx, &user.SignInRequest{
		Account: strings.Join(account.Values, " "),
		Passwd:  strings.Join(passwd.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = response.Rescode
	/*b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})*/
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	fmt.Println(rsp.Body)

	return nil
}

func (s *User) SignUp(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	account, ok := req.Get["account"]
	if !ok || len(account.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.user", "account cannot be blank")
	}

	passwd, ok := req.Get["passwd"]
	if !ok || len(account.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.user", "account cannot be blank")
	}

	response, err := s.Client.SignUp(ctx, &user.SignUpRequest{
		Account: strings.Join(account.Values, " "),
		Passwd:  strings.Join(passwd.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	/* b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	}) */
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	fmt.Println(rsp.Body)
	return nil
}

func (s *User) GetInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received GetInfo API request")

	userID, ok := req.Get["userid"]
	if !ok || len(userID.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.user", "userid cannot be blank")
	}

	strUserID := strings.Join(userID.Values, " ")
	intUserID, err := strconv.ParseInt(strUserID, 10, 64)
	if err != nil {
		//
		return errors.BadRequest("go.mnhosted.api.node", "userid cannot be error")
	}

	response, err := s.Client.GetInfo(ctx, &user.GetInfoRequest{
		UserID: intUserID,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = response.Rescode

	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	fmt.Println(rsp.Body)
	return nil
}

func (s *User) SignOut(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received SignOut API request")

	userID, ok := req.Get["userid"]
	if !ok || len(userID.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.user", "userid cannot be blank")
	}

	strUserID := strings.Join(userID.Values, " ")
	intUserID, err := strconv.ParseInt(strUserID, 10, 64)
	if err != nil {
		//
		return errors.BadRequest("go.mnhosted.api.user", "userid cannot be error")
	}

	response, err := s.Client.SignOut(ctx, &user.SignOutRequest{
		UserID: intUserID,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = response.Rescode

	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	fmt.Println(rsp.Body)
	return nil
}

func (s *User) MailCode(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received MailCode API request")

	account, ok := req.Get["account"]
	if !ok || len(account.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.user", "userid cannot be blank")
	}

	response, err := s.Client.MailCode(ctx, &user.MailCodeRequest{
		Account: strings.Join(account.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = response.Rescode

	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	fmt.Println(rsp.Body)
	return nil
}

func (s *User) Reset(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Reset API request")

	account, ok := req.Get["account"]
	if !ok || len(account.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.user", "userid cannot be blank")
	}

	passwd, ok := req.Get["password"]
	if !ok || len(passwd.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.user", "passwd cannot be blank")
	}

	authcode, ok := req.Get["authcode"]
	if !ok || len(authcode.Values) == 0 {
		return errors.BadRequest("go.mnhosted.srv.user", "authcode cannot be blank")
	}

	response, err := s.Client.Reset(ctx, &user.ResetRequest{
		Account:  strings.Join(account.Values, " "),
		Passwd:   strings.Join(passwd.Values, " "),
		Authcode: strings.Join(authcode.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = response.Rescode

	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	fmt.Println(rsp.Body)
	return nil
}

func main() {
	liblog.InitLog("/var/log/mn-hosted/rpcsvr/user", "user.log")
	db.Init()
	redisclient.Init()
	reg := register.NewRegistry()

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.mnhosted.api.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: user.NewUserService("go.mnhosted.srv.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
