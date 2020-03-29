package handler

import (
	"github.com/xssnick/goeasy"
)

type Service struct {
	goeasy.BasicService

	Database string
}

func (s *Service) MiddlewareChain() []goeasy.HttpMiddleware {
	return []goeasy.HttpMiddleware{NewLog(s.Handler)}
}
