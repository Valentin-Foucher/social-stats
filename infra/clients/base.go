package clients

import (
	"net/http"

	"github.com/Valentin-Foucher/social-stats/configs"
	"github.com/Valentin-Foucher/social-stats/domain"
)

type HttpClients struct {
	Strava *StravaClient
}

func New(conf *configs.Configuration) *HttpClients {
	return &HttpClients{
		Strava: newStravaClient(conf),
	}
}

func (clients *HttpClients) GetStrava() domain.IStravaClient {
	return clients.Strava
}

func newClient(conf *configs.Configuration, baseUrl string) *HttpClient {
	return &HttpClient{
		inner:      http.DefaultClient,
		baseUrl:    baseUrl,
		maxRetries: conf.Http.MaxRetries,
	}
}
