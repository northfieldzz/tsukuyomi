package line

import (
	"context"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"tsukuyomi/ent"
	"tsukuyomi/ent/lineuser"
)

func eventFollow(event *linebot.Event) {
	client := ent.GetClient()
	ctx := context.Background()
	switch event.Source.Type {
	case linebot.EventSourceTypeUser:
		users, err := client.LineUser.Query().Where(lineuser.IsActive(true)).All(ctx)
		if err != nil {
			// TODO:
		}
		if len(users) == 0 {
			_, err := client.LineUser.
				Create().
				SetID(event.Source.UserID).
				Save(ctx)
			if err != nil {
				// TODO:
			}
		} else {
			_, err := client.LineUser.
				UpdateOneID(event.Source.UserID).
				SetIsActive(true).
				Save(ctx)
			if err != nil {
				// TODO:
			}
		}
	case linebot.EventSourceTypeGroup:
	case linebot.EventSourceTypeRoom:
	}
}

func eventUnFollow(event *linebot.Event) {
	client := ent.GetClient()
	ctx := context.Background()
	switch event.Source.Type {
	case linebot.EventSourceTypeUser:
		_, err := client.LineUser.UpdateOneID(event.Source.UserID).SetIsActive(false).Save(ctx)
		if err != nil {
			// TODO:
		}
	case linebot.EventSourceTypeGroup:
	case linebot.EventSourceTypeRoom:
	}
}
