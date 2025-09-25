package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppPort int
}

func LoadConfig() (*Config, error) {
	portStr := os.Getenv("APP_PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	return &Config{
		AppPort: port,
	}, nil
}
