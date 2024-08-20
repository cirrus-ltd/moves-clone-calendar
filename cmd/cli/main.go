package main

import (
	"log"
	"os"

	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/config"
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/infrastructure/migrations"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "main",
		Usage: "A new cli application",
		Commands: []*cli.Command{
			{
				Name:  "migration",
				Usage: "DBマイグレーション",
				Action: func(c *cli.Context) error {
					cfg, err := config.LoadConfig()
					if err != nil {
						return err
					}
					migrations.RunMigrations(cfg)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
