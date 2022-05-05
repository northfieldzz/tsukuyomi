package rakuten

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"tsukuyomi/pkg/rakuten/recipe/category"
	"tsukuyomi/pkg/rakuten/recipe/ranking"
)

const (
	categoriesPath = "Recipe/CategoryList/20170426"
	rankingPath    = "Recipe/CategoryRanking/20170426"
)

type CategoryQuery struct {
	FormatVersion int    `json:"formatVersion,string"`
	CategoryType  string `json:"categoryType,string"`
}

func (r *Rakuten) FetchCategories(categoryType string) (*category.Data, error) {
	var err error
	query := map[string]string{"formatVersion": "2"}
	if categoryType != "" {
		query["categoryType"] = categoryType
	}
	r.SetQuery(categoriesPath, query)
	client := &http.Client{}
	var request *http.Request
	if request, err = http.NewRequest(
		"GET",
		r.url.String(),
		strings.NewReader(url.Values{}.Encode()),
	); err != nil {
		return nil, err
	}
	var response *http.Response
	if response, err = client.Do(request); err != nil {
		return nil, err
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	var body []byte
	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}
	if response.StatusCode == 200 {
		var d category.Data
		if err := json.Unmarshal(body, &d); err != nil {
			return nil, err
		}
		return &d, nil
	} else {
		var d ErrorResponse
		if err := json.Unmarshal(body, &d); err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func (r *Rakuten) FetchRecipeRanking(categoryId int32) (*ranking.Data, error) {
	var err error
	r.SetQuery(rankingPath, map[string]string{
		"formatVersion": "2",
		"categoryId":    strconv.Itoa(int(categoryId)),
	})
	client := &http.Client{}
	var request *http.Request
	if request, err = http.NewRequest(
		"GET",
		r.url.String(),
		strings.NewReader(url.Values{}.Encode()),
	); err != nil {
		return nil, err
	}
	var response *http.Response
	if response, err = client.Do(request); err != nil {
		return nil, err
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	var body []byte
	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		var d ranking.Data
		if err := json.Unmarshal(body, &d); err != nil {
			return nil, err
		}
		return &d, nil
	} else {
		var d ErrorResponse
		if err := json.Unmarshal(body, &d); err != nil {
			return nil, err
		}
		return nil, nil
	}
}
