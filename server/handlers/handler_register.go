package handlers

import (
	m "PockitGolangBoilerplate/middleware"
	"PockitGolangBoilerplate/responses"
	s "PockitGolangBoilerplate/server"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterHandler struct {
	server *s.Server
}

func NewHandlerRegisterFinish(server *s.Server) *RegisterHandler {
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
func (registerHandler *RegisterHandler) Register(c echo.Context) error {
	fmt.Println("Register")

	name := c.Param("name")
	u, ok := registerHandler.server.Sr.GetUsers()[name]
	fmt.Println("POST name: , u: ", name, u)
	if !ok {
		fmt.Println("POST !ok")

		return c.NoContent(http.StatusNotFound)
	}

	sess := m.SessionFromContext(c)
	fmt.Println("POST SessionFromContext: ", sess)

	// registerHandler.server.WA.FinishRegistration(c.Request(), c.Response(), u, webauthn.WrapMap(sess.Values))

	return responses.MessageResponse(c, http.StatusAccepted, "Request successfully accepted")
}
