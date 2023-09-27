package server

import (
	"PockitGolangBoilerplate/config"
	"PockitGolangBoilerplate/configuration"
	"PockitGolangBoilerplate/db"
	m "PockitGolangBoilerplate/middleware"
	"PockitGolangBoilerplate/model"
	"PockitGolangBoilerplate/server/handler"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"

	// "github.com/gorilla/sessions"
	"github.com/fasthttp/session/v2"
	_ "github.com/koesie10/webauthn/attestation"
	"github.com/valyala/fasthttp"

	"go.uber.org/zap"

	// "github.com/koesie10/webauthn/webauthn"

	"github.com/jinzhu/gorm"
	// "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo *echo.Echo
	DB   *gorm.DB
	Cfg  *config.Config
	WA   *webauthn.WebAuthn
	Sr   *Storage

	RequestCtx *fasthttp.RequestCtx
	Log        *zap.Logger
	Providers  *m.Providers
	Config     *configuration.Config
}

var (
	w   *webauthn.WebAuthn
	err error
)

type UserProvider interface {
	Get(name string) (user *model.User, err error)
	Set(user *model.User) (err error)
}

type Providers struct {
	Webauthn *webauthn.WebAuthn
	Session  *session.Session

	User UserProvider
}

func (ctx *Server) DestroyUserSession() (err error) {
	if err = ctx.RegenerateUserSession(); err != nil {
		return err
	}

	store := session.NewStore()

	return ctx.Providers.Session.Save(ctx.RequestCtx, store)
}

func (ctx *Server) RegenerateUserSession() (err error) {
	return ctx.Providers.Session.Regenerate(ctx.RequestCtx)
}

func (ctx *Server) GetUserSession() (session *model.UserSession, err error) {
	fmt.Println("GetUserSession Providers: ", ctx, ctx.Providers)
	fmt.Println("GetUserSession Session: ", ctx.Providers.Session)
	fmt.Println("GetUserSession RequestCtx: ", ctx.RequestCtx)
	store, err := ctx.Providers.Session.Get(ctx.RequestCtx)
	if err != nil {
		fmt.Println("ctx.Providers.Session.Get err: ", ctx.RequestCtx)

		return nil, err
	}

	session = &model.UserSession{}

	sessionBytes, ok := store.Get("user").([]byte)
	if !ok {
		fmt.Println("store.Get err: ", ctx.RequestCtx)

		return session, nil
	}

	fmt.Println("json.Unmarshal: ", sessionBytes)
	if err = json.Unmarshal(sessionBytes, session); err != nil {
		fmt.Println("json.Unmarshal err: ")
		return nil, err
	}

	return session, nil
}

func (ctx *Server) SaveUserSession(session *model.UserSession) (err error) {
	fmt.Println("SaveUserSession", session)

	store, err := ctx.Providers.Session.Get(ctx.RequestCtx)
	if err != nil {
		return err
	}

	fmt.Println("SaveUserSession Get", store)
	sessionJSON, err := json.Marshal(*session)
	if err != nil {
		return err
	}

	fmt.Println("SaveUserSession Marshal", sessionJSON)
	store.Set("user", sessionJSON)

	return ctx.Providers.Session.Save(ctx.RequestCtx, store)
}

func (ctx *Server) CreateKO(message interface{}) (ko model.MessageResponse) {
	ko = model.MessageResponse{
		Status: "KO",
	}

	switch m := message.(type) {
	case *model.ErrorJSON:
		ko.Message = m.Info()
	case error:
		ko.Message = m.Error()
	case string:
		ko.Message = m
	}

	return ko
}

func (ctx *Server) CreatedJSON(message string) {
	response := model.MessageResponse{
		Status:  "OK",
		Message: message,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		ctx.Log.Error("failed to marshal JSON 201 Created response", zap.Error(err))

		return
	}

	ctx.RequestCtx.SetStatusCode(fasthttp.StatusCreated)
	ctx.RequestCtx.SetBody(responseJSON)
}

func (ctx *Server) OKJSON(data interface{}) {
	response := model.DataResponse{
		Status: "OK",
		Data:   data,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		ctx.Log.Error("failed to marshal JSON 200 OK response", zap.Error(err))

		return
	}

	ctx.RequestCtx.SetStatusCode(fasthttp.StatusOK)
	ctx.RequestCtx.SetBody(responseJSON)
}

func (ctx *Server) ErrorJSON(err error, status int) {
	response := ctx.CreateKO(err)

	responseJSON, err := json.Marshal(response)
	if err != nil {
		ctx.Log.Error(fmt.Sprintf("failed to marshal JSON %d %s response", status, fasthttp.StatusMessage(status)), zap.Error(err))

		return
	}

	ctx.RequestCtx.SetStatusCode(status)
	ctx.RequestCtx.SetBody(responseJSON)
}

func (ctx *Server) BadRequestJSON(err error) {
	ctx.ErrorJSON(err, fasthttp.StatusBadRequest)
}

func (ctx *Server) ForbiddenJSON(err error) {
	ctx.ErrorJSON(err, fasthttp.StatusForbidden)
}

func (ctx *Server) UnauthorizedJSON(err error) {
	ctx.ErrorJSON(err, fasthttp.StatusUnauthorized)
}

func NewServer(cfg *config.Config, config *configuration.Config) *Server {
	var providers *m.Providers

	if providers, err = m.NewProviders(config); err != nil {
		return nil
	}

	efs := handler.NewEmbeddedFS(handler.EmbeddedFSConfig{
		Prefix:     "public_html",
		IndexFiles: []string{"index.html"},
		TemplatedFiles: map[string]handler.TemplatedEmbeddedFSFileConfig{
			"index.html": {
				Data: struct{ ExternalURL string }{config.ExternalURL.String()},
			},
		},
	}, assets)

	if err = efs.Load(); err != nil {
		return nil
	}

	var storage = &Storage{
		authenticators: make(map[string]*Authenticator),
		users:          make(map[string]*User),
	}

	// wc, err := webauthn.New(&webauthn.Config{
	// 	RelyingPartyName:   "webauthn-demo",
	// 	Debug:              true,
	// 	AuthenticatorStore: storage,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// wconfig := &webauthn.Config{
	// 	RPDisplayName: "Go Webauthn",                               // Display Name for your site
	// 	RPID:          "go-webauthn.local",                         // Generally the FQDN for your site
	// 	RPOrigins:     []string{"https://login.go-webauthn.local"}, // The origin URLs allowed for WebAuthn requests
	// }

	// if w, err = webauthn.New(wconfig); err != nil {
	// 	fmt.Println(err)
	// }

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, HEAD, OPTIONS")

			c.Request().Header.Set("Access-Control-Allow-Origin", "*")
			c.Request().Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Request().Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, HEAD, OPTIONS")

			fmt.Println("Cors: ", c.Response().Header()["Access-Control-Allow-Origin"])

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(http.StatusNoContent)
			}

			return next(c)
		}
	})

	e.Debug = true
	e.HideBanner = true

	// Add logger and recover middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	// Create the cookie store with an insecure key and use the middleware so sessions are saved
	// store := sessions.NewCookieStore([]byte("thisisanunsecurecookiestorepassword"))
	// e.Use(session.Middleware(store), m.SessionMiddleware)

	fmt.Println("NewServer: ", config.DisplayName, config.ExternalURL, config.ListenAddress)
	return &Server{
		Echo:       e, //echo.New(),
		DB:         db.Init(cfg),
		Cfg:        cfg,
		WA:         w, //wc,
		Sr:         storage,
		Providers:  providers,
		RequestCtx: new(fasthttp.RequestCtx),
		// Log * zap.Logger
		Config: config,
	}
}

func (server *Server) Start(addr string) error {
	fmt.Println("Start addr: ", addr)

	return server.Echo.Start(":" + addr)
	// return server.Echo.StartAutoTLS(":" + addr)
	// return server.Echo.StartTLS(":"+addr, "server/public/cert/cert.pem", "server/public/cert/key.pem")
}
