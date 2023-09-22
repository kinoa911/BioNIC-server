package handlers

import (
	"PockitGolangBoilerplate/requests"
	s "PockitGolangBoilerplate/server"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RegisterHandler struct {
	server *s.Server
}

type TemplateRenderer struct {
	templates *template.Template
}

func NewRegisterHandler(server *s.Server) *RegisterHandler {
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

	user, _ := strconv.Atoi(c.Param("user"))

	registerRequest := new(requests.RegisterRequest)

	if err := c.Bind(registerRequest); err != nil {
		return err
	}

	// if err := registerRequest.Validate(); err != nil {
	// 	return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	// }

	// existUser := models.User{}
	// userRepository := repositories.NewUserRepository(registerHandler.server.DB)
	// userRepository.GetUserByEmail(&existUser, registerRequest.Email)

	// if existUser.ID != 0 {
	// 	return responses.ErrorResponse(c, http.StatusBadRequest, "User already exists")
	// }

	// userService := user.NewUserService(registerHandler.server.DB)
	// if err := userService.Register(registerRequest); err != nil {
	// 	return responses.ErrorResponse(c, http.StatusInternalServerError, "Server error")
	// }

	// return responses.MessageResponse(c, http.StatusCreated, "User successfully created")
	c.Echo().Static("/", "server/public")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("server/views/*.html")),
	}

	c.Echo().Renderer = renderer

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"username":    user,
		"displayname": registerRequest.Name,
		"action":      "signup",
	})
}
