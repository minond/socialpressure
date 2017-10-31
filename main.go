package main

import (
	"encoding/json"
	"fmt"

	config "github.com/minond/gofigure"
	"github.com/minond/peer-pressure/provider"
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

	todoist := provider.Todoist{provider.Auth{Token: keys.Todoist.Token}}

	dump(todoist.Query(provider.TodoistQuery{
		TaskId:  "2297620443",
		Message: "Have I mediated and/or exercised today?",
	}))
}
