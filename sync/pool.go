package sync

import stdsync "sync"

type Pool[T any] struct {
	pool stdsync.Pool
}

func NewPool[T any](newfunc func() any) *Pool[T] {
	return &Pool[T]{
		pool: stdsync.Pool{
			New: newfunc,
		},
	}
}

func (p *Pool[T]) Put(x T) {
	p.pool.Put(x)
}

func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}
