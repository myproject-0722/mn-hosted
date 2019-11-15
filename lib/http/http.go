package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/myproject-0722/mn-hosted/conf"
)

func RpcCall(s string) (string, error) {
	str := strings.Fields(s)
	//url := "http://mnhosted:123456@127.0.0.1:18332"
	curl1 := `{"jsonrpc":"1.0","id":"curltest","method":"`
	curl2 := `","params":[`
	curl3 := `]}`

	var quest string
	switch len(str) {
	case 1:
		quest = fmt.Sprintln(curl1 + str[0] + curl2 + curl3)
	case 2:
		quest = fmt.Sprintln(curl1 + str[0] + curl2 + "\"" + str[1] + "\"" + curl3)
	case 3:
		quest = fmt.Sprintln(curl1 + str[0] + curl2 + "\"" + str[1] + "\"" + ",\"" + str[2] + "\"" + curl3)
	case 4:
		quest = fmt.Sprintln(curl1 + str[0] + curl2 + "\"" + str[1] + "\"" + ",\"" + str[2] + "\"" + "," + str[3] + curl3)
	}

	fmt.Println(quest)
	var jsonStr = []byte(quest)
	req, err := http.NewRequest("POST", conf.WalletBaseUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
		//panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return string(body), nil
}

/*
func GetRpcCallResult(cmd string) (string, error) {
	r, err := RpcCall(cmd)
	if err != nil {
		return "", err
	}

	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(r), &dat); err == nil {
		//fmt.Println(dat)
		return dat["result"].(string), nil
	}
	return "", err
}*/

func GetRpcCallResult(cmd string) (interface{}, error) {
	r, err := RpcCall(cmd)
	if err != nil {
		return 0, err
	}

	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(r), &dat); err == nil {
		res := dat["result"]
		return res, nil
	}
	return 0, err
}

// Request method:GET or POST
func Request(method string, reqURL string, reqBody string) (res []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, reqURL, strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	/*req.Header.Set("Content-Type", "content-type: text/plain;")
	username := conf.WalletUser
	password := conf.WalletPassword
	authString := username + ":" + password
	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(authString))
	req.Header.Set("Authorization", "Basic "+s64) //验证信息放入header*/

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if res, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	return res, err
}

func WalletRequest(reqBody string) ([]byte, error) {
	res, err := Request("POST", conf.WalletBaseUrl, reqBody)
	if err != nil {
		log.Println(res)
	}
	return res, nil
}
