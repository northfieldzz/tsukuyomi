package server

import (
	"github.com/gin-gonic/gin"
	"tsukuyomi/internal/tsukuyomi/controllers"
	"tsukuyomi/internal/tsukuyomi/line"
)

func NewRouter() (*gin.Engine, error) {
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return nil, err
	}
	router.Use()
	router.GET("/", controllers.Index)
	router.GET("/version", controllers.Version)
	webhookGroup := router.Group("webhook")
	{
		lineGroup := webhookGroup.Group("line")
		{
			lineGroup.POST("/callback", line.Callback)
		}
	}
	return router, nil
}
