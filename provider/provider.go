package provider

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Auth struct {
	Token string
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
