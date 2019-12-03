package main

import (
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
	"github.com/myproject-0722/mn-hosted/conf"
	"github.com/myproject-0722/mn-hosted/lib/auth"
	"github.com/myproject-0722/mn-hosted/lib/token"
)

//const name = "API gateway"
func init() {
	token := &token.Token{}
	token.InitConfig(conf.GetConsulHosts(), "micro", "config", "jwt-key", "key")

	//plugin.Register(cors.NewPlugin())

	plugin.Register(plugin.NewPlugin(
		plugin.WithName("auth"),
		plugin.WithHandler(
			auth.JWTAuthWrapper(token),
		),
	))
	/*
		plugin.Register(plugin.NewPlugin(
			plugin.WithName("tracer"),
			plugin.WithHandler(
				stdhttp.TracerWrapper,
			),
		))
		plugin.Register(plugin.NewPlugin(
			plugin.WithName("breaker"),
			plugin.WithHandler(
				hystrix.BreakerWrapper,
			),
		))
		plugin.Register(plugin.NewPlugin(
			plugin.WithName("metrics"),
			plugin.WithHandler(
				prometheus.MetricsWrapper,
			),
		))*/
}

func main() {
	//stdhttp.SetSamplingFrequency(50)
	/*t, io, err := tracer.NewTracer(name, "")
	if err != nil {
					log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)*/

	/*hystrixStreamHandler := ph.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "8000"), hystrixStreamHandler)*/

	cmd.Init()
}
