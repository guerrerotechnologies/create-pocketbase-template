package main

import (
	    "log"

		"github.com/pocketbase/pocketbase"
		"github.com/pocketbase/pocketbase/plugins/migratecmd"

		// Run this before adding migrations import below.
		// Also make sure app is not running while running this.
		// go run . migrate collections
		
		// Enable this once you have at least one migration
		// _ "yourpackage/migrations"
		// Example:                                                     
		// _ "github.com/guerrerotechnologies/create-pocketbase/migrations"
)

func main() {
	app := pocketbase.New()

		migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: false,
	})

		if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
