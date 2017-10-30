package main

import (
	"fmt"
	config "github.com/minond/gofigure"
	"github.com/minond/peer-pressure/provider"
)

func main() {
	var keys struct {
		Todoist struct {
			Token string `yaml:"token"`
		} `yaml:"todoist"`
	}

	config.AddVariants("local")
	config.Load("keys", &keys)

	todoist := provider.Todoist{provider.Auth{Token: keys.Todoist.Token}}
	tasks, err := todoist.GetTasks()

	if err != nil {
		panic(fmt.Sprintf("Error making request: %v", err))
	}

	fmt.Println(tasks)
}
