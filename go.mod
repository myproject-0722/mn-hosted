module github.com/myproject-0722/mn-hosted

go 1.12

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/bitly/go-simplejson v0.5.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/micro/go-micro v1.16.0
	github.com/micro/go-plugins v1.5.1
	github.com/micro/micro v1.16.0
	github.com/myproject-0722/my-micro v0.0.0-20191024020128-780e2e7d7c45
	github.com/opentracing/opentracing-go v1.1.0
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/robfig/cron v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/smallnest/rpcx v0.0.0-20191115100340-4c760a7be45d
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/yaml.v2 v2.2.4
)

replace github.com/micro/micro => github.com/micro/micro v1.5.0

replace github.com/micro/go-micro => github.com/micro/go-micro v1.5.0
