package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/myproject-0722/mn-hosted/conf"
	log "github.com/sirupsen/logrus"
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
	req, err := http.NewRequest("POST", conf.GetWalletUrl(), bytes.NewBuffer(jsonStr))
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
func Request(method string, reqURL string, reqBody []byte) (res []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, reqURL, bytes.NewReader(reqBody))
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

func VpsRequest(api string, reqBody []byte) (interface{}, error) {
	r, err := Request("POST", conf.GetVpsUrl()+api, reqBody)
	if err != nil {
		log.Error("VpsRequest", err.Error())
		return "1", err
	}

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(r), &dat); err == nil {
		res := dat["Errno"]
		return res, nil
	}
	return "1", nil
}

func VpsRequestGet(api string, reqBody []byte) ([]byte, error) {
	r, err := Request("POST", conf.GetVpsUrl()+api, reqBody)
	if err != nil {
		log.Error("VpsRequest", err.Error())
		return nil, err
	}

	return r, nil
}

func AddVpsNode(orderid int64) bool {
	jsondata := make(map[string]interface{})
	strOrderID := strconv.FormatInt(orderid, 10)
	jsondata["id"] = strOrderID
	bytesData, err := json.Marshal(jsondata)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	res, err := VpsRequest("vps/new", bytesData)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	if res == "0" {
		//fmt.Println("添加主节点成功")
		return true
	}
	return false
}

func DelVpsNode(nodeid int64) bool {
	jsondata := make(map[string]interface{})
	strNodeID := strconv.FormatInt(nodeid, 10)
	jsondata["id"] = strNodeID
	bytesData, err := json.Marshal(jsondata)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	res, err := VpsRequest("vps/del", bytesData)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	if res == "0" {
		//fmt.Println("添加主节点成功")
		return true
	}
	return false
}

func UpdateVpsNode(nodeid int64) bool {
	jsondata := make(map[string]interface{})
	strNodeID := strconv.FormatInt(nodeid, 10)
	jsondata["id"] = strNodeID
	bytesData, err := json.Marshal(jsondata)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	res, err := VpsRequest("vps/update", bytesData)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	if res == "0" {
		//fmt.Println("添加主节点成功")
		return true
	}
	return false
}

func GetAllVps() []byte {
	jsondata := make(map[string]interface{})
	jsondata["id"] = "0"
	bytesData, err := json.Marshal(jsondata)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	res, err := VpsRequestGet("vps/get", bytesData)
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	return res
}

func GetIpByVpsID(vpsid int64) string {
	jsondata := make(map[string]interface{})
	strVpsID := strconv.FormatInt(vpsid, 10)
	jsondata["id"] = strVpsID
	bytesData, err := json.Marshal(jsondata)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	res, err := VpsRequestGet("vps/getallnodefromvps", bytesData)
	if err != nil {
		log.Error(err.Error())
		return ""
	}

	var dat map[string]interface{}
	err = json.Unmarshal(res, &dat)
	if err == nil {
		res := dat["Errno"]
		if res == 0 {
			return dat["vpsip"].(string)
		}
		return ""
	}
	return ""
}

func GetCoinsPrice() (string, error) {
	resp, error := http.Get("https://api.coincap.io/v2/assets?ids=bitcoin,dash")
	if error != nil {
		return "", error
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body), nil
}

/*
func GetDashMNStatus(ips string) ([]byte, error) {
	baseUrl := "https://www.dashninja.pl/api/masternodes?testnet=0&exstatus=1&balance=1&ips="
	url := baseUrl + ips
	resp, error := http.Get(url)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
	//fmt.Println(string(body))
	//return string(body), nil
}
*/
func GetDashBlockData() ([]byte, error) {
	baseUrl := "https://www.dashninja.pl/data/blocks24h-0.json"
	url := baseUrl
	resp, error := http.Get(url)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
	//fmt.Println(string(body))
	//return string(body), nil
}

func GetSnowgemMNRewards(address string) (float64, error) {
	baseUrl := "http://127.0.0.1:3002/ext/getbalance/"
	url := baseUrl + address
	resp, error := http.Get(url)
	if error != nil {
		return 0, error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return 0, error
	}

	rewards, error := strconv.ParseFloat(string(body), 64)
	if error != nil {
		return 0, error
	}
	return rewards, nil
}

func GetSnowgemMNStatus(address string) (string, error) {
	baseUrl := "http://127.0.0.1:3002/ext/masternodestatus/"
	url := baseUrl + address
	resp, error := http.Get(url)
	if error != nil {
		return "", error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return "", error
	}

	return string(body), nil
}

func GetVdsMNRewards(address string) (float64, error) {
	baseUrl := "http://127.0.0.1:3001/ext/getbalance/"
	url := baseUrl + address
	resp, error := http.Get(url)
	if error != nil {
		return 0, error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return 0, error
	}

	rewards, error := strconv.ParseFloat(string(body), 64)
	if error != nil {
		return 0, error
	}
	return rewards, nil
}

func GetVdsMNStatus(address string) (string, error) {
	baseUrl := "http://127.0.0.1:3001/ext/masternodestatus/"
	url := baseUrl + address
	resp, error := http.Get(url)
	if error != nil {
		return "", error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return "", error
	}

	return string(body), nil
}

func GetDashMNRewards(address string) (float64, error) {
	baseUrl := "http://127.0.0.1:3003/ext/getbalance/"
	url := baseUrl + address
	resp, error := http.Get(url)
	if error != nil {
		return 0, error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return 0, error
	}

	rewards, error := strconv.ParseFloat(string(body), 64)
	if error != nil {
		return 0, error
	}
	return rewards, nil
}

func GetDashMNStatus(address string) (string, error) {
	baseUrl := "http://127.0.0.1:3003/ext/masternodestatus/"
	url := baseUrl + address
	resp, error := http.Get(url)
	if error != nil {
		return "", error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return "", error
	}

	return string(body), nil
}

func GetDashMNPayee(address string) (string, error) {
	baseUrl := "http://127.0.0.1:3003/ext/masternodemnpayee/"
	url := baseUrl + address
	resp, error := http.Get(url)
	if error != nil {
		return "", error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return "", error
	}

	return string(body), nil
}

func GetVdsMNPayee(address string) (string, error) {
	baseUrl := "http://127.0.0.1:3001/ext/masternodemnpayee/"
	url := baseUrl + address
	resp, error := http.Get(url)
	if error != nil {
		return "", error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return "", error
	}

	return string(body), nil
}

func GetSnowgemMNPayee(address string) (string, error) {
	baseUrl := "http://127.0.0.1:3002/ext/masternodepayee/"
	url := baseUrl + address
	resp, error := http.Get(url)
	if error != nil {
		return "", error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return "", error
	}

	return string(body), nil
}

/*
package main

import (
  "fmt"
  "os"
  "path/filepath"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "api.coincap.io/v2/assets"
  method := "GET"

  client := &http.Client {
    CheckRedirect: func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
    },
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
  }
  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))
}*/
