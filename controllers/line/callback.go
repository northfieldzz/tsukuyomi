package line

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"net/http"
	"tsukuyomi/log"
)

func Callback(c *gin.Context) {
	logger := log.GetLogger()
	client, err := GetClient()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed initialize line api: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	events, err := client.ParseRequest(c.Request)
	if err != nil {
		if err != linebot.ErrInvalidSignature {
			logger.Error("Invalid signature to line api")
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			logger.Error(fmt.Sprintf("Failed parse line request: %v", err))
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
		return
	}
	err = selector(events)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed Api Process%v", err))
		c.JSON(http.StatusOK, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func selector(events []*linebot.Event) error {
	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeFollow:
			return eventFollow(event)
		case linebot.EventTypeUnfollow:
			return eventUnFollow(event)
		case linebot.EventTypeMessage:
			//return eventMessage(event)
		case linebot.EventTypeJoin:
		case linebot.EventTypeMemberJoined:
		case linebot.EventTypeMemberLeft:
		}
	}
	return nil
}
