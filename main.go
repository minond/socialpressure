package main

import (
	"fmt"

	config "github.com/minond/gofigure"
	"github.com/minond/socialpressure/api"
)

func dump(q api.Query, err error) {
	if err != nil {
		fmt.Printf("Error (%s): %v\n", q.Message, err)
	} else {
		ans := "No"

		if q.Ok {
			ans = "Yes"
		}

		fmt.Printf("%s %s.\n", q.Message, ans)
	}
}

func query(qch chan<- api.Query, ech chan<- error, todoist api.Todoist, q api.TodoistQuery) {
	go func() {
		q, e := todoist.Query(q)
		qch <- q
		ech <- e
	}()
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
	qch := make(chan api.Query)
	ech := make(chan error)

	query(qch, ech, todoist, api.TodoistQuery{
		TaskID:  "2297620443",
		Message: "Has Marcos mediated and/or exercised today?",
	})

	query(qch, ech, todoist, api.TodoistQuery{
		TaskID:  "2313429809",
		Message: "Has Marcos done his class work today?",
	})

	dump(<-qch, <-ech)
	dump(<-qch, <-ech)
}
