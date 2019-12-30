package main

import (
	"fmt"
	httpStd "net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	config "txCancel/config/api"
	"txCancel/dbqueries"
	"txCancel/http"
	"txCancel/middleware"
)

const (
	ConfigFlag        = "config"
	ConfigPathDefault = "../../config_api.yaml"
)

func main() {
	app := cli.NewApp()
	app.Usage = "api service"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  ConfigFlag + ", c",
			Value: ConfigPathDefault,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "serve",
			Usage:  "starts api",
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

	db.DB.SetMaxOpenConns(conf.MaxOpenConnections)
	db.DB.SetMaxIdleConns(conf.MaxIdleConnections)

	userQ := dbqueries.NewUserQ(db)
	txQ := dbqueries.NewTransactionQ(db)
	counterQ := dbqueries.NewCounterQ(db)
	compositeQ := dbqueries.NewCompositeQ(db, userQ, txQ, counterQ)

	handler := http.NewHandler(userQ, compositeQ)
	middlewares := middleware.NewMiddlewares(conf)
	r := getRouter(handler, middlewares)

	addr := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	logrus.Infof("starting listening %v", addr)
	if err = httpStd.ListenAndServe(addr, r); err != nil {
		return errors.Wrap(err, "failed to run server")
	}

	return nil
}
