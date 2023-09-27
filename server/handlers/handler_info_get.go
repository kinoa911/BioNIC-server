package handlers

import (
	"PockitGolangBoilerplate/model"
	"PockitGolangBoilerplate/responses"
	s "PockitGolangBoilerplate/server"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewHandlerInfoGET(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (registerHandler *RegisterHandler) InfoGET(c echo.Context) error {
	fmt.Println("InfoGET")

	ctx := registerHandler.server
	session, err := ctx.GetUserSession()
	if err != nil {
		ctx.BadRequestJSON(model.NewErrorJSON().WithError(err).WithInfo("Could not get session."))
	}

	response := model.Info{
		Username: session.Username,
	}

	ctx.OKJSON(response)

	return responses.MessageResponse(c, http.StatusOK, "Request successfully accepted")
}
