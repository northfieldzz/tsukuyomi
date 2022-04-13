package line

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"net/http"
	"os"
	"tsukuyomi/log"
)

func GetClient() (*linebot.Client, error) {
	return linebot.New(os.Getenv("LINEBOT_SECRET_KEY"), os.Getenv("LINEBOT_CHANNEL_ACCESS_TOKEN"))
}

func Callback(c *gin.Context) {
	logger := log.GetLogger()
	client, err := GetClient()
	if err != nil {
		message := fmt.Sprintf("Failed initialize line api: %v", err)
		logger.Error(message)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": []string{message},
		})
		return
	}
	events, err := client.ParseRequest(c.Request)
	if err != nil {
		if err != linebot.ErrInvalidSignature {
			message := "Invalid signature to line api"
			logger.Error(message)
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": []string{message},
			})
			return
		} else {
			message := fmt.Sprintf("Failed parse line request: %v", err)
			logger.Error(message)
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors": []string{message},
			})
			return
		}
	}
	err = selector(events)
	err = errors.New("")
	if err != nil {
		logger.Error(fmt.Sprintf("Failed Api Process%v", err))
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
	return
}

func selector(events []*linebot.Event) error {
	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeFollow:
			return eventFollow(event)
		case linebot.EventTypeUnfollow:
			return eventUnFollow(event)
		case linebot.EventTypeMessage:
			return eventMessage(event)
		case linebot.EventTypeJoin:
		case linebot.EventTypeMemberJoined:
		case linebot.EventTypeMemberLeft:
		}
	}
	return nil
}
