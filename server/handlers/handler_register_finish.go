package handlers

import (
	m "PockitGolangBoilerplate/middleware"
	"PockitGolangBoilerplate/responses"
	s "PockitGolangBoilerplate/server"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/koesie10/webauthn/webauthn"
	"github.com/labstack/echo/v4"
)

type RegisterHandler struct {
	server *s.Server
}

type TemplateRenderer struct {
	templates *template.Template
}

func NewHandlerRegisterFinish(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
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

	registerHandler.server.WA.FinishRegistration(c.Request(), c.Response(), u, webauthn.WrapMap(sess.Values))

	return responses.MessageResponse(c, http.StatusAccepted, "Request successfully accepted")
}
