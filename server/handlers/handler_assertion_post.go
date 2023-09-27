package handlers

import (
	s "PockitGolangBoilerplate/server"
	"fmt"

	"github.com/labstack/echo/v4"
)

func NewHandlerAssertionPOST(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (registerHandler *RegisterHandler) AssertionPOST(c echo.Context) error {
	fmt.Println("AssertionPOST")

	// var (
	// 	user           *model.User
	// 	credential     *webauthn.Credential
	// 	parsedResponse *protocol.ParsedCredentialAssertionData
	// 	session        *model.UserSession
	// 	err            error
	// )

	// ctx := registerHandler.server
	// if session, err = ctx.GetUserSession(); err != nil {
	// 	ctx.Log.Error("failed to retrieve user session", zap.Error(err))

	// 	ctx.ForbiddenJSON(model.NewErrorJSON().WithError(err).WithInfo("Forbidden."))

	// 	return nil
	// }

	// defer func() {
	// 	session.Webauthn = nil

	// 	if err := ctx.SaveUserSession(session); err != nil {
	// 		ctx.Log.Error("failed to save user session", zap.Error(err))
	// 	}
	// }()

	// discoverable, err := strconv.ParseBool(c.QueryParam(queryArgDiscoverable))

	// if session == nil {

	// }

	// if parsedResponse, err = protocol.ParseCredentialRequestResponseBody(bytes.NewReader(ctx.PostBody())); err != nil {
	// 	ctx.Log.Error("failed to parse credential request response body", model.ProtoErrToFields(err)...)

	// 	ctx.ForbiddenJSON(model.NewErrorJSON().WithError(err).WithInfo("Forbidden."))

	// 	return nil
	// }

	// if discoverable {
	// 	if credential, err = ctx.Providers.Webauthn.ValidateDiscoverableLogin(func(_, userHandle []byte) (_ webauthn.User, err error) {
	// 		if user, err = ctx.Providers.User.Get(string(userHandle)); err != nil {
	// 			return nil, err
	// 		}

	// 		return user, nil
	// 	}, *session.Webauthn, parsedResponse); err != nil {
	// 		ctx.Log.Error("failed to validate discoverable assertion", model.ProtoErrToFields(err)...)

	// 		ctx.UnauthorizedJSON(model.NewErrorJSON().WithError(err).WithInfo("Unauthorized."))

	// 		return nil
	// 	}

	// 	session.Username = user.Name
	// } else {
	// 	if user, err = ctx.Providers.User.Get(session.Username); err != nil {
	// 		ctx.Log.Error("failed to lookup user from session username", zap.String("username", session.Username), zap.Error(err))

	// 		ctx.BadRequestJSON(model.NewErrorJSON().WithErrorStr("failed to lookup user").WithInfo("Bad Request."))

	// 		return nil
	// 	}

	// 	if credential, err = ctx.Providers.Webauthn.ValidateLogin(user, *session.Webauthn, parsedResponse); err != nil {
	// 		ctx.Log.Error("failed to validate assertion", model.ProtoErrToFields(err)...)

	// 		ctx.UnauthorizedJSON(model.NewErrorJSON().WithError(err).WithInfo("Unauthorized."))

	// 		return nil
	// 	}
	// }

	// user.CredentialsSignIn = append(user.CredentialsSignIn, *credential)

	// if err = ctx.Providers.User.Set(user); err != nil {
	// 	ctx.Log.Error("failed to save user", zap.String("user_id", user.ID), zap.Error(err))

	// 	ctx.ForbiddenJSON(model.NewErrorJSON().WithError(err).WithInfo("Forbidden."))

	// 	return nil
	// }

	// ctx.OKJSON(user)

	return nil
}
