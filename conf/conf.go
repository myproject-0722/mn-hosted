package conf

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

/*
//consul
var (
	ConsulAddresses string = "127.0.0.1:8500" //consul addresses(such as 127.0.0.1:8500;127.0.0.1:8600)
)

//mq、cache、db
var (
	NSQIP     string = "127.0.0.1:4150" //NSQ
	KafkaIP          = []string{"127.0.0.1:9092"}
	TracingIP string = "127.0.0.1:6831"
	RedisIP   string = "127.0.0.1:6379" //redis
	MySQL     string = "mnhosted:123456@tcp(localhost:3306)/mnhosted?charset=utf8&parseTime=true"
)

//wallet
var (
	WalletBaseUrl string = "http://mnhosted:123456@127.0.0.1:19998/" //wallet
	//WalletBaseUrl  string = "http://127.0.0.1:19998/" //wallet
	WalletUser      string = "mnhosted"
	WalletPassword  string = "123454"
	MNHostedAddress string = "yhWiRNXYYLdNGogh3EYBne6FcSR8dRAvuf"
)

//vps
var (
	VpsBaseUrl string = "http://127.0.0.1:18889/api/v1/"
)
*/

var LogLevelMap map[string]int

type Conf struct {
	Version string              `yaml:"version"`
	Host    map[string]string   `yaml:"host"`
	Hosts   map[string][]string `yaml:"hosts"`
	Wallet  map[string]string   `yaml:"wallet"`
	Vps     map[string]string   `yaml:"vps"`
	Log     map[string]string   `yaml:"log"`
}

var config Conf

/*
type wallet struct {
	BaseUrl  string `yaml:"baseurl"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Address  string `yaml:"mnhostedaddress"`
}*/

func GetVersion() string {
	return config.Version
}

func GetHost(name string) string {
	return config.Host[name]
}

func GetHosts(name string) []string {
	return config.Hosts[name]
}

func GetWallet(name string) string {
	return config.Wallet[name]
}

func GetVps(name string) string {
	return config.Vps[name]
}

func GetLog(name string) string {
	return config.Log[name]
}

func GetConsulHosts() []string {
	return GetHosts("consul")
}

func GetJaegerHost() string {
	return GetHost("jaeger")
}

func GetRedisHost() string {
	return GetHost("redis")
}

func GetMysqlUrl() string {
	return GetHost("mysql")
}

func GetKafkaHosts() []string {
	return GetHosts("kafka")
}

func GetWalletUrl() string {
	return GetWallet("baseurl")
}

func GetWalletUser() string {
	return GetWallet("user")
}

func GetWalletPassword() string {
	return GetWallet("password")
}

func GetWalletAddress() string {
	return GetWallet("mnhostedaddress")
}

func GetVpsUrl() string {
	return GetVps("baseurl")
}

func GetLogDir() string {
	return GetLog("basedir")
}

func GetLogLevel() int {
	return LogLevelMap[GetLog("level")]
}

func init() {
	prefixPath := os.Getenv("mnhosted-path")
	if prefixPath == "" {
		gopath := os.Getenv("GOPATH")
		prefixPath = gopath + "/src/mn-hosted"
	}
	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("%s/conf/mn-hosted.yaml", prefixPath))
	if err != nil {
		fmt.Println("read yaml config error: ", err)
	}
	err = yaml.UnmarshalStrict(yamlFile, &config)

	LogLevelMap = map[string]int{"panic": 0, "fatal": 1, "error": 2, "warn": 3, "info": 4, "debug": 5}
}
