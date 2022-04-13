package line

import (
	"context"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"tsukuyomi/ent"
	"tsukuyomi/ent/lineuser"
	"tsukuyomi/log"
)

func eventFollow(event *linebot.Event) error {
	logger := log.GetLogger()
	logger.Info("Entry Follow event")
	client := ent.GetClient()
	ctx := context.Background()
	switch event.Source.Type {
	case linebot.EventSourceTypeUser:
		logger.Info("Entry Event source type")
		users, err := client.LineUser.Query().Where(lineuser.IsActive(true)).All(ctx)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed read users: %v", err))
			return err
		}
		if len(users) == 0 {
			logger.Info("Entry Event source type   Create")
			_, err := client.LineUser.
				Create().
				SetID(event.Source.UserID).
				Save(ctx)
			if err != nil {
				logger.Error(fmt.Sprintf("Failed create user: %v", err))
				return err
			}
		} else {
			logger.Info("Entry Event source type recreate ")
			_, err := client.LineUser.
				UpdateOneID(event.Source.UserID).
				SetIsActive(true).
				Save(ctx)
			if err != nil {
				logger.Error(fmt.Sprintf("Failed revival user: %v", err))
				return err
			}
		}
	case linebot.EventSourceTypeGroup:
	case linebot.EventSourceTypeRoom:
		logger.Info("Entry Event source group or room")
	}
	return nil
}

func eventUnFollow(event *linebot.Event) error {
	logger := log.GetLogger()
	logger.Info("Entry Unfollow event")
	client := ent.GetClient()
	ctx := context.Background()
	switch event.Source.Type {
	case linebot.EventSourceTypeUser:
		_, err := client.LineUser.UpdateOneID(event.Source.UserID).SetIsActive(false).Save(ctx)
		if err != nil {
			return err
		}
	case linebot.EventSourceTypeGroup:
	case linebot.EventSourceTypeRoom:
	}
	return nil
}
