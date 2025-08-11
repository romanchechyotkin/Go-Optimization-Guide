package main

import "sync"

type Batcher[T any] struct {
	mu     sync.Mutex
	buffer []T
	size   int
	flush  func([]T)
}

func NewBatcher[T any](size int, flush func([]T)) *Batcher[T] {
	return &Batcher[T]{
		buffer: make([]T, 0, size),
		size:   size,
		flush:  flush,
	}
}

func (b *Batcher[T]) Add(item T) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.buffer = append(b.buffer, item)
	if len(b.buffer) >= b.size {
		b.flushNow()
	}
}

func (b *Batcher[T]) flushNow() {
	if len(b.buffer) == 0 {
		return
	}
	b.flush(b.buffer)
	b.buffer = b.buffer[:0]
}
