package server

import (
	"github.com/gin-gonic/gin"
	"tsukuyomi/controllers"
	"tsukuyomi/controllers/line"
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
			lineGroup.POST("/v1", line.WebhookV1)
			lineGroup.POST("/v2", line.WebhookV2)
		}

		//discordGroup := webhookGroup.Group("discord")
		//{
		//	discordGroup.POST("/v1", discord.WebhookV1)
		//	discordGroup.POST("/v2", discord.WebhookV2)
		//}
	}
	return router, nil
}
