package twitter

import "github.com/sivchari/gotwtr"

func GetClient() *gotwtr.Client {
	return gotwtr.New("")
}
