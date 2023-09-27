package application

import (
	"PockitGolangBoilerplate/config"
	"PockitGolangBoilerplate/configuration"
	"PockitGolangBoilerplate/server"
	"PockitGolangBoilerplate/server/routes"
	"log"
)

func Start(cfg *config.Config, config *configuration.Config) {
	app := server.NewServer(cfg, config)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
