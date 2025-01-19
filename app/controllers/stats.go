package controllers

import (
	"net/http"

	"github.com/Valentin-Foucher/social-stats/common"
	"github.com/Valentin-Foucher/social-stats/domain"
	"github.com/Valentin-Foucher/social-stats/domain/providers"
	"github.com/go-chi/chi"
)

type StatsRouter struct {
	clients domain.Clients
}

func GetStatsRouter(clients domain.Clients) func(chi.Router) {
	r := StatsRouter{
		clients: clients,
	}

	return func(router chi.Router) {
		router.Get("/", makeResponse(r.getStats))
	}
}

func (r StatsRouter) getStats(res http.ResponseWriter, req *http.Request) (any, *common.AppError) {
	providerID := chi.URLParam(req, "providerID")

	provider, err := providers.New(providerID, r.clients)
	if err != nil {
		return nil, err
	}

	token, err := getAccessToken(req, providerID)
	if err != nil {
		return nil, err
	}

	return provider.GetStats(token)
}
