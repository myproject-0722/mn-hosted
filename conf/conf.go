package conf

//consul
var (
	ConsulAddresses string = "127.0.0.1:8500" //consul addresses(such as 127.0.0.1:8500;127.0.0.1:8600)
)

//mq、cache、db
var (
	NSQIP   string = "127.0.0.1:4150" //NSQ
	RedisIP string = "127.0.0.1:6379" //redis
	MySQL          = "root:123456@tcp(localhost:3306)/mn-hosted?charset=utf8&parseTime=true"
)
