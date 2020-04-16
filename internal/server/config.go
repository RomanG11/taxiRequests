package server

import "github.com/caarlos0/env"

//
// Config contains server configuration
//
type Config struct {
	HttpPort string `env:"HTTP_PORT,required"`
}

//
// ConfigFromEnv func - reads env by struct's fields 'env' annotation
//
func ConfigFromEnv() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
