package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Valentin-Foucher/social-stats/common"
)

type errorPayload struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

func httpStatusCode(e *common.AppError) int {
	switch e.Kind() {
	case common.NotFound:
		return http.StatusNotFound
	case common.InternalError:
		return http.StatusInternalServerError
	case common.InvalidParameters:
		return http.StatusBadRequest
	case common.Unauthorized:
		return http.StatusUnauthorized
	}
	panic("unreachable")
}

func makeResponse(f func(w http.ResponseWriter, r *http.Request) (any, *common.AppError)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var status int

		data, err := f(w, r)

		if data != nil {
			status = http.StatusOK
		} else if err == nil {
			status = http.StatusNoContent
		} else {
			status = httpStatusCode(err)
			data = errorPayload{Error: err.Error(), Description: err.Description()}
		}

		if data != nil {
			b, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				status = http.StatusInternalServerError
			}

			w.Write(b)
		}

		w.WriteHeader(status)
	})
}
