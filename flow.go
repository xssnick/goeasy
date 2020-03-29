package goeasy

import (
	"context"
	"time"
)

type Flow interface {
	context.Context

	UpdateContext(ctx context.Context)
}

type BasicFlow struct {
	ctx context.Context
}

func (f *BasicFlow) Deadline() (deadline time.Time, ok bool) {
	return f.ctx.Deadline()
}

func (f *BasicFlow) Done() <-chan struct{} {
	return f.ctx.Done()
}

func (f *BasicFlow) Err() error {
	return f.ctx.Err()
}

func (f *BasicFlow) Value(key interface{}) interface{} {
	return f.ctx.Value(key)
}

func (f *BasicFlow) UpdateContext(ctx context.Context) {
	f.ctx = ctx
}

func NewBasicFlow(ctx context.Context) Flow {
	return &BasicFlow{ctx: ctx}
}
