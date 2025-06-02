package main

import (
	"net/http"

	"github.com/aminoxix/highliSense-server/configs"
	"github.com/aminoxix/highliSense-server/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.Use(configs.HandleCors)

  // grouping routes with prefix of:
  v1 := router.Group("/api/v1")

  v1.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, "pong")
  })

  v1.POST("/run", handlers.PostContent)

  router.Run() // listen and serve on 0.0.0.0:8080
}