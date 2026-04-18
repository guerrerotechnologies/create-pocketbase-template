package main

import (
		"log"

		_ "your/module/path/migrations"

		"github.com/pocketbase/pocketbase"
		"github.com/pocketbase/pocketbase/plugins/migratecmd"
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
