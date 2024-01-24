package main

import (
	"app/internal/application"
	"fmt"
	"os"
)

func main() {
	// env
	// ...

	// application
	fmt.Println("Starting application...")
	fmt.Println("Server address: " + os.Getenv("SERVER_ADDR"))
	fmt.Println("Database file: " + os.Getenv("DB_FILE"))
	// - config
	cfg := application.ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}
	app := application.NewApplicationDefault(&cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
