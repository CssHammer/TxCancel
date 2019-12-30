package api

import (
	"io/ioutil"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Host               string   `yaml:"host"`
	Port               int      `yaml:"port"`
	DbConnection       string   `yaml:"db_connection"`
	MaxOpenConnections int      `yaml:"max_open_connections"`
	MaxIdleConnections int      `yaml:"max_idle_connections"`
	SourceTypes        []string `yaml:"source_types"`
}

func (cfg *Config) Validate() error {
	return validation.ValidateStruct(cfg,
		validation.Field(&cfg.Host, validation.Required),
		validation.Field(&cfg.Port, validation.Required),
		validation.Field(&cfg.DbConnection, validation.Required),
		validation.Field(&cfg.MaxOpenConnections, validation.Required),
		validation.Field(&cfg.MaxIdleConnections, validation.Required),
		validation.Field(&cfg.SourceTypes, validation.Required),
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
