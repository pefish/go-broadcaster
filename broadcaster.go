package go_broadcaster

import (
	"sync"
)

type Broadcaster[T any] struct {
	chans sync.Map
}

func NewBroadcaster[T any]() *Broadcaster[T] {
	return &Broadcaster[T]{}
}

func (b *Broadcaster[T]) Sub(key string) {
	b.chans.Store(key, make(chan T, 0))
}

func (b *Broadcaster[T]) C(key string) <-chan T {
	v, ok := b.chans.Load(key)
	if !ok {
		return nil
	}
	return v.(chan T)
}

func (b *Broadcaster[T]) Pub(value T) {
	b.chans.Range(func(key, v any) bool {
		v.(chan T) <- value
		return true
	})
}
