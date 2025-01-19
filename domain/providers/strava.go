package providers

import (
	"github.com/Valentin-Foucher/social-stats/common"
	"github.com/Valentin-Foucher/social-stats/domain"
)

type StravaProvider struct {
	c domain.IStravaClient
}

func (p *StravaProvider) GetStats(token string) (map[string]any, *common.AppError) {
	result, err := p.c.GetAuthenticatedAthlete(token)
	if err != nil {
		return nil, common.NewInternalError(err, "error connecting to Strava")
	}

	return result, nil
}
