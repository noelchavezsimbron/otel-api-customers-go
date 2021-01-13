package tracer

import (
	"fmt"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
	"os"
)

func InitTracer() func() {
	jaegerHost := os.Getenv("OTEL_JAEGER_ENDPOINT")
	jaegerServiceName := os.Getenv("OTEL_JAEGER_SERVICE_NAME")
	jaegerExporter := os.Getenv("OTEL_EXPORTER")

	fmt.Printf("jaegerHost: %s\n", jaegerHost)
	fmt.Printf("jaegerServiceName: %s\n", jaegerServiceName)
	fmt.Printf("jaegerExporter: %s\n", jaegerExporter)

	// Create and install Jaeger export pipeline
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(jaegerHost+"/api/traces"),
		jaeger.WithProcess(
			jaeger.Process{
				ServiceName: jaegerServiceName,
				Tags: []label.KeyValue{
					label.String("exporter", jaegerExporter),
					label.Float64("float", 312.23),
				},
			},
		),
		jaeger.WithSDK(
			&sdktrace.Config{
				DefaultSampler: sdktrace.AlwaysSample(),
			},
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	return flush
}
