package line

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"net/http"
	"os"
	"tsukuyomi/pkg/log"
)

func Callback(c *gin.Context) {
	logger := log.GetLogger()
	client, err := linebot.New(
		os.Getenv("LINEBOT_SECRET_KEY"),
		os.Getenv("LINEBOT_CHANNEL_ACCESS_TOKEN"),
	)
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
	er := EventsResource{
		Events: events,
		Client: client,
	}
	err = er.Callback()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed Api Process%v", err))
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
	return
}
