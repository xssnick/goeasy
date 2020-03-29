package goeasy

import (
	"context"
)

type Flow interface {
	Context() context.Context
	UpdateContext(ctx context.Context)
}

type BasicFlow struct {
	ctx context.Context
}

func (f *BasicFlow) Context() context.Context {
	return f.ctx
}

func (f *BasicFlow) UpdateContext(ctx context.Context) {
	f.ctx = ctx
}

func NewBasicFlow(ctx context.Context) Flow {
	return &BasicFlow{ctx: ctx}
}
