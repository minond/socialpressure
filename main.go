package main

import (
	"encoding/json"
	"fmt"

	config "github.com/minond/gofigure"
	"github.com/minond/socialpressure/api"
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

	if keys.Todoist.Token == "" {
		panic("I need a Todoist API Token (Settings -> Integrations -> API Token)")
	}

	todoist := api.Todoist{api.Auth{Token: keys.Todoist.Token}}

	dump(todoist.Query(api.TodoistQuery{
		TaskId:  "2297620443",
		Message: "Has Marcos mediated and/or exercised today?",
	}))
}
