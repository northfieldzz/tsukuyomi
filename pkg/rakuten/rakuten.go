package rakuten

import (
	"fmt"
	"net/url"
	"os"
)

func New() *Rakuten {
	return &Rakuten{
		url: &url.URL{
			Scheme: "https",
			Host:   "app.rakuten.co.jp",
		},
	}
}

type Rakuten struct {
	url           *url.URL
	recipesPath   string
	applicationID string
}

func (r *Rakuten) SetQuery(path string, values map[string]string) {
	u := r.url
	u.Path = fmt.Sprintf("services/api/%s", path)
	q := u.Query()
	q.Set("applicationId", os.Getenv("RAKUTEN_APPLICATION_ID"))
	for k, v := range values {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
}
