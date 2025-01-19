package providers

import (
	"fmt"

	"github.com/Valentin-Foucher/social-stats/common"
	"github.com/Valentin-Foucher/social-stats/domain"
)

const (
	STRAVA = "strava"
)

type Provider interface {
	GetStats(token string) (map[string]any, *common.AppError)
}

func New(providerID string, clients domain.Clients) (Provider, *common.AppError) {
	var provider Provider
	switch providerID {
	case STRAVA:
		if clients != nil {
			provider = &StravaProvider{
				c: clients.GetStrava(),
			}
		}
	default:
		return nil, common.NewDefaultInternalError(fmt.Errorf("invalid provider %s", providerID))
	}
	return provider, nil
}
