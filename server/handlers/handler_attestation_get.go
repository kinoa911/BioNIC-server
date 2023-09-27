package handlers

import (
	"PockitGolangBoilerplate/model"
	"PockitGolangBoilerplate/responses"
	s "PockitGolangBoilerplate/server"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewHandlerAttestationGET(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (registerHandler *RegisterHandler) AttestationGET(c echo.Context) error {
	fmt.Println("AttestationGET")

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

	if user, err = ctx.Providers.User.Get(session.Username); err != nil {
		ctx.Log.Error("failed to retrieve user", zap.Error(err))

		ctx.ForbiddenJSON(model.NewErrorJSON().WithError(err).WithInfo("Forbidden."))

		return nil
	}

	// discoverable := ctx.QueryArgs().GetBool(queryArgDiscoverable)
	discoverable, err := strconv.ParseBool(c.QueryParam(queryArgDiscoverable))
	fmt.Println("AttestationGET discoverable: ", c.QueryParam(queryArgDiscoverable), discoverable)

	var selection protocol.AuthenticatorSelection

	if discoverable {
		selection = ctx.Config.AuthenticatorSelection(protocol.ResidentKeyRequirementRequired)
	} else {
		selection = ctx.Config.AuthenticatorSelection(protocol.ResidentKeyRequirementDiscouraged)
	}

	opts, data, err := ctx.Providers.Webauthn.BeginRegistration(user,
		webauthn.WithAuthenticatorSelection(selection),
		webauthn.WithConveyancePreference(ctx.Config.ConveyancePreference),
		webauthn.WithExclusions(user.WebAuthnCredentialDescriptors()),
		webauthn.WithAppIdExcludeExtension(ctx.Config.ExternalURL.String()),
	)
	fmt.Println("BeginRegistration", data, err)

	if err != nil {
		fmt.Println("BeginRegistration err")
		ctx.Log.Error("failed to generate attestation options", model.ProtoErrToFields(err)...)

		ctx.UnauthorizedJSON(model.NewErrorJSON().WithError(err).WithInfo("Unauthorized."))

		return nil
	}

	session.Webauthn = data

	if err = ctx.SaveUserSession(session); err != nil {
		fmt.Println("SaveUserSession err")

		ctx.Log.Error("failed to save user session", zap.String("username", session.Username), zap.Error(err))

		ctx.UnauthorizedJSON(model.NewErrorJSON().WithError(err).WithInfo("Unauthorized."))

		return nil
	}

	ctx.OKJSON(opts)
	fmt.Println("BeginRegistration return opts: ", opts)

	return responses.Response(c, http.StatusOK, opts)
}
