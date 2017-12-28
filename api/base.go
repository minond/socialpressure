package api

import (
	"encoding/json"
	"net/http"
)

type Auth struct {
	Token string
}

type Query struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func unmarshal(res *http.Response, store interface{}, lastError error) error {
	if lastError != nil {
		return lastError
	}

	return json.NewDecoder(res.Body).Decode(&store)
}
