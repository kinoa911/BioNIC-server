package handlers

import (
	"PockitGolangBoilerplate/model"
	"PockitGolangBoilerplate/responses"
	s "PockitGolangBoilerplate/server"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewHandlerLogoutGET(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (registerHandler *RegisterHandler) LogoutGET(c echo.Context) error {
	fmt.Println("LogoutGET")

	ctx := registerHandler.server
	if err := ctx.DestroyUserSession(); err != nil {
		ctx.ForbiddenJSON(model.NewErrorJSON().WithErrorStr("could not regenerate session").WithInfo("Invalid Credentials."))

		return nil
	}

	ctx.OKJSON(nil)

	return responses.MessageResponse(c, http.StatusOK, "Request successfully accepted")
}
