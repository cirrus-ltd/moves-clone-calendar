package main

import (
	"log"
	"os"

	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/infrastructure/migrations"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "migration",
			Usage:  "DBマイグレーション",
			Action: migrations.RunMigrations,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
