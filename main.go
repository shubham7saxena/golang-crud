package main

import (
	app "crud/appcontext"
	"crud/config"
	migrationRunner "crud/migration"
	"crud/server"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {
	config.Load()
	app.Initiate()

	clientApp := cli.NewApp()
	clientApp.Name = "crud"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start HTTP server",
			Action: func(c *cli.Context) error {
				router := server.Router()
				log.Fatal(http.ListenAndServe(":8080", router))
				return nil
			},
		},
		{
			Name:        "migrate",
			Description: "Run database migrations",
			Action: func(c *cli.Context) error {
				migrationRunner.Init()
				return migrationRunner.RunDatabaseMigrations()
			},
		},
		{
			Name:        "rollback",
			Description: "Rollback latest database migration",
			Action: func(c *cli.Context) error {
				migrationRunner.Init()
				return migrationRunner.RollbackLatestMigration()
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
