package line

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"os"
)

func GetClient() (*linebot.Client, error) {
	client, err := linebot.New(os.Getenv("LINEBOT_SECRET_KEY"), os.Getenv("LINEBOT_CHANEL_ACCESS_TOKEN"))
	if err != nil {
		return nil, err
	}
	return client, nil
}
