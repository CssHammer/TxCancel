package cron

import (
	"io/ioutil"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DbConnection string `yaml:"db_connection"`
	IntervalMins int    `yaml:"interval_mins"`
	RecordCount  int    `yaml:"record_count"`
}

func (cfg *Config) Validate() error {
	return validation.ValidateStruct(cfg,
		validation.Field(&cfg.DbConnection, validation.Required),
		validation.Field(&cfg.IntervalMins, validation.Required),
		validation.Field(&cfg.RecordCount, validation.Required),
	)
}

func NewConfig(path string) (*Config, error) {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithMessagef(err, "unable to read config file from path %s", path)
	}

	config := new(Config)
	err = yaml.Unmarshal(rawConfig, config)
	if err != nil {
		return nil, errors.WithMessagef(err, "unable to unmarshal config: %s", rawConfig)

	}

	err = config.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "invalid configuration")
	}

	return config, nil
}
