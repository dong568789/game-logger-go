package router

import (
	v1 "game-logger-go/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	g := r.Group("api/v1")
	{
		//写日志
		g.POST("cy_log", v1.LogWrite)
	}
}
