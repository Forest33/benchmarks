package main

import (
	"bytes"
	"sync"
)

type BytesPool struct {
	pool sync.Pool
}

func NewBytesPool() *BytesPool {
	return &BytesPool{
		pool: sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}

func (p *BytesPool) Get() *bytes.Buffer {
	return p.pool.Get().(*bytes.Buffer)
}

func (p *BytesPool) Release(b *bytes.Buffer) {
	b.Reset()
	p.pool.Put(b)
}
