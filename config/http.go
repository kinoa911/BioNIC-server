package config

import "os"

type HTTPConfig struct {
	Host       string
	Port       string
	ExposePort string
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Host:       os.Getenv("HOST"),
		Port:       "8000", //os.Getenv("PORT"),
		ExposePort: os.Getenv("EXPOSE_PORT"),
	}
}
