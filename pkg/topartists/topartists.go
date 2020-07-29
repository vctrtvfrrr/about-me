package topartists

import (
	"fmt"
	"os"

	"github.com/shkh/lastfm-go/lastfm"
)

// TopArtists searches last.fm for my top 10 artists of the past week.
func TopArtists() (result [10]string, err error) {
	APIKey := os.Getenv("LASTFM_API_KEY")
	APISecret := os.Getenv("LASTFM_API_SECRET")
	api := lastfm.New(APIKey, APISecret)

	response, err := api.User.GetTopArtists(lastfm.P{
		"user":   os.Getenv("LASTFM_USER"),
		"period": "7day",
		"limit":  10,
	})
	if err != nil {
		return
	}

	for i, artist := range response.Artists {
		result[i] = fmt.Sprintf(
			"%dº [%s](%s)",
			i+1,
			artist.Name,
			artist.Url,
		)
	}

	return
}
