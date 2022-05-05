package rakuten

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"math/rand"
	"regexp"
	"tsukuyomi/pkg/rakuten"
)

const (
	RecipeGenreRandom  int32 = 0
	RecipeGenreMeatId  int32 = 10
	RecipeGenreFishId  int32 = 11
	RecipeGenrePastaId int32 = 15
	RecipeGenreSoupId  int32 = 17
	RecipeGenreSaladId int32 = 18
)

const (
	WantRandomRecipeMessage = "適当なレシピ"
	WantMeatRecipeMessage   = "肉のレシピ"
	WantFishRecipeMessage   = "魚のレシピ"
	WantPastaRecipeMessage  = "パスタのレシピ"
	WantSoupRecipeMessage   = "スープのレシピ"
	WantSaladRecipeMessage  = "サラダのレシピ"
)

func SelectRecipe(replyId string, client *linebot.Client, message *linebot.TextMessage) {
	var categoryId int32
	switch {
	case regexp.MustCompile(WantRandomRecipeMessage).MatchString(message.Text):
		categoryId = RecipeGenreRandom
	case regexp.MustCompile(WantMeatRecipeMessage).MatchString(message.Text):
		categoryId = RecipeGenreMeatId
	case regexp.MustCompile(WantFishRecipeMessage).MatchString(message.Text):
		categoryId = RecipeGenreFishId
	case regexp.MustCompile(WantPastaRecipeMessage).MatchString(message.Text):
		categoryId = RecipeGenrePastaId
	case regexp.MustCompile(WantSoupRecipeMessage).MatchString(message.Text):
		categoryId = RecipeGenreSoupId
	case regexp.MustCompile(WantSaladRecipeMessage).MatchString(message.Text):
		categoryId = RecipeGenreSaladId
	}
	if err := SendRecipes(replyId, client, categoryId); err != nil {
		return
	}
}

func QuickReplyRecipeGenre(replyId string, client *linebot.Client) error {
	text := linebot.NewTextMessage("何系？")
	sending := text.WithQuickReplies(
		&linebot.QuickReplyItems{
			Items: []*linebot.QuickReplyButton{
				{Action: linebot.NewMessageAction("ランダム", WantRandomRecipeMessage)},
				{Action: linebot.NewMessageAction("肉", WantMeatRecipeMessage)},
				{Action: linebot.NewMessageAction("魚", WantFishRecipeMessage)},
				{Action: linebot.NewMessageAction("パスタ", WantPastaRecipeMessage)},
				{Action: linebot.NewMessageAction("スープ", WantSoupRecipeMessage)},
				{Action: linebot.NewMessageAction("サラダ", WantSaladRecipeMessage)},
			},
		},
	)

	if _, err := client.ReplyMessage(replyId, sending).Do(); err != nil {
		return err
	}
	return nil
}

func SendRecipes(replyId string, client *linebot.Client, categoryId int32) error {
	api := rakuten.New()
	if categoryId == RecipeGenreRandom {
		categories, err := api.FetchCategories("large")
		if err != nil {
			return err
		}
		var c []int32
		for _, v := range categories.Result.Large {
			c = append(c, v.Id)
		}

		categoryId = c[rand.Intn(len(c))]

		print(fmt.Sprintf("%v"), categories)
	}

	recipes, err := api.FetchRecipeRanking(categoryId)
	if err != nil {
		return err
	}

	var contents []*linebot.BubbleContainer
	currentCount := 0
	maxCount := 10
	for _, r := range recipes.Results {
		hero := &linebot.ImageComponent{
			Type:        linebot.FlexComponentTypeImage,
			URL:         r.FoodImageUrl,
			AspectRatio: linebot.FlexImageAspectRatioType1to1,
			AspectMode:  linebot.FlexImageAspectModeTypeCover,
			Action: &linebot.URIAction{
				Label: "Recipe Site",
				URI:   r.Url,
				AltURI: &linebot.URIActionAltURI{
					Desktop: r.Url,
				},
			},
		}
		body := &linebot.BoxComponent{
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Wrap: true,
					Text: r.Title,
				},
			},
		}
		contents = append(
			contents,
			&linebot.BubbleContainer{
				Size: linebot.FlexBubbleSizeTypeNano,
				Hero: hero,
				Body: body,
			},
		)
		currentCount++
		if currentCount > maxCount {
			break
		}
	}
	container := &linebot.CarouselContainer{Contents: contents}

	if _, err := client.ReplyMessage(
		replyId,
		linebot.NewFlexMessage("tsukuyomiは告げています.", container),
	).Do(); err != nil {
		return err
	}
	return nil
}
