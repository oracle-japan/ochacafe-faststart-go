package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/db"
	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/http"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var tracerExp *otlptrace.Exporter

func retryInitTracer() func() {
	var shutdown func()
	go func() {
		for {
			// otel will reconnected and re-send spans when otel col recover. so, we don't need to re-init tracer exporter.
			if tracerExp == nil {
				shutdown = initTracer()
			} else {
				break
			}
			time.Sleep(time.Minute * 5)
		}
	}()
	return shutdown
}

func initTracer() func() {
	// temporarily set timeout to 10s
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	otlpHTTPExporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithURLPath("/20200101/opentelemetry/private/v1/traces"),
		otlptracehttp.WithHeaders(map[string]string{
			"Authorization": "dataKey " + os.Getenv("APM_PRIVATE_DATA_KEY"),
		}),
	)
	if err != nil {
		handleErr(err, "OTLP Trace Creation")
		return nil
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(otlpHTTPExporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL)))

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	tracerExp = otlpHTTPExporter
	return func() {
		// Shutdown will flush any remaining spans and shut down the exporter.
		handleErr(tracerProvider.Shutdown(ctx), "failed to shutdown TracerProvider")
	}
}

func handleErr(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %v", message, err)
	}
}

func main() {
	if shutdown := retryInitTracer(); shutdown != nil {
		defer shutdown()
	}
	db.SetupDB()
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(otelgin.Middleware("ochacafe-app"))
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"POST",
			"GET",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	router.GET("/items", http.GetAll)
	router.GET("/items/:id", http.GetItemById)
	router.POST("/items", http.UpdateItem)
	router.DELETE("/items/:id", http.DeleteItem)
	return router
}
