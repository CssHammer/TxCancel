package main

import (
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	config "txCancel/config/cron"
	"txCancel/dbqueries"
)

const (
	ConfigFlag        = "config"
	ConfigPathDefault = "../../config_cron.yaml"
)

func main() {
	app := cli.NewApp()
	app.Usage = "cron service"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  ConfigFlag + ", c",
			Value: ConfigPathDefault,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "serve",
			Usage:  "starts cron",
			Action: serve,
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Error("failed to run app")
	}
}

func serve(c *cli.Context) error {
	conf, err := config.NewConfig(c.GlobalString(ConfigFlag))
	if err != nil {
		return errors.Wrap(err, "failed to init config")
	}

	db, err := sqlx.Connect("pgx", conf.DbConnection)
	if err != nil {
		return errors.Wrap(err, "failed to init db")
	}

	userQ := dbqueries.NewUserQ(db)
	txQ := dbqueries.NewTransactionQ(db)
	counterQ := dbqueries.NewCounterQ(db)
	compositeQ := dbqueries.NewCompositeQ(db, userQ, txQ, counterQ)

	for {
		logrus.Info("starting iteration")
		err = compositeQ.CancelTransactions(conf.RecordCount)
		if err != nil {
			logrus.WithError(err).Error("failed to cancel transactions")
		}
		logrus.Info("finished iteration")
		time.Sleep(time.Duration(conf.IntervalMins) * time.Minute)
	}
}
