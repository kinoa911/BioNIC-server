package server

import (
	"PockitGolangBoilerplate/config"
	"PockitGolangBoilerplate/db"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Echo:   echo.New(),
		DB:     db.Init(cfg),
		Config: cfg,
	}
}

func (server *Server) Start(addr string) error {
	fmt.Println("Start addr: ", addr)

	return server.Echo.Start(":" + addr)
	// return server.Echo.StartAutoTLS(":" + addr)
	// return server.Echo.StartTLS(":"+addr, "server/public/cert/cert.pem", "server/public/cert/key.pem")
}
