package handler

import (
	"PockitGolangBoilerplate/middleware"
)

func DebugGET(ctx *middleware.RequestCtx) {

	ctx.SetBodyString("Hello")
}
