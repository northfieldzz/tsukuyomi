package line

import "github.com/line/line-bot-sdk-go/v7/linebot"

func Selector(events []*linebot.Event) {
	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeFollow:
			eventFollow(event)
		case linebot.EventTypeUnfollow:
			eventUnFollow(event)
		case linebot.EventTypeMessage:
			eventMessage(event)
		case linebot.EventTypeJoin:
		case linebot.EventTypeMemberJoined:
		case linebot.EventTypeMemberLeft:
		}
	}
}
