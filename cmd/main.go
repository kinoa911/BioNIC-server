package main

import (
	application "PockitGolangBoilerplate"
	"PockitGolangBoilerplate/config"
)

func main() {
	cfg := config.NewConfig()

	// docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)

	application.Start(cfg)
}
