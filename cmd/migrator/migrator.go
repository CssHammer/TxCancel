package main

import (
	"github.com/jmoiron/sqlx"

	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	_ "github.com/jackc/pgx/v4/stdlib"

	config "txCancel/config/api"
)

type Migrator struct {
	migrations migrate.MigrationSource
}

func (m *Migrator) MigrateDown(ctx *cli.Context) error {
	return m.Migrate(ctx.GlobalString(ConfigFlag), migrate.Down)
}

func (m *Migrator) MigrateUp(ctx *cli.Context) error {
	return m.Migrate(ctx.GlobalString(ConfigFlag), migrate.Up)
}

func (m *Migrator) Migrate(configPath string, direction migrate.MigrationDirection) error {
	conf, err := config.NewConfig(configPath)
	if err != nil {
		return errors.Wrap(err, "failed to init config")
	}

	db, err := sqlx.Connect("pgx", conf.DbConnection)
	if err != nil {
		return errors.Wrapf(err, "failed to open connection: %s", conf.DbConnection)
	}

	n, err := migrate.Exec(db.DB, "postgres", m.migrations, direction)
	if err != nil {
		return errors.Wrap(err, "migrations failed")

	}

	logrus.Infof("Applied %d migrations!", n)
	return nil
}
