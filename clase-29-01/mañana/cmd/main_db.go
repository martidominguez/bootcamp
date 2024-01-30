package main

import (
	"app/internal/application"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// application
	// - config
	cfg := &application.ConfigAppDefault{
		ServerAddress: ":8080",
		ConfigMySQL: &mysql.Config{
			User:   "root",
			Passwd: "",
			Net:    "tcp",
			Addr:   "localhost:3306",
			DBName: "my_db",
		},
	}
	app := application.NewApplicationDefaultDB(cfg)

	// - setup
	db, err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close() // this is here because if not when the function ends, the connection is closed

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
