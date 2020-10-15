package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/zsais/go-gin-prometheus"
	"net/http"
)

func main(){

	defaultRoute := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(defaultRoute)

	defaultRoute.GET("/ping-secondary", replyToPing)
	defaultRoute.GET("/fail-secondary", replyToFail)

	err := http.ListenAndServe("localhost:8081", defaultRoute)
	if err != nil {
		log.Error("unable to start server..", err)
	}

	log.Info("secondary service started..")
}

func replyToPing(c *gin.Context){
	c.JSON(500, "pong-secondary!!")
}

func replyToFail(c *gin.Context){
	c.JSON(500, "pong-secondary-failed..")
}


