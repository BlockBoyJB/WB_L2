package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"path"
)

type Config struct {
	HTTP HTTP
}

type (
	HTTP struct {
		Port string `yaml:"port" env:"HTTP_PORT"`
	}
)

func NewConfig() (*Config, error) {
	c := &Config{}

	if err := cleanenv.ReadConfig(path.Join("./", "config/config.yaml"), c); err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	if err := cleanenv.ReadEnv(c); err != nil {
		return nil, fmt.Errorf("error reading config env: %w", err)
	}
	return c, nil
}
