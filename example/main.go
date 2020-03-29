package main

import (
	"log"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/xssnick/goeasy/example/service"
	"github.com/xssnick/goeasy/server"
)

func main() {
	srv := server.New(server.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
	})

	svc := service.Service{Database: "SQL"}
	srv.Register(fasthttp.MethodGet, "/:id", &service.Simple{Service: svc})

	log.Println(srv.Listen(":7777"))
}
