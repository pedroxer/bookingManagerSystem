package routes

import (
	fasthttprouter "github.com/fasthttp/router"
	"github.com/pedroxer/BookingManagerSystem/internal/storage"
	"github.com/valyala/fasthttp"
)

type Router struct {
	port    int
	storage storage.Storage
	rtr     fasthttprouter.Router
	srv     fasthttp.Server
}
