package routes

import (
	"fmt"
	fasthttprouter "github.com/fasthttp/router"
	config "github.com/pedroxer/BookingManagerSystem/internal/configs"
	proto_gen "github.com/pedroxer/BookingManagerSystem/internal/proto_gen/protos"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"time"
)

type Router struct {
	rtr            *fasthttprouter.Router
	srv            *fasthttp.Server
	logger         *logrus.Logger
	cfg            *config.Config
	resourceClient proto_gen.ResourceServiceClient
	bookingClient  proto_gen.BookingServiceClient
}

func NewRouter(logger *logrus.Logger, cfg *config.Config, resourceClient proto_gen.ResourceServiceClient, bookingClient proto_gen.BookingServiceClient) *Router {

	rtr := fasthttprouter.New()
	r := &Router{
		rtr: rtr,
		srv: &fasthttp.Server{
			Handler:            cors(rtr.Handler),
			MaxRequestBodySize: 100_000_000,
			ReadTimeout:        time.Duration(cfg.Api.ReadTimeout) * time.Second,
			WriteTimeout:       time.Duration(cfg.Api.WriteTimeout) * time.Second,
			IdleTimeout:        time.Duration(cfg.Api.IdleTimeout) * time.Second},
		logger:         logger,
		cfg:            cfg,
		resourceClient: resourceClient,
		bookingClient:  bookingClient,
	}

	registerResource(r)
	r.rtr.GET("/status", statusHandler)
	r.rtr.GET("/", test)
	r.rtr.HandleMethodNotAllowed = true
	r.rtr.MethodNotAllowed = methodNotAllowedHandler
	r.rtr.NotFound = notFoundHandler
	return r
}

func (r *Router) Start() error {
	return r.srv.ListenAndServe(fmt.Sprintf(":%d", r.cfg.Api.Port))
}

func (r *Router) Shutdown() error {
	return r.srv.Shutdown()
}

func test(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}
func methodNotAllowedHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
}

func notFoundHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}

// Хендлер для прохождения liveness/readiness проб kubernetes.
func statusHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func cors(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		handler(ctx)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "authorization, content-type, set-cookie, cookie, server")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, DELETE")
	}
}
