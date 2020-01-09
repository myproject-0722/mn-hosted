package main

import (
	"flag"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/myproject-0722/mn-hosted/conf"
	"github.com/myproject-0722/mn-hosted/lib/topic"
	log "github.com/sirupsen/logrus"
)

func main() {
	var level string

	flag.StringVar(&level, "l", "5", "loglevel 0-PanicLevel 1-FatalLevel 2-ErrorLevel 3-WarnLevel 4-InfoLevel 5-DebugLevel 6-TraceLevel")
	flag.Parse()

	bk := kafka.NewBroker(func(options *broker.Options) {
		options.Addrs = conf.GetKafkaHosts()
	})

	if err := bk.Connect(); err != nil {
		log.Error("Broker Connect error: ", err)
	}

	//service := register.NewMicroService("go.mnhosted.set.log.level")
	/*
		brokerKafka := kafka.NewBroker(func(options *broker.Options) {
			options.Addrs = conf.GetKafkaHosts()
		})

		micro.Broker(brokerKafka)
	*/
	//service.Init()

	//bk := service.Server().Options().Broker

	err := bk.Publish(topic.SetLogLevel, &broker.Message{
		Header: map[string]string{
			"loglevel": level,
		},
		Body: []byte(""),
	})

	if err != nil {
		log.Error(err.Error())
	}

	//defer bk.Close()
	/*if err := service.Run(); err != nil {
		log.Fatal(err)
	}*/
}
