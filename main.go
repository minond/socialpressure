package main

import (
	"encoding/json"
	"fmt"
	config "github.com/minond/gofigure"
	"github.com/minond/peer-pressure/provider"
	"time"
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

	task, err := todoist.GetTask("2297620443")

	if err != nil {
		panic(fmt.Sprintf("Error getting task: %v", err))
	}

	mediatedToday := task.Due.Date.After(time.Now())
	fmt.Printf("Have I mediated and/or exercised today? %v\n", mediatedToday)
}
