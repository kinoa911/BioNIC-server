package handler

import (
	"PockitGolangBoilerplate/model"

	"PockitGolangBoilerplate/middleware"
)

func LogoutGET(ctx *middleware.RequestCtx) {
	if err := ctx.DestroyUserSession(); err != nil {
		ctx.ForbiddenJSON(model.NewErrorJSON().WithErrorStr("could not regenerate session").WithInfo("Invalid Credentials."))

		return
	}

	ctx.OKJSON(nil)
}
