package chesscompubapi

type Streamer struct {
	Username            string `json:"username"`
	Avatar              string `json:"avatar"`
	TwitchURL           string `json:"twitch_url"`
	URL                 string `json:"url"`
	IsLive              bool   `json:"is_live"`
	IsCommunityStreamer bool   `json:"is_community_streamer"`
}

// ListStreamers lists information about Chess.com streamers.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-streamers.
func (c *Client) ListStreamers() ([]Streamer, error) {
	const urlTemplate = "streamers"
	streamers := &struct {
		Streamers []Streamer `json:"streamers"`
	}{}
	err := c.getInto(urlTemplate, streamers)
	return streamers.Streamers, err
}
