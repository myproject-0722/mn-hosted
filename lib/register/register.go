package register

import (
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/myproject-0722/mn-hosted/conf"
	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	"github.com/myproject-0722/mn-hosted/lib/tracer"
	"github.com/opentracing/opentracing-go"
)

var microService micro.Service
var reg registry.Registry

func NewRegistry() registry.Registry {
	if reg != nil {
		return reg
	}

	reg = consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = conf.GetConsulHosts()
	})

	return reg
	/* reg := etcdv3.NewRegistry(func(op *registry.Options) {
	    op.Addrs = []string{
	    "http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
	   }
	})*/
}

func NewMicroService(servername string) micro.Service {
	liblog.InitLog(conf.GetLogDir(), servername+".log")
	tracer.InitTracer(servername)

	/*brokerKafka := kafka.NewBroker(func(options *broker.Options) {
					options.Addrs = conf.KafkaIP
	})
	if err := brokerKafka.Connect(); err != nil {
					log.Error("Broker Connect error: ", err)
	}*/

	reg := NewRegistry()

	microService := micro.NewService(
		micro.Registry(reg),
		micro.Name(servername),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		//micro.Broker(brokerKafka),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// optionally setup command line usage
	microService.Init()

	return microService
}
