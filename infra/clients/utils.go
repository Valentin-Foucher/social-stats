package clients

import (
	"encoding/json"
	"io"
)

func decodeJson[T any](body io.Reader, data T) (T, error) {
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return zero[T](), err
	}

	return data, nil
}

func zero[T any]() T {
	var result T
	return result
}
