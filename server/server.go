package server

import (
	"PockitGolangBoilerplate/config"
	"PockitGolangBoilerplate/db"
	m "PockitGolangBoilerplate/middleware"
	"fmt"
	"log"
	"net/http"

	// "github.com/go-webauthn/webauthn/webauthn"
	"github.com/gorilla/sessions"
	_ "github.com/koesie10/webauthn/attestation"
	"github.com/koesie10/webauthn/webauthn"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Config *config.Config
	WA     *webauthn.WebAuthn
	Sr     *Storage
}

func NewServer(cfg *config.Config) *Server {
	var storage = &Storage{
		authenticators: make(map[string]*Authenticator),
		users:          make(map[string]*User),
	}

	wc, err := webauthn.New(&webauthn.Config{
		RelyingPartyName:   "webauthn-demo",
		Debug:              true,
		AuthenticatorStore: storage,
	})
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// AllowOrigins: []string{"http://localhost:3000"},
	// AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

			c.Request().Header.Set("Access-Control-Allow-Origin", "*")
			c.Request().Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Request().Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

			fmt.Println("Cors: ", c.Response().Header()["Access-Control-Allow-Origin"])

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(http.StatusNoContent)
			}

			return next(c)
		}
	})

	e.Debug = true
	e.HideBanner = true

	// Add logger and recover middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	// Create the cookie store with an insecure key and use the middleware so sessions are saved
	store := sessions.NewCookieStore([]byte("thisisanunsecurecookiestorepassword"))
	e.Use(session.Middleware(store), m.SessionMiddleware)

	return &Server{
		Echo:   e, //echo.New(),
		DB:     db.Init(cfg),
		Config: cfg,
		WA:     wc,
		Sr:     storage,
	}
}

func (server *Server) Start(addr string) error {
	fmt.Println("Start addr: ", addr)

	return server.Echo.Start(":" + addr)
	// return server.Echo.StartAutoTLS(":" + addr)
	// return server.Echo.StartTLS(":"+addr, "server/public/cert/cert.pem", "server/public/cert/key.pem")
}
