package app

import (
	"github.com/Valentin-Foucher/social-stats/app/controllers"
	"github.com/Valentin-Foucher/social-stats/domain"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func InitializeRouter(clients domain.Clients) chi.Router {
	router := chi.NewRouter()

	setupMiddlewares(router)
	setupRoutes(router, clients)

	return router
}

func setupMiddlewares(router chi.Router) {
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowCredentials: true,
	}))
	router.Use(accessLogs)
}

func setupRoutes(router chi.Router, clients domain.Clients) {
	router.Route("/api/auth", controllers.GetAuthRouter())
	router.Route("/api/{providerID}/stats", controllers.GetStatsRouter(clients))
}
