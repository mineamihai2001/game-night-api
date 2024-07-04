package tracing

import (
	"context"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"

	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func Init() *sdktrace.TracerProvider {

	tracerOptions := GetTracerProviderOptions()

	traceProvider := sdktrace.NewTracerProvider(tracerOptions...)
	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return traceProvider
}

func GetTracerProviderOptions() []sdktrace.TracerProviderOption {
	tracerOptions := []sdktrace.TracerProviderOption{
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	}

	var serviceName string
	var serviceVersion string

	if name, exists := os.LookupEnv("SERVICE_NAME"); exists {
		serviceName = name
	} else {
		serviceName = "service"
	}

	if version, exists := os.LookupEnv("SERVICE_VERSION"); exists {
		serviceVersion = version
	} else {
		serviceVersion = "service"
	}

	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("service.version", serviceVersion),
		),
	)
	if err != nil {
		log.Fatal("Could not set resources")
	}

	tracerOptions = append(tracerOptions, sdktrace.WithResource(resources))

	if _, exists := os.LookupEnv("OTEL_EXPORTERS"); !exists {
		return tracerOptions
	}

	exporters := os.Getenv("OTEL_EXPORTERS")
	exporters = strings.ReplaceAll(exporters, " ", "")
	exportersArr := strings.Split(exporters, ",")

	for _, e := range exportersArr {
		if e == "http" {
			httpExporter, err := otlptracehttp.New(
				context.Background(),
				otlptracehttp.WithHeaders(map[string]string{
					"Content-Type": "application/json",
				}),
				otlptracehttp.WithInsecure(),
			)

			if err != nil {
				log.Fatal("Could not create http exporter")
			}

			tracerOptions = append(tracerOptions, sdktrace.WithBatcher(httpExporter))
		}
		if e == "stdout" {
			stdoutExporter, err := stdout.New(stdout.WithPrettyPrint())

			if err != nil {
				log.Fatal("Could not create stdout exporter")
			}

			tracerOptions = append(tracerOptions, sdktrace.WithBatcher(stdoutExporter))
		}
	}

	return tracerOptions
}

func IgnorePaths(r *http.Request) bool {
	var ignoredPaths []string
	if val, exists := os.LookupEnv("OTEL_IGNORED_PATHS"); exists {
		val = strings.ReplaceAll(val, " ", "")
		ignoredPaths = strings.Split(val, ",")
	} else {
		ignoredPaths = make([]string, 0)
	}

	return slices.Index(ignoredPaths, r.URL.Path) == -1
}
