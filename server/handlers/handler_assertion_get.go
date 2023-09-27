package handlers

import (
	"PockitGolangBoilerplate/model"
	s "PockitGolangBoilerplate/server"
	"fmt"
	"strconv"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewHandlerAssertionGET(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (registerHandler *RegisterHandler) AssertionGET(c echo.Context) error {
	fmt.Println("AssertionGET")

	var (
		user    *model.User
		session *model.UserSession
		err     error
	)

	ctx := registerHandler.server
	if session, err = ctx.GetUserSession(); err != nil {
		fmt.Println("AssertionGET GetUserSession err")

		ctx.Log.Error("failed to retrieve user session", zap.Error(err))

		ctx.ForbiddenJSON(model.NewErrorJSON().WithError(err).WithInfo("Forbidden."))

		return nil
	}

	var opts = []webauthn.LoginOption{
		webauthn.WithUserVerification(protocol.VerificationPreferred),
	}

	// discoverable := ctx.RequestCtx.QueryArgs().GetBool(queryArgDiscoverable)
	discoverable, err := strconv.ParseBool(c.QueryParam(queryArgDiscoverable))
	fmt.Println("AssertionGET discoverable: ", c.QueryParam(queryArgDiscoverable), discoverable)
	if !discoverable {
		fmt.Println("AssertionGET !discoverable session.Username: ", session.Username /*ctx.Config.ExternalURL.String()*/, "http://localhost:5000")

		if user, err = ctx.Providers.User.Get(session.Username); err != nil {
			ctx.Log.Error("failed to retrieve user", zap.Error(err))

			ctx.ForbiddenJSON(model.NewErrorJSON().WithError(err).WithInfo("Forbidden."))

			return nil
		}

		credentials := user.WebAuthnCredentialDescriptors()

		opts = append(opts, webauthn.WithAllowedCredentials(credentials), webauthn.WithAppIdExtension( /*ctx.Config.ExternalURL.String()*/ "http://localhost:5000"))
	}

	var (
		assertion *protocol.CredentialAssertion
		data      *webauthn.SessionData
	)

	if discoverable {
		fmt.Println("AssertionGET discoverable")

		ctx.Log.Debug("begin assertion", zap.Bool("discoverable", true))

		if assertion, data, err = ctx.Providers.Webauthn.BeginDiscoverableLogin(opts...); err != nil {
			ctx.Log.Error("error in begin discoverable assertion", append([]zap.Field{zap.Bool("discoverable", true)}, model.ProtoErrToFields(err)...)...)

			ctx.UnauthorizedJSON(model.NewErrorJSON().WithError(err).WithInfo("Unauthorized."))

			return nil
		}
	} else {
		fmt.Println("AssertionGET !discoverable")

		fmt.Println("AssertionGET Debug: ", user, user.ID, opts)
		// ctx.Log.Debug("begin assertion", zap.Bool("discoverable", false), zap.String("user", user.ID))

		fmt.Println("AssertionGET Debug")

		if assertion, data, err = ctx.Providers.Webauthn.BeginLogin(user, opts...); err != nil {
			// ctx.Log.Error("error in begin assertion", append([]zap.Field{zap.Bool("discoverable", false), zap.String("user", user.ID)}, model.ProtoErrToFields(err)...)...)

			ctx.UnauthorizedJSON(model.NewErrorJSON().WithError(err).WithInfo("Unauthorized."))

			return nil
		}
	}

	fmt.Println("AssertionGET SaveUserSession")

	session.Webauthn = data
	if err = ctx.SaveUserSession(session); err != nil {
		fmt.Println("SaveUserSession err")

		ctx.Log.Error("failed to save user session", zap.Error(err))

		ctx.UnauthorizedJSON(model.NewErrorJSON().WithError(err).WithInfo("Unauthorized."))

		return nil
	}

	fmt.Println("assertion", assertion)

	ctx.OKJSON(assertion)

	return nil
}
