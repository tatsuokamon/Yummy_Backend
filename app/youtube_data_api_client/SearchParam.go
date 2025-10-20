package ytbdtc

import (
	"fmt"
	"net/url"
)

const KEY = ""

// search エンドポイントへの構造体ってことにしましょう

const (
	MAXRESULTS = "50"
)

func CreateNewSearchParam(query, typ string) *SearchParam {
	return &SearchParam{Query: query, Type: typ, Part: "snippet", MaxResults: MAXRESULTS}
}

type SearchParam struct {
	Query      string `url:"q"`
	Type       string `url:"type"`
	Part       string `url:"part"`
	Order      string `url:"order"`
	MaxResults string `url:"maxResults"`
	NextPageToken string `url:"nextPageToken"`

	Key string `url:"key"`

	RegionCode string `url:"regionCode"`
}
func (sp *SearchParam)SetNextPageToken(token string){
	sp.NextPageToken = token
}

func (sp *SearchParam) ToURL() string {
	param := url.Values{}

	for field, value := range map[string]string{
		"q":          sp.Query,
		"type":       sp.Type,
		"part":       sp.Part,
		"order":      sp.Order,
		"maxResults": sp.MaxResults,
		"key":        sp.Key,
		"regionCode": sp.RegionCode,
		"nextPageToken": sp.NextPageToken,
	} {
		if value != "" {
			param.Set(field, value)
		}
	}

	return fmt.Sprintf("%s?%s", "https://www.googleapis.com/youtube/v3/search", param.Encode())
}

func (sp *SearchParam) SetKey(key string) {
	sp.Key = key
}
