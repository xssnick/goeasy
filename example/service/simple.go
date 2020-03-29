package service

import (
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/xssnick/goeasy"
)

type SimpleRequest struct {
	ID int64
}

type SimpleResponse struct {
	Name string
}

type Simple struct {
	Service
}

func (s *Simple) MiddlewareChain() []goeasy.HttpMiddleware {
	return append(s.Service.MiddlewareChain(), new(Auth))
}

func (s *Simple) OnHttpRequest(flow goeasy.Flow, req *fasthttp.Request) interface{} {
	sid := flow.Value("id").(string)

	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		return goeasy.NewBasicError(goeasy.ErrCodeBadRequest, "err: %v", err)
	}

	return SimpleRequest{ID: id}
}

func (s *Simple) OnProcess(flow goeasy.Flow, p interface{}) interface{} {
	req := p.(SimpleRequest)

	if req.ID == 0 {
		return goeasy.NewBasicError(goeasy.ErrCodeBadRequest, "id=%d", req.ID)
	}

	if req.ID == 7 {
		return SimpleResponse{Name: "Super Man > DB:" + s.Database}
	}

	return SimpleResponse{Name: "unknown"}
}
