package routes

import (
	s "PockitGolangBoilerplate/server"
	"PockitGolangBoilerplate/server/handlers"
	"fmt"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *s.Server) {
	handlerRegisterRequest := handlers.NewHandlerRegisterRequest(server)
	handlerRegisterFinish := handlers.NewHandlerRegisterFinish(server)

	handlerLoginRequest := handlers.NewHandlerLoginRequest(server)
	handlerLogin := handlers.NewHandlerLogin(server)

	postHandler := handlers.NewPostHandlers(server)
	authHandler := handlers.NewAuthHandler(server)

	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/register-request/:name", handlerRegisterRequest.RegisterRequest)
	server.Echo.POST("/register-finish/:name", handlerRegisterFinish.Register)

	server.Echo.POST("/login-request/:name", handlerLoginRequest.LoginRequest)
	server.Echo.POST("/login-finish/:name", handlerLogin.Login)

	server.Echo.POST("/login-request/:user", authHandler.Login)
	server.Echo.POST("/login-finish/:user", authHandler.Login)

	server.Echo.POST("/refresh", authHandler.RefreshToken)

	fmt.Println(server.Config.Auth.AccessSecret)

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
