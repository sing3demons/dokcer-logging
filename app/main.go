package main

import (
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/sing3demons/app/serve"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		resp := gin.H{
			"message": "Hello World!",
		}
		logger.Info("message", zap.Any("data", resp))
		c.JSON(200, resp)
	})
	serve.ServeHttp(":8080", "app", r)
}
