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

type ClubMemberActivity struct {
	Weekly  []ClubMember `json:"weekly"`
	Monthly []ClubMember `json:"monthly"`
	AllTime []ClubMember `json:"all_time"`
}

type ClubMember struct {
	Username string               `json:"username"`
	Joined   UnixSecondsTimestamp `json:"joined"`
}

// GetClub gets additional details about a club.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-club-profile.
func (c *Client) GetClub(id string) (Club, error) {
	const urlTemplate = "club/%s"
	club := Club{}
	err := c.getInto(fmt.Sprintf(urlTemplate, id), &club)
	return club, err
}

// GetClubMemberActivity gets a list of club members (usernames and joined date timestamp), grouped by club-activity frequency.
// Details about the endpoint can be found at https://www.chess.com/news/view/published-data-api#pubapi-endpoint-club-members.
func (c *Client) GetClubMemberActivity(id string) (ClubMemberActivity, error) {
	const urlTemplate = "club/%s/members"
	activity := ClubMemberActivity{}
	err := c.getInto(fmt.Sprintf(urlTemplate, id), &activity)
	return activity, err
}
