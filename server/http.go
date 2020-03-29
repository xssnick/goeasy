package server

import (
	"context"
	"log"
	"reflect"
	"runtime/debug"
	"time"

	//"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"

	"github.com/xssnick/goeasy"
)

type handlerInfo struct {
	typ     reflect.Type
	regFunc func()
}

type FlowMaker func(ctx context.Context) goeasy.Flow

type Server struct {
	r        *fasthttprouter.Router
	srv      fasthttp.Server
	flowMaker FlowMaker
	handlers []handlerInfo
}

type Config struct {
	ReadBufferSize  int
	WriteBufferSize int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
}

func New(cfg Config) *Server {
	s := &Server{
		r: fasthttprouter.New(),
		flowMaker: goeasy.NewBasicFlow,
	}

	s.srv = fasthttp.Server{
		Handler:         s.r.Handler,
		ReadBufferSize:  cfg.ReadBufferSize,
		WriteBufferSize: cfg.WriteBufferSize,
		ReadTimeout:     cfg.ReadTimeout,
		WriteTimeout:    cfg.WriteTimeout,
	}

	return s
}

func (s *Server) Register(method, path string, h goeasy.HttpHandler) {
	s.handlers = append(s.handlers, handlerInfo{
		typ: reflect.TypeOf(h).Elem(),
		regFunc: func() {
			h.OnInit(h)

			switch method {
			case fasthttp.MethodGet:
				s.r.GET(path, s.handler(h))
			case fasthttp.MethodPost:
				s.r.POST(path, s.handler(h))
			case fasthttp.MethodDelete:
				s.r.DELETE(path, s.handler(h))
			case fasthttp.MethodPatch:
				s.r.PATCH(path, s.handler(h))
			case fasthttp.MethodPut:
				s.r.PUT(path, s.handler(h))
			case fasthttp.MethodOptions:
				s.r.OPTIONS(path, s.handler(h))
			case fasthttp.MethodHead:
				s.r.HEAD(path, s.handler(h))
			default:
				panic("unknown method")
			}
		},
	})
}

func (s *Server) SetCustomFlowMaker(fm FlowMaker) {
	s.flowMaker = fm
}

func (s *Server) Listen(addr string) error {
	log.Println("Initializing handlers...")

	for _, h := range s.handlers {
		log.Println("Initializing", h.typ.String())
		h.regFunc()
	}

	log.Println("Initialization done!")

	log.Println("Starting http server on", addr)

	return s.srv.ListenAndServe(addr)
}

func (s *Server) Stop() {
	if err := s.srv.Shutdown(); err != nil {
		log.Println("Graceful shutdown didn't complete:", err.Error())
	}
}

func (s *Server) handler(h goeasy.HttpHandler) func(ctx *fasthttp.RequestCtx) {
	preproc := h.MiddlewareChain()

	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("recovered goeasy handler:", r)
				debug.PrintStack()
				ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			}
		}()

		flow := s.flowMaker(ctx)

		for _, p := range preproc {
			if !process(flow, ctx, p) {
				return
			}
		}

		process(flow, ctx, h)
	}
}

func process(flow goeasy.Flow, ctx *fasthttp.RequestCtx, h goeasy.HttpMiddleware) bool {
	res := h.OnHttpRequest(flow, &ctx.Request)

	if _, ok := res.(goeasy.Error); !ok {
		res = h.OnProcess(flow, res)
		if res == nil {
			return true
		}
	}

	h.OnHttpResponse(flow, res, &ctx.Response)

	return false
}
