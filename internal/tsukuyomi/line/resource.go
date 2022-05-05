package line

import (
	"context"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"regexp"
	"tsukuyomi/ent"
	"tsukuyomi/ent/linesession"
	"tsukuyomi/ent/lineuser"
	"tsukuyomi/internal/tsukuyomi/line/rakuten"
	"tsukuyomi/internal/tsukuyomi/line/youtube"
	"tsukuyomi/pkg/log"
)

type EventsResource struct {
	Events []*linebot.Event
	Client *linebot.Client
}

func (r *EventsResource) Callback() error {
	for _, event := range r.Events {
		er := EventResource{
			event:  event,
			Client: r.Client,
		}
		switch event.Type {
		case linebot.EventTypeFollow:
			return er.follow()
		case linebot.EventTypeUnfollow:
			return er.unFollow()
		case linebot.EventTypeMessage:
			return er.message()
		case linebot.EventTypeJoin:
		case linebot.EventTypeMemberJoined:
		case linebot.EventTypeMemberLeft:
		}
	}
	return nil
}

type EventResource struct {
	event  *linebot.Event
	Client *linebot.Client
}

func (r *EventResource) follow() error {
	logger := log.GetLogger()
	client := ent.GetClient()
	ctx := context.Background()
	switch r.event.Source.Type {
	case linebot.EventSourceTypeUser:
		userId := r.event.Source.UserID
		users, err := client.LineUser.Query().Where(lineuser.ID(userId)).All(ctx)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed read users: %v", err))
			return err
		}
		if len(users) == 0 {
			_, err := client.LineUser.Create().SetID(userId).Save(ctx)
			if err != nil {
				logger.Error(fmt.Sprintf("Failed create user: %v", err))
				return err
			}
		} else {
			_, err := client.LineUser.UpdateOneID(userId).SetIsActive(true).Save(ctx)
			if err != nil {
				logger.Error(fmt.Sprintf("Failed revival user: %v", err))
				return err
			}
		}
	case linebot.EventSourceTypeGroup:
	case linebot.EventSourceTypeRoom:
	}
	return nil
}

func (r *EventResource) unFollow() error {
	logger := log.GetLogger()
	logger.Info("Entry Unfollow event")
	client := ent.GetClient()
	ctx := context.Background()
	switch r.event.Source.Type {
	case linebot.EventSourceTypeUser:
		_, err := client.LineUser.UpdateOneID(r.event.Source.UserID).SetIsActive(false).Save(ctx)
		if err != nil {
			return err
		}
	case linebot.EventSourceTypeGroup:
	case linebot.EventSourceTypeRoom:
	}
	return nil
}

func (r *EventResource) message() error {
	var err error
	if err != nil {
		return err
	}
	switch message := r.event.Message.(type) {
	case *linebot.TextMessage:
		err = r.choiceTextMessage(message)
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
		if _, err = r.Client.ReplyMessage(r.event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
			return err
		}
	}
	return nil
}

func (r *EventResource) choiceTextMessage(message *linebot.TextMessage) (err error) {
	if regexp.MustCompile(`^Tsukuyomi wakeup$`).MatchString(message.Text) {
		r.WakeupCommand()
	} else if regexp.MustCompile(`^Tsukuyomi shutdown$`).MatchString(message.Text) {
		r.ShutdownCommand()
	}

	client := ent.GetClient()
	ctx := context.Background()
	source := r.event.Source
	if sessions, err := client.LineSession.
		Query().
		Where(linesession.TypeEQ(r.sourceType())).
		Where(linesession.UserIDEQ(source.UserID)).
		Where(linesession.GroupIDEQ(source.GroupID)).
		Where(linesession.RoomIDEQ(source.RoomID)).
		All(ctx); err != nil {
		return err
	} else if len(sessions) == 0 {
		return nil
	}

	if regexp.MustCompile(`^Tsukuyomi$`).MatchString(message.Text) {
		r.Called()
	} else if regexp.MustCompile(`^youtube$`).MatchString(message.Text) {
		err = youtube.MakeMessageYoutube(r.event.ReplyToken, r.Client)
	} else if regexp.MustCompile(`^please recipe`).MatchString(message.Text) {
		if err := rakuten.QuickReplyRecipeGenre(r.event.ReplyToken, r.Client); err != nil {
			return err
		}
		//client := ent.GetClient()
		//ctx := context.Background()
		//_, err = client.ReplySession.Create().SetID().Save(ctx)
		return nil

	}
	rakuten.SelectRecipe(r.event.ReplyToken, r.Client, message)

	if err != nil {
		return err
	}
	return nil
}

func (r *EventResource) Called() {
	logger := log.GetLogger()
	profile, err := r.Client.GetProfile(r.event.Source.UserID).Do()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed get profile: %v", err))
	}
	if _, err = r.Client.ReplyMessage(
		r.event.ReplyToken,
		linebot.NewTextMessage(fmt.Sprintf("なんや、%s", profile.DisplayName)),
	).Do(); err != nil {
		logger.Error(fmt.Sprintf("Failed error send message: %v", err))
	}
}

func (r *EventResource) WakeupCommand() {
	logger := log.GetLogger()
	client := ent.GetClient()
	ctx := context.Background()
	source := r.event.Source
	if sessions, err := client.LineSession.
		Query().
		Where(linesession.TypeEQ(r.sourceType())).
		Where(linesession.UserIDEQ(source.UserID)).
		Where(linesession.GroupIDEQ(source.GroupID)).
		Where(linesession.RoomIDEQ(source.RoomID)).
		All(ctx); err != nil {
		panic(err)
	} else {
		var message string
		if len(sessions) == 0 {
			if _, err := client.LineSession.
				Create().
				SetType(r.sourceType()).
				SetUserID(source.UserID).
				SetGroupID(source.GroupID).
				SetRoomID(source.RoomID).
				Save(ctx); err != nil {
				panic(err)
			}
			message = "起動します."
		} else {
			message = "私はもう起きてますよ?"
		}
		if _, err := r.Client.ReplyMessage(
			r.event.ReplyToken,
			linebot.NewTextMessage(message),
		).Do(); err != nil {
			logger.Error(fmt.Sprintf("Failed error send message: %v", err))
		}
	}
}

func (r *EventResource) ShutdownCommand() {
	logger := log.GetLogger()
	client := ent.GetClient()
	ctx := context.Background()
	source := r.event.Source
	if count, err := client.LineSession.
		Delete().
		Where(linesession.TypeEQ(r.sourceType())).
		Where(linesession.UserIDEQ(source.UserID)).
		Where(linesession.GroupIDEQ(source.GroupID)).
		Where(linesession.RoomIDEQ(source.RoomID)).
		Exec(ctx); err != nil {
		panic(err)
	} else {
		var message string
		logger.Info(fmt.Sprintf("Deleted line sessions : %d", count))
		if count == 0 {
			message = "Zzzzz..."
		} else {
			message = "シャットダウンします."
		}
		if _, err := r.Client.ReplyMessage(
			r.event.ReplyToken,
			linebot.NewTextMessage(message),
		).Do(); err != nil {
			logger.Error(fmt.Sprintf("Failed error send message: %v", err))
		}
	}
}

const (
	SourceTypeUser  = 0
	SourceTypeGroup = 1
	SourceTypeRoom  = 2
)

func (r *EventResource) sourceType() int8 {
	var sourceType int8
	switch r.event.Source.Type {
	case linebot.EventSourceTypeUser:
		sourceType = SourceTypeUser
		break
	case linebot.EventSourceTypeGroup:
		sourceType = SourceTypeGroup
		break
	case linebot.EventSourceTypeRoom:
		sourceType = SourceTypeRoom
		break
	}
	return sourceType
}
