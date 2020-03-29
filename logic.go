package goeasy

import (
	"github.com/valyala/fasthttp"
)

type Handler interface {
	OnInit(realType interface{})
	OnProcess(flow Flow, p interface{}) interface{}
}

type HttpHandler interface {
	Handler
	MiddlewareChain() []HttpMiddleware
	OnHttpRequest(flow Flow, req *fasthttp.Request) interface{}
	OnHttpResponse(flow Flow, result interface{}, resp *fasthttp.Response)
}

type HttpMiddleware interface {
	Handler
	OnHttpRequest(flow Flow, req *fasthttp.Request) interface{}
	OnHttpResponse(flow Flow, result interface{}, resp *fasthttp.Response)
}
