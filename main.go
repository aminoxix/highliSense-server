package main

import (
	"net/http"

	"github.com/aminoxix/highliSense-server/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, "pong")
  })

  router.GET("/run", func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,  helpers.Scraper())
  })

  router.Run() // listen and serve on 0.0.0.0:8080
}