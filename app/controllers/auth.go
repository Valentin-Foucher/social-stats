package controllers

import (
	"net/http"
	"time"

	"github.com/Valentin-Foucher/social-stats/common"
	"github.com/Valentin-Foucher/social-stats/domain/validators"
	"github.com/go-chi/chi"
	"github.com/markbates/goth/gothic"
)

func GetAuthRouter() func(chi.Router) {
	return func(router chi.Router) {
		router.Get("/{providerID}", makeResponse(login))
		router.Get(("/{providerID}/callback"), makeResponse(callback))
	}
}

type AccessTokenResponse struct {
	Token string `json:"token"`
}

func login(res http.ResponseWriter, req *http.Request) (any, *common.AppError) {
	providerID := chi.URLParam(req, "providerID")
	if err := validators.ValidateProvider(providerID); err != nil {
		return nil, err
	}

	addQueryParams(
		req,
		queryParameter{key: "provider", value: providerID},
	)

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		gothic.BeginAuthHandler(res, req)
		return nil, nil
	}

	return AccessTokenResponse{Token: user.AccessToken}, nil
}

func callback(res http.ResponseWriter, req *http.Request) (any, *common.AppError) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		return nil, common.NewUnauthorizedError(err)
	}

	http.SetCookie(res, &http.Cookie{
		Value:   user.AccessToken,
		Name:    "strava_token",
		Path:    "/",
		Expires: time.Now().Add(6 * time.Hour),
	})

	return nil, nil
}
