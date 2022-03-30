package line

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func eventMessage(event *linebot.Event) {
	client, err := GetClient()
	if err != nil {
		// TODO:
	}
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		text := ""
		if _, err = client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text)).Do(); err != nil {
			// TODO:
		}
	case *linebot.StickerMessage:
		replyMessage := fmt.Sprintf(
			"sticker id is %s, stickerResourceType is %s",
			message.StickerID,
			message.StickerResourceType,
		)
		if _, err = client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
			// TODO:
		}
	}
}
