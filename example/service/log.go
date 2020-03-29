package service

import (
	"log"
	"reflect"

	"github.com/valyala/fasthttp"

	"github.com/xssnick/goeasy"
)

type Log struct {
	Service
	method string
}

func (l *Log) OnHttpRequest(flow goeasy.Flow, req *fasthttp.Request) interface{} {
	log.Println(l.method, req.URI().String())
	return nil
}

func NewLog(target interface{}) *Log {
	return &Log{
		method: reflect.TypeOf(target).String(),
	}
}
