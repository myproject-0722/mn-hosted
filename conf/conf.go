package conf

//consul
var (
	ConsulAddresses string = "127.0.0.1:8500" //consul addresses(such as 127.0.0.1:8500;127.0.0.1:8600)
)

//mq、cache、db
var (
	NSQIP   string = "127.0.0.1:4150" //NSQ
	RedisIP string = "127.0.0.1:6379" //redis
	MySQL   string = "mnhosted:123456@tcp(localhost:3306)/mnhosted?charset=utf8&parseTime=true"
)

//wallet
var (
	WalletBaseUrl string = "http://mnhosted:123456@127.0.0.1:19998/" //wallet
	//WalletBaseUrl  string = "http://127.0.0.1:19998/" //wallet
	WalletUser      string = "mnhosted"
	WalletPassword  string = "123454"
	MNHostedAddress string = "yhWiRNXYYLdNGogh3EYBne6FcSR8dRAvuf"
)
