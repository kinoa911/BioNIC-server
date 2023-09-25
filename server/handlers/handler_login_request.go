package handlers

import (
	m "PockitGolangBoilerplate/middleware"
	s "PockitGolangBoilerplate/server"
	"fmt"

	"github.com/koesie10/webauthn/webauthn"
	"github.com/labstack/echo/v4"
)

func NewHandlerLoginRequest(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

// Register godoc
// @Summary Register
// @Description New user registration
// @ID user-register
// @Tags User Actions
// @Accept json
// @Produce json
// @Param params body requests.RegisterRequest true "User's email, user's password"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /register [post]
func (registerHandler *RegisterHandler) LoginRequest(c echo.Context) error {
	fmt.Println("Register")

	name := c.Param("name")
	// u, ok := storage.users[name]
	u, ok := registerHandler.server.Sr.GetUsers()[name]

	sess := m.SessionFromContext(c)

	if ok {
		registerHandler.server.WA.StartLogin(c.Request(), c.Response(), u, webauthn.WrapMap(sess.Values))
	} else {
		registerHandler.server.WA.StartLogin(c.Request(), c.Response(), nil, webauthn.WrapMap(sess.Values))
	}
	return nil
}
