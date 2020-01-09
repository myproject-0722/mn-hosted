package register

import (
	"strconv"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/go-plugins/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/myproject-0722/mn-hosted/conf"
	liblog "github.com/myproject-0722/mn-hosted/lib/log"
	"github.com/myproject-0722/mn-hosted/lib/topic"
	"github.com/myproject-0722/mn-hosted/lib/tracer"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
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

	brokerKafka := kafka.NewBroker(func(options *broker.Options) {
		options.Addrs = conf.GetKafkaHosts()
	})

	if err := brokerKafka.Connect(); err != nil {
		log.Error("Broker Connect error: ", err)
	}

	reg := NewRegistry()

	microService := micro.NewService(
		micro.Registry(reg),
		micro.Name(servername),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Broker(brokerKafka),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	bk := microService.Server().Options().Broker

	_, err := bk.Subscribe(topic.SetLogLevel, func(p broker.Event) error {
		brokerHeader := p.Message().Header
		level := brokerHeader["loglevel"]
		logLevel, err := strconv.Atoi(level)
		if err != nil {
			log.Error(servername, "Recv change log level error: ", err.Error())
			return nil
		}
		log.Info(servername, " Recv change log level:", logLevel)
		log.SetLevel(log.Level(logLevel))
		return nil
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	// optionally setup command line usage
	microService.Init()

	return microService
}
