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
	Ok      bool   `json:"ok"`
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
