package domain

type Clients interface {
	GetStrava() IStravaClient
}

type IStravaClient interface {
	GetAuthenticatedAthlete(token string) (map[string]any, error)
}
