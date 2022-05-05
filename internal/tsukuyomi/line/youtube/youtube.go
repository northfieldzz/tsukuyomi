package youtube

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"tsukuyomi/pkg/youtubeApi"
)

func MakeMessageYoutube(replyId string, client *linebot.Client) error {
	video := youtubeApi.FetchLatestVideo()
	uri, err := youtubeApi.GenerateURL(video)
	if err != nil {
		return err
	}

	hero := &linebot.ImageComponent{
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
	}

	body := &linebot.BoxComponent{
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
						Text: youtubeApi.BroadcastStatus(video),
					},
				},
			},
		},
	}

	footer := &linebot.BoxComponent{
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
	}

	container := &linebot.BubbleContainer{
		Type:   linebot.FlexContainerTypeBubble,
		Hero:   hero,
		Body:   body,
		Footer: footer,
	}
	if _, err := client.ReplyMessage(
		replyId,
		linebot.NewFlexMessage("tsukuyomiは告げています.", container),
	).Do(); err != nil {
		return err
	}
	return nil
}
