package service

import (
	"github.com/valyala/fasthttp"

	"github.com/xssnick/goeasy"
)

type AuthRequest struct {
	Token string
}

type Auth struct {
	Service
}

func (s *Auth) OnHttpRequest(flow goeasy.Flow, req *fasthttp.Request) interface{} {
	return AuthRequest{Token: string(req.Header.Peek("Auth"))}
}

func (s *Auth) OnProcess(flow goeasy.Flow, p interface{}) interface{} {
	req := p.(AuthRequest)

	if req.Token == "" {
		return goeasy.NewBasicError(goeasy.ErrCodeUnauthorized, "empty token")
	}

	return nil
}
