package configs

import "github.com/gin-gonic/gin"

func HandleCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Credentials", "true")
    c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

    if c.Request.Method == "OPTIONS" {
        c.AbortWithStatus(204)
        return
    }

    c.Next()
}