package routes

import (
	s "PockitGolangBoilerplate/server"
	"net/http"

	"PockitGolangBoilerplate/server/handlers"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SomeHandler(c echo.Context) error {
	fmt.Println("SomeHandler")

	return c.JSON(
		http.StatusOK,
		map[string]any{"message": "Hello, 世界 !"},
	)
}

func ConfigureRoutes(server *s.Server) {
	handlerRegisterRequest := handlers.NewHandlerRegisterRequest(server)
	handlerRegisterFinish := handlers.NewHandlerRegisterFinish(server)

	handlerLoginRequest := handlers.NewHandlerLoginRequest(server)
	handlerLogin := handlers.NewHandlerLogin(server)

	postHandler := handlers.NewPostHandlers(server)
	authHandler := handlers.NewAuthHandler(server)

	// server.Echo.GET("/", m.CORS(efs.Handler()))
	// server.Echo.GET("/{filepath:*}", m.CORS(efs.Handler()))

	// server.Echo.OPTIONS("/debug", m.CORS(handler.Nil))
	// server.Echo.GET("/debug", m.CORS(handler.DebugGET))

	// server.Echo.OPTIONS("/api/info", m.CORS(handler.Nil))
	// server.Echo.OPTIONS("/api/login", m.CORS(handler.Nil))
	// server.Echo.OPTIONS("/api/logout", m.CORS(handler.Nil))
	// server.Echo.OPTIONS("/api/u2f/register", m.CORS(handler.Nil))
	// server.Echo.GET("/api/u2f/register", m.Authenticated(m.CORS(handler.U2FRegisterGET)))
	// server.Echo.POST("/api/u2f/register", m.Authenticated(m.CORS(handler.U2FRegisterPOST)))

	// server.Echo.OPTIONS("/api/webauthn/attestation", m.CORS(handler.Nil))
	// server.Echo.OPTIONS("/api/webauthn/assertion", m.CORS(handler.Nil))

	handlerAssertionGET := handlers.NewHandlerAssertionGET(server)
	handlerAssertionPOST := handlers.NewHandlerAssertionPOST(server)
	server.Echo.GET("/api/webauthn/assertion", handlerAssertionGET.AssertionGET)
	server.Echo.POST("/api/webauthn/assertion", handlerAssertionPOST.AssertionPOST)

	handlerAttestationGET := handlers.NewHandlerAttestationGET(server)
	handlerAttestationPOST := handlers.NewHandlerAttestationPOST(server)
	server.Echo.GET("/api/webauthn/attestation", handlerAttestationGET.AttestationGET)
	server.Echo.POST("/api/webauthn/attestation", handlerAttestationPOST.AttestationPOST)

	handlerLoginPOST := handlers.NewHandlerLoginPOST(server)
	handlerLogoutGET := handlers.NewHandlerLogoutGET(server)
	server.Echo.POST("/api/login", handlerLoginPOST.LoginPOST)
	server.Echo.GET("/api/logout", handlerLogoutGET.LogoutGET)

	handlerInfoGET := handlers.NewHandlerInfoGET(server)
	server.Echo.GET("/api/info", handlerInfoGET.InfoGET)

	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/register-request/:name", handlerRegisterRequest.RegisterRequest)
	server.Echo.POST("/register/:name", handlerRegisterFinish.Register)

	server.Echo.POST("/login-request/:name", handlerLoginRequest.LoginRequest)
	server.Echo.POST("/login/:name", handlerLogin.Login)

	// server.Echo.POST("/login-request/:user", authHandler.Login)
	// server.Echo.POST("/login-finish/:user", authHandler.Login)

	server.Echo.POST("/refresh", authHandler.RefreshToken)

	// fmt.Println(server.Config.Auth.AccessSecret)

	r := server.Echo.Group("")
	// config := middleware.JWTConfig{
	// 	Claims:     &token.JwtCustomClaims{},
	// 	SigningKey: []byte(server.Config.Auth.AccessSecret),
	// }
	// r.Use(middleware.JWTWithConfig(config))

	r.GET("/posts", postHandler.GetPosts)
	r.POST("/posts", postHandler.CreatePost)
	r.DELETE("/posts/:id", postHandler.DeletePost)
	r.PUT("/posts/:id", postHandler.UpdatePost)
}
