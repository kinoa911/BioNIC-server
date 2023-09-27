package handlers

import (
	"PockitGolangBoilerplate/model"
	"PockitGolangBoilerplate/responses"
	s "PockitGolangBoilerplate/server"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewHandlerLoginPOST(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (registerHandler *RegisterHandler) LoginPOST(c echo.Context) error {
	var form model.LoginForm

	ctx := registerHandler.server
	fmt.Println("LoginPOST: ", ctx.RequestCtx)

	d := json.NewDecoder(c.Request().Body)
	d.DisallowUnknownFields()
	// d.Decode(&form)
	// fmt.Println("LoginPOST PostBody: ", form)

	if err := d.Decode(&form); /*json.Unmarshal(ctx.RequestCtx.PostBody(), &form)*/ err != nil {
		fmt.Println("LoginPOST Bad Request")

		ctx.ForbiddenJSON(model.NewErrorJSON().WithErrorStr("invalid credentials").WithInfo("Bad Request."))

		return nil
	}

	fmt.Println("PostBody")

	if err := ctx.RegenerateUserSession(); err != nil {
		ctx.ForbiddenJSON(model.NewErrorJSON().WithErrorStr("could not regenerate session").WithInfo("Failed to Regenerate Session."))

		return nil
	}

	fmt.Println("RegenerateUserSession")

	session := &model.UserSession{
		Username: form.Username,
	}

	if err := ctx.SaveUserSession(session); err != nil {
		ctx.ForbiddenJSON(model.NewErrorJSON().WithErrorStr("could not save session").WithInfo("Failed to Save Session."))

		return nil
	}

	fmt.Println("SaveUserSession")

	if _, err := ctx.Providers.User.Get(form.Username); err != nil {
		err = ctx.Providers.User.Set(&model.User{
			ID:          form.Username,
			Name:        form.Username,
			DisplayName: form.Username,
		})

		if err != nil {
			ctx.ForbiddenJSON(model.NewErrorJSON().WithErrorStr("could not save new user").WithInfo("Invalid Credentials."))

			return nil
		}
	}

	fmt.Println("Get")

	ctx.OKJSON(&model.User{
		ID:          "form.Username",
		Name:        "form.Username",
		DisplayName: "form.Username",
	})

	return responses.MessageResponse(c, http.StatusOK, "Request successfully accepted")
}
