package clients

import (
	"net/http"

	"github.com/Valentin-Foucher/social-stats/configs"
)

type StravaClient struct {
	*HttpClient
}

func newStravaClient(conf *configs.Configuration) *StravaClient {
	return &StravaClient{newClient(conf, conf.Strava.BaseUrl)}
}

func (c *StravaClient) GetAuthenticatedAthlete(token string) (map[string]any, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	response, err := c.get("/athlete", []int{http.StatusOK}, headers, nil)
	if err != nil {
		return nil, err
	}

	data, err := decodeJson(response.Body, map[string]any{})
	if err != nil {
		return nil, err
	}

	return data, nil
}
