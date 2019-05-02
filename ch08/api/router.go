package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	trace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"log"
	"micro/ch08/api/service"
	"os"
	"time"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Zipkin())

	router.GET("/", service.IndexApi) //首页
	//router.POST("/person", service.AddPersonApi)    //新增
	//router.GET("/person/:id", service.GetPersonApi) //获取单条记录

	return router
}

func Zipkin() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		zipkin_addr := "http://localhost:9411/api/v1/spans"
		hostname, _ := os.Hostname()
		collector, err1 := zipkin.NewHTTPCollector(zipkin_addr)
		if err1 != nil {
			log.Fatalf("unable to create Zipkin HTTP collector: %v", err1)
			return
		}
		defer collector.Close()
		recorder := zipkin.NewRecorder(collector, false, hostname, "gin.cli")

		tracer, err2 := zipkin.NewTracer(
			recorder,
			zipkin.ClientServerSameSpan(true),
			zipkin.TraceID128Bit(true),
		)
		if err2 != nil {
			log.Fatalf("unable to create Zipkin tracer: %v", err2)
			return
		}
		opentracing.InitGlobalTracer(tracer)

		ctx, span, _ := trace.StartSpanFromContext(c.Request.Context(), c.HandlerName())
		defer span.Finish()
		c.Keys = make(map[string]interface{})
		c.Keys["ctx"] = ctx

		//请求之前
		c.Next()
		//请求之后
		latency := time.Since(t)
		fmt.Println("time: ", latency)
	}
}
