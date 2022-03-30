package line

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"net/http"
)

func WebhookV1(c *gin.Context) {
	client, err := GetClient()
	if err != nil {
		// TODO:
	}
	events, err := client.ParseRequest(c.Request)
	if err != nil {
		if err != linebot.ErrInvalidSignature {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
		return
	}
	Selector(events)
	c.JSON(http.StatusOK, gin.H{})
}
