package youtubeApi

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"os"
)

func GetService() (*youtube.Service, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	service, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return service, nil
}

func FetchLatestVideo() *youtube.SearchResult {
	service, err := GetService()
	if err != nil {
		// TODO:
	}
	youtubeSearchList := service.Search.List([]string{"snippet"}).MaxResults(5)
	res, err := youtubeSearchList.Do()
	if err != nil {
		// TODO:
	}

	return res.Items[0]
}
