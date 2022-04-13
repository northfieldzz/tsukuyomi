package line

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"regexp"
	"tsukuyomi/log"
)

func eventMessage(event *linebot.Event) error {
	client, err := GetClient()
	if err != nil {
		return err
	}
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		err = receivedTextMessage(event.Source, event.ReplyToken, message)
	case *linebot.ImageMessage:
	case *linebot.VideoMessage:
	case *linebot.AudioMessage:
	case *linebot.FileMessage:
	case *linebot.LocationMessage:
		break
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

func receivedTextMessage(source *linebot.EventSource, replyToken string, message *linebot.TextMessage) error {
	logger := log.GetLogger()
	client, err := GetClient()
	if err != nil {
		return err
	}
	text := ""
	r := regexp.MustCompile(`Tsukuyomi`)
	if r.MatchString(message.Text) {
		profile, err := client.GetProfile(source.UserID).Do()
		if err != nil {
			logger.Error(fmt.Sprintf("Failed get profile: %v", err))
			return err
		}
		text = fmt.Sprintf("なんや、%s", profile.DisplayName)
	}
	if _, err = client.ReplyMessage(replyToken, linebot.NewTextMessage(text)).Do(); err != nil {
		logger.Error(fmt.Sprintf("Failed error send message: %v", err))
		return err
	}
	return nil
}
