package line

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"os"
)

func GetClient() (*linebot.Client, error) {
	return linebot.New(os.Getenv("LINEBOT_SECRET_KEY"), os.Getenv("LINEBOT_CHANNEL_ACCESS_TOKEN"))
}
