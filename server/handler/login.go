package handler

import (
	"encoding/json"
	"fmt"

	"PockitGolangBoilerplate/model"

	"PockitGolangBoilerplate/middleware"
)

func LoginPOST(ctx *middleware.RequestCtx) {
	fmt.Println("LoginPOST")

	var form model.LoginForm

	if err := json.Unmarshal(ctx.PostBody(), &form); err != nil {
		ctx.ForbiddenJSON(model.NewErrorJSON().WithErrorStr("invalid credentials").WithInfo("Bad Request."))

		return
	}

	if err := ctx.RegenerateUserSession(); err != nil {
		ctx.ForbiddenJSON(model.NewErrorJSON().WithErrorStr("could not regenerate session").WithInfo("Failed to Regenerate Session."))

		return
	}

	session := &model.UserSession{
		Username: form.Username,
	}

	if err := ctx.SaveUserSession(session); err != nil {
		ctx.ForbiddenJSON(model.NewErrorJSON().WithErrorStr("could not save session").WithInfo("Failed to Save Session."))

		return
	}

	if _, err := ctx.Providers.User.Get(form.Username); err != nil {
		err = ctx.Providers.User.Set(&model.User{
			ID:          form.Username,
			Name:        form.Username,
			DisplayName: form.Username,
		})

		if err != nil {
			ctx.ForbiddenJSON(model.NewErrorJSON().WithErrorStr("could not save new user").WithInfo("Invalid Credentials."))

			return
		}
	}

	ctx.OKJSON(&model.User{
		ID:          "form.Username",
		Name:        "form.Username",
		DisplayName: "form.Username",
	})
}
