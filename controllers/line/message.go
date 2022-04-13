package line

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"google.golang.org/api/youtube/v3"
	"net/url"
	"path"
	"regexp"
	"tsukuyomi/log"
	"tsukuyomi/youtubeApi"
)

func eventMessage(event *linebot.Event) error {
	client, err := GetClient()
	if err != nil {
		return err
	}
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		err = receivedTextMessage(event, message)
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

func receivedTextMessage(event *linebot.Event, message *linebot.TextMessage) error {
	client, err := GetClient()
	if err != nil {
		return err
	}
	var r *regexp.Regexp
	r = regexp.MustCompile(`Tsukuyomi`)
	if r.MatchString(message.Text) {
		err = makeMessageTsukuyomi(client, event)
	}
	r = regexp.MustCompile(`youtube`)
	if r.MatchString(message.Text) {
		err = makeMessageYoutube(client, event.Source)
	}
	if err != nil {
		return err
	}
	return nil
}

func makeMessageTsukuyomi(client *linebot.Client, event *linebot.Event) error {
	logger := log.GetLogger()
	profile, err := client.GetProfile(event.Source.UserID).Do()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed get profile: %v", err))
		return err
	}
	if _, err = client.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(fmt.Sprintf("なんや、%s", profile.DisplayName)),
	).Do(); err != nil {
		logger.Error(fmt.Sprintf("Failed error send message: %v", err))
		return err
	}
	return nil
}

func makeMessageYoutube(client *linebot.Client, source *linebot.EventSource) error {
	video := youtubeApi.FetchLatestVideo()
	uri, err := GenerateURL(video)
	if err != nil {
		return err
	}
	container := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:        linebot.FlexComponentTypeImage,
			URL:         video.Snippet.Thumbnails.Medium.Url,
			Size:        linebot.FlexImageSizeTypeFull,
			AspectRatio: linebot.FlexImageAspectRatioType16to9,
			AspectMode:  linebot.FlexImageAspectModeTypeCover,
			Action: &linebot.URIAction{
				Label: "hero01",
				URI:   uri.String(),
				AltURI: &linebot.URIActionAltURI{
					Desktop: uri.String(),
				},
			},
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Size:   linebot.FlexTextSizeTypeMd,
					Weight: linebot.FlexTextWeightTypeBold,
					Wrap:   true,
					Text:   video.Snippet.Title,
				},
				&linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeBaseline,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type: linebot.FlexComponentTypeText,
							Size: linebot.FlexTextSizeTypeSm,
							Text: BroadcastStatus(video),
						},
					},
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Height: linebot.FlexButtonHeightTypeSm,
					Style:  linebot.FlexButtonStyleTypeLink,
					Action: &linebot.URIAction{
						Label: video.Snippet.ChannelTitle,
						URI:   "https://www.youtube.com/channel/UCS9uQI-jC3DE0L4IpXyvr6w",
						AltURI: &linebot.URIActionAltURI{
							Desktop: "https://www.youtube.com/channel/UCS9uQI-jC3DE0L4IpXyvr6w",
						},
					},
				},
				&linebot.ButtonComponent{
					Height: linebot.FlexButtonHeightTypeSm,
					Style:  linebot.FlexButtonStyleTypeLink,
					Action: &linebot.URIAction{
						Label: "ホロライブ公式",
						URI:   "https://www.youtube.com/channel/UCJFZiqLMntJufDCHc6bQixg/featured",
						AltURI: &linebot.URIActionAltURI{
							Desktop: "https://www.youtube.com/channel/UCJFZiqLMntJufDCHc6bQixg/featured",
						},
					},
				},
			},
		},
	}
	if _, err := client.PushMessage(
		getReplyId(source),
		linebot.NewFlexMessage("tsukuyomiは告げています.", container),
	).Do(); err != nil {
		return err
	}
	return nil
}

func getReplyId(source *linebot.EventSource) string {
	switch source.Type {
	case linebot.EventSourceTypeUser:
		return source.UserID
	case linebot.EventSourceTypeGroup:
		return source.GroupID
	case linebot.EventSourceTypeRoom:
		return source.RoomID
	}
	return ""
}

// GenerateURL 動画URLを生成.
func GenerateURL(sr *youtube.SearchResult) (*url.URL, error) {
	endpoint := "https://www.youtube.com"
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "watch")
	q := u.Query()
	q.Set("v", sr.Id.VideoId)
	u.RawQuery = q.Encode()
	return u, nil
}

// BroadcastStatus 配信ステータスを取得
func BroadcastStatus(v *youtube.SearchResult) string {
	switch v.Snippet.LiveBroadcastContent {
	case "live":
		return "配信中"
	case "upcoming":
		return "配信予定"
	}
	return "配信終了"
}
