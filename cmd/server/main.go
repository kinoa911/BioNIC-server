package main

import (
	application "PockitGolangBoilerplate"
	"PockitGolangBoilerplate/config"
	"PockitGolangBoilerplate/configuration"
	"PockitGolangBoilerplate/logging"
)

func main() {
	cfg := config.NewConfig()

	// docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)

	config, err := configuration.Load([]string{"config.yaml"}, true, nil)
	if err != nil {
		panic(err)
	}

	logger, err := logging.Configure(&config.Log)
	if err != nil {
		panic(err)
	}

	logger.Info("server is starting")

	application.Start(cfg, config)
}
