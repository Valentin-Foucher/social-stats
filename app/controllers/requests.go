package controllers

import (
	"fmt"
	"net/http"

	"github.com/Valentin-Foucher/social-stats/common"
)

type queryParameter struct {
	key   string
	value string
}

func addQueryParams(req *http.Request, parameters ...queryParameter) {
	q := req.URL.Query()
	for _, p := range parameters {
		q.Add(p.key, p.value)
	}

	req.URL.RawQuery = q.Encode()
}

func getAccessToken(req *http.Request, providerID string) (string, *common.AppError) {
	cookie, err := req.Cookie(fmt.Sprintf("%s_token", providerID))
	if err != nil {
		return "", common.NewUnauthorizedError(err)
	}

	return cookie.Value, nil
}
