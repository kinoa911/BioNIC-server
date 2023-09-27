package handlers

import (
	s "PockitGolangBoilerplate/server"
	"fmt"
	"net/http"

	"github.com/koesie10/webauthn/webauthn"
	"github.com/labstack/echo/v4"
)

func NewHandlerLogin(server *s.Server) *RegisterHandler {
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
func (registerHandler *RegisterHandler) Login(c echo.Context) error {
	fmt.Println("Register")

	// name := c.Param("name")
	// u, ok := storage.users[name]
	// u, ok := registerHandler.server.Sr.GetUsers()[name]

	// sess := m.SessionFromContext(c)

	var authenticator webauthn.Authenticator
	// if ok {
	// authenticator = registerHandler.server.WA.FinishLogin(c.Request(), c.Response(), u, webauthn.WrapMap(sess.Values))
	// } else {
	// authenticator = registerHandler.server.WA.FinishLogin(c.Request(), c.Response(), nil, webauthn.WrapMap(sess.Values))
	// }
	if authenticator == nil {
		return nil
	}

	authr, ok := authenticator.(*s.Authenticator)
	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, authr.User)
}
