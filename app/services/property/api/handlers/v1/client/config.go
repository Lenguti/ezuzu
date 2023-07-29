package client

import (
	"fmt"
	"os"
)

// Config - client configuration.
type Config struct {
	Host string
	Port string
}

// NewConfig - initalizes and parses env vars for configuration.
func NewConfig() (Config, error) {
	var (
		cfg Config
		err error
	)
	cfg, err = cfg.parseEnv()
	if err != nil {
		return cfg, fmt.Errorf("new config: unable to parse environment config: %w", err)
	}

	return cfg, nil
}

func (c Config) parseEnv() (Config, error) {
	var (
		host = os.Getenv("PROPERTY_SERVICE_HOST")
		port = os.Getenv("PROPERTY_SERVICE_PORT")
	)

	switch "" {
	case host, port:
		return c, fmt.Errorf("parse env: invalid config provided")
	}

	c.Host = host
	c.Port = port
	return c, nil
}
