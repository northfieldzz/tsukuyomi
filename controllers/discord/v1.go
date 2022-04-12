package discord

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebhookV1(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
