package goeasy

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type BasicService struct {
	Handler interface{}
}

func (b *BasicService) OnInit(handler interface{}) {
	b.Handler = handler
}

func (b *BasicService) OnProcess(flow Flow, p interface{}) interface{} {
	return nil
}

func (b *BasicService) MiddlewareChain() []HttpMiddleware {
	return []HttpMiddleware{}
}

func (b *BasicService) OnHttpRequest(flow Flow, req *fasthttp.Request) interface{} {
	return nil
}

func (b *BasicService) OnHttpResponse(flow Flow, result interface{}, resp *fasthttp.Response) {
	resp.Header.SetContentType("application/json")

	if err, ok := result.(Error); ok {
		resp.Header.SetStatusCode(err.Code())
	}

	json.NewEncoder(resp.BodyWriter()).Encode(result)
}
