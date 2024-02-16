package main

import (
	"exercise-03/internal/application"
	"fmt"
)

func main() {
	// env
	// ...

	// app
	// - config
	app := application.NewApplicationDefault(":8080")
	// - tear down
	defer app.TearDown()
	// - set up
	if err := app.SetUp(); err != nil {
		fmt.Println(err)
		return
	}
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}