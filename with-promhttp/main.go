package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func main() {
	defaultRoute := gin.Default()

	prometheus.MustRegister(incomingTotal)
	prometheus.MustRegister(outboundTotal)
	defaultRoute.Use(prometheusMiddleware)

	defaultRoute.GET("/ping-primary", replyToPing)
	defaultRoute.GET("/metrics", prometheusHandler)
	defaultRoute.GET("/call-secondary", callSecondary)

	err := http.ListenAndServe("localhost:8080", defaultRoute)
	if err != nil {
		log.Error("unable to start server..", err)
	}

	log.Info("primary service started.....")
}

func replyToPing(c *gin.Context) {
	c.JSON(200, "pong-primary!!")
}

func callSecondary(c *gin.Context) {
	resp, responseBytes := httpClient("http://localhost:8081/ping-secondary")
	c.JSON(resp.StatusCode, string(responseBytes))
}

func prometheusHandler(c *gin.Context) {
	h := promhttp.Handler()
	h.ServeHTTP(c.Writer, c.Request)
	return
}

func prometheusMiddleware(c *gin.Context) {
	if c.Request.URL.String() == "/metrics" {
		c.Next()
		return
	}
	status := strconv.Itoa(c.Writer.Status())
	incomingTotal.
		WithLabelValues(status, c.Request.Method, c.HandlerName(), c.Request.Host, c.Request.URL.String()).
		Inc()
	return
}
