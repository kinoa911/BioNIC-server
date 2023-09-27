package handlers

import (
	"PockitGolangBoilerplate/model"
	s "PockitGolangBoilerplate/server"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewHandlerAttestationPOST(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (registerHandler *RegisterHandler) AttestationPOST(c echo.Context) error {
	fmt.Println("AttestationPOST")

	var (
		user    *model.User
		session *model.UserSession
		err     error
	)

	ctx := registerHandler.server
	if session, err = ctx.GetUserSession(); err != nil {
		ctx.Log.Error("failed to retrieve user session", zap.Error(err))

		ctx.ForbiddenJSON(model.NewErrorJSON().WithError(err).WithInfo("Forbidden."))

		return nil
	}

	defer func() {
		session.Webauthn = nil

		if err := ctx.SaveUserSession(session); err != nil {
			ctx.Log.Error("failed to save user session", zap.Error(err))
		}
	}()

	if user, err = ctx.Providers.User.Get(session.Username); err != nil {
		ctx.Log.Error("failed to retrieve user", zap.Error(err))

		ctx.UnauthorizedJSON(model.NewErrorJSON().WithError(err).WithInfo("Unauthorized."))

		return nil
	}

	parsedResponse, err := protocol.ParseCredentialCreationResponseBody(c.Request().Body /*bytes.NewReader(ctx.PostBody())*/)
	if err != nil {
		ctx.Log.Error("failed to parse credential creation response body", model.ProtoErrToFields(err)...)

		ctx.BadRequestJSON(model.NewErrorJSON().WithError(err).WithInfo("Bad Request."))

		return nil
	}

	cred, err := ctx.Providers.Webauthn.CreateCredential(user, *session.Webauthn, parsedResponse)
	if err != nil {
		ctx.Log.Error("failed to create credential", model.ProtoErrToFields(err)...)

		ctx.UnauthorizedJSON(model.NewErrorJSON().WithError(err).WithInfo("Unauthorized."))

		return nil
	}

	user.Credentials = append(user.Credentials, *cred)

	if err = ctx.Providers.User.Set(user); err != nil {
		ctx.Log.Error("failed to save user", zap.String("user_id", user.ID), zap.Error(err))

		ctx.UnauthorizedJSON(model.NewErrorJSON().WithError(err).WithInfo("Unauthorized."))

		return nil
	}

	ctx.CreatedJSON("Done.")

	return nil
}
