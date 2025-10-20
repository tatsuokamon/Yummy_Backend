package ytbdtc

import "fmt"

type SearchResponse struct {
	ETag          string   `json:"etag"`
	Items         []Item   `json:"items"`
	Kind          string   `json:"kind"`
	NextPageToken string   `json:"nextPageToken"`
	PageInfo      PageInfo `json:"pageInfo"`
	RegionCode    string   `json:"regionCode"`
}

type PageInfo struct {
	ResultsPerPage int `json:"resultsPerPage"`
	TotalResults   int `json:"totalResults"`
}

type Item struct {
	ETag string `json:"etag"`
	ID   struct {
		Kind       string `json:"kind"`
		VideoID    string `json:"videoId"`
		ChannelID  string `json:"channelId"`
		PlaylistID string `json:"playlistId"`
	} `json:"id"`
	Snippet Snippet `json:"snippet"`
}

func (i Item) String() string {
	var id string
	if id = i.ID.VideoID; id == "" {
		if id = i.ID.ChannelID; id == "" {
			id = i.ID.PlaylistID
		}
	}
	return fmt.Sprintf(`
	-- kind: %s
	-- Id: %s
	-- Snippet: %s
	`, i.ID.Kind, id, i.Snippet)
}

type Thumbnail struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

type Snippet struct {
	ChannelID            string `json:"channelId"`
	ChannelTitle         string `json:"channelTitle"`
	Description          string `json:"description"`
	LiveBroadcastContent string `json:"liveBroadcastContent"`
	PublishTime          string `json:"publishTime"`
	PublishedAt          string `json:"publishedAt"`
	Thumbnails           struct {
		Default Thumbnail `json:"default"`
		High    Thumbnail `json:"high"`
		Medium  Thumbnail `json:"medium"`
	} `json:"thumbnails"`

	Title string `json:"title"`
}

func (s Snippet) String() string {
	return fmt.Sprintf(`
	-- ChannelID: %s
	-- ChannelTitle: %s
	-- Title: %s
	`, s.ChannelID, s.ChannelTitle, s.Title)
}
