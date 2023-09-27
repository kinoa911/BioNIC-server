package handlers

import (
	s "PockitGolangBoilerplate/server"
	"fmt"

	// "github.com/koesie10/webauthn/webauthn"
	"github.com/labstack/echo/v4"
)

func NewHandlerRegisterRequest(server *s.Server) *RegisterHandler {
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
func (registerHandler *RegisterHandler) RegisterRequest(c echo.Context) error {
	fmt.Println("Register")

	// user := datastore.GetUser() // Find or create the new user
	// options, session, err := w.BeginRegistration(user)
	// handle errors if present
	// store the sessionData values
	// JSONResponse(w, options, http.StatusOK) // return the options generated
	// options.publicKey contain our registration options

	// name := c.Param("name")
	// u, ok := registerHandler.server.Sr.GetUsers()[name]
	// fmt.Println("POST name: , u: ", name, u)
	// if !ok {
	// 	fmt.Println("POST !ok")
	// 	u = &s.User{
	// 		Name:           name,
	// 		Authenticators: make(map[string]*s.Authenticator),
	// 	}
	// 	registerHandler.server.Sr.SetUsers(name, u)
	// 	// storage.users[name] = u
	// }

	// sess := m.SessionFromContext(c)

	// fmt.Println("POST SessionFromContext: ", sess)

	// r := c.Request()
	// fmt.Println("POST c.Request: ", r)
	// rw := c.Response()
	// fmt.Println("POST c.Response: ", rw)
	// session := webauthn.WrapMap(sess.Values)
	// fmt.Println("POST: ", session)
	// registerHandler.server.WA.StartRegistration(c.Request(), c.Response(), u, webauthn.WrapMap(sess.Values))

	// fmt.Println("POST StartRegistration")
	// return nil

	// user, _ := strconv.Atoi(c.Param("user"))

	// registerRequest := new(requests.RegisterRequest)

	// if err := c.Bind(registerRequest); err != nil {
	// 	return err
	// }

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

	// return responses.MessageResponse(c, http.StatusAccepted, "Request successfully accepted")
	return nil
}
