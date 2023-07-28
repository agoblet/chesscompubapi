package chesscompubapi

import "fmt"

type Club struct {
	ID                 StringFromPathSuffix   `json:"@id"`
	Name               string                 `json:"name"`
	URL                string                 `json:"url"`
	ClubID             int                    `json:"club_id"`
	Icon               string                 `json:"icon"`
	Visibility         string                 `json:"visibility"`
	JoinRequest        string                 `json:"join_request"`
	Description        string                 `json:"description"`
	Country            StringFromPathSuffix   `json:"country"`
	AverageDailyRating int                    `json:"average_daily_rating"`
	MembersCount       int                    `json:"members_count"`
	Created            UnixSecondsTimestamp   `json:"created"`
	LastActivity       UnixSecondsTimestamp   `json:"last_activity"`
	Admins             []StringFromPathSuffix `json:"admin"`
}

// GetClub gets additional details about a club.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-club-profile.
func (c *Client) GetClub(id string) (Club, error) {
	const urlTemplate = "club/%s"
	club := Club{}
	err := c.getInto(fmt.Sprintf(urlTemplate, id), &club)
	return club, err
}
