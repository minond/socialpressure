package main

import (
	"encoding/json"
	"fmt"

	config "github.com/minond/gofigure"
	"github.com/minond/peer-pressure/api"
)

func dump(thing interface{}, err error) {
	if err != nil {
		panic(fmt.Sprintf("Error : %v", err))
		return
	}

	pretty, _ := json.MarshalIndent(thing, "", "  ")
	fmt.Println(string(pretty))
}

func main() {
	var keys struct {
		Todoist struct {
			Token string `yaml:"token"`
		} `yaml:"todoist"`
	}

	config.AddVariants("local")
	config.Load("keys", &keys)

	todoist := api.Todoist{api.Auth{Token: keys.Todoist.Token}}

	dump(todoist.Query(api.TodoistQuery{
		TaskId:  "2297620443",
		Message: "Have I mediated and/or exercised today?",
	}))
}
