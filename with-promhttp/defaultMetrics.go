package main

import "github.com/prometheus/client_golang/prometheus"

var incomingTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "",
		Subsystem: "test",
		Name: "requests_total",
		Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
	},
	[]string{"statusCode", "method", "handler", "host", "url"},
)

var outboundTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "",
		Subsystem: "test",
		Name: "cots_request_total",
		Help: "External API calls made",
	},
	[]string{"statusCode", "method", "handler", "host", "url"},
)

