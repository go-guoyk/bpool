package bpool

import (
	"bytes"
	"sync"
)

type ReturnFunc func()

type Pool interface {
	Retrieve() (*bytes.Buffer, ReturnFunc)
}

type pool struct {
	*sync.Pool
}

func (p *pool) Retrieve() (*bytes.Buffer, ReturnFunc) {
	buf := p.Get().(*bytes.Buffer)
	buf.Reset()
	return buf, func() {
		p.Put(buf)
	}
}

func New() Pool {
	return &pool{
		Pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}
