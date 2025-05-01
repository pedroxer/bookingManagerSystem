package common_errors

import (
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FormError(ctx *fasthttp.RequestCtx, errMessage string, errStatusCode int) {
	ctx.SetContentType("application/json")
	ctx.SetBodyString("{\"error\": \"" + errMessage + "\"}")
	ctx.SetStatusCode(errStatusCode)
}

func CastGrpcErrors(err error) int {
	st, _ := status.FromError(err)
	st.Code()
	switch st.Code() {
	case codes.Internal:
		return fasthttp.StatusInternalServerError
	case codes.InvalidArgument:
		return fasthttp.StatusBadRequest
	case codes.NotFound:
		return fasthttp.StatusNotFound
	default:
		return fasthttp.StatusInternalServerError
	}
	return fasthttp.StatusOK
}

func CastGrpcMessage(err error) string {
	st, _ := status.FromError(err)
	return st.Message()
}
