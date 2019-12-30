package main

import (
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	ConfigFlag        = "config"
	ConfigPathDefault = "../../config_api.yaml"
)

func main() {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "../../migrations"),
	}

	migrator := Migrator{
		migrations: migrations,
	}

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  ConfigFlag + ", c",
			Value: ConfigPathDefault,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "up",
			Usage:  "migrate up",
			Action: migrator.MigrateUp,
		},
		{
			Name:   "down",
			Usage:  "migrate down",
			Action: migrator.MigrateDown,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.WithError(err).Error("failed to run app")
		return
	}
}
