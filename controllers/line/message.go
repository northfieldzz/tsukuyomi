package line

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func eventMessage(event *linebot.Event) error {
	client, err := GetClient()
	if err != nil {
		return err
	}
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		text := "Hello"
		if _, err = client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text)).Do(); err != nil {
			return err
		}
	case *linebot.StickerMessage:
		replyMessage := fmt.Sprintf(
			"sticker id is %s, stickerResourceType is %s",
			message.StickerID,
			message.StickerResourceType,
		)
		if _, err = client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
			return err
		}
	}
	return nil
}
