package youtubeApi

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"net/url"
	"os"
	"path"
)

func New() (*YoutubeApi, error) {
	key := os.Getenv("GOOGLE_API_KEY")
	service, err := youtube.NewService(context.Background(), option.WithAPIKey(key))
	if err != nil {
		return nil, err
	}
	return &YoutubeApi{
		Service: service,
		Url:     "https://www.youtube.com",
	}, nil
}

type YoutubeApi struct {
	Service *youtube.Service
	Url     string
}

func (g *YoutubeApi) FetchPlaylist(length int64) ([]*youtube.SearchResult, error) {
	query := g.Service.Search.List([]string{"snippet"}).MaxResults(length)
	if res, err := query.Do(); err != nil {
		return nil, err
	} else {
		return res.Items, nil
	}
}

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

// GenerateURL 動画URLを生成.
func GenerateURL(sr *youtube.SearchResult) (*url.URL, error) {
	u, err := url.Parse("https://www.youtube.com")
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
