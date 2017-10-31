package provider

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Auth struct {
	Token string
}

type Query struct {
	Message string `json:"message"`
	OkToday bool   `json:"ok_today"`
}

type StandardQuery interface {
	GetMessage() string
	IsOkToday() bool
	Prepare() Query
}

func unmarshal(res *http.Response, store interface{}, lastError error) error {
	if lastError != nil {
		return lastError
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(body, &store)
}
