package tracer

import (
	"io"
	"time"

	"github.com/myproject-0722/mn-hosted/conf"
	opentracing "github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer 创建一个jaeger Tracer
func NewTracer(servicename string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: servicename,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		return nil, nil, err
	}

	reporter := jaeger.NewRemoteReporter(sender)
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Reporter(reporter),
	)

	return tracer, closer, err
}

func InitTracer(servicename string) {
	t, io, err := NewTracer(servicename, conf.GetJaegerHost())
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
}