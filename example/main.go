package main

import (
	"log"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/xssnick/goeasy/example/handler"
	"github.com/xssnick/goeasy/server"
)

func main() {
	srv := server.New(server.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
	})

	svc := handler.Service{Database: "SQL"}
	srv.MustRegister(fasthttp.MethodGet, "/:id", &handler.Simple{Service: svc})

	log.Println(srv.Listen(":7777"))
}
