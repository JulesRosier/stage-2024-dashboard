package broadcast

import (
	"context"
	"log/slog"
)

type BroadcastServer[T any] interface {
	Subscribe() <-chan T
	CancelSubscription(<-chan T)
}

type broadcastServer[T any] struct {
	source         <-chan T
	listeners      []chan T
	addListener    chan chan T
	removeListener chan (<-chan T)
}

func (s *broadcastServer[T]) Subscribe() <-chan T {
	newListener := make(chan T, 100)
	s.addListener <- newListener
	return newListener
}

func (s *broadcastServer[T]) CancelSubscription(channel <-chan T) {
	s.removeListener <- channel
}

func (s *broadcastServer[T]) serve(ctx context.Context) {
	defer func() {
		for _, listener := range s.listeners {
			if listener != nil {
				close(listener)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			slog.Info("Stopping broadcaster server")
			return
		case newListener := <-s.addListener:
			s.listeners = append(s.listeners, newListener)
		case listenerToRemove := <-s.removeListener:
			for i, ch := range s.listeners {
				if ch == listenerToRemove {
					// replace item we want to remove with last item
					s.listeners[i] = s.listeners[len(s.listeners)-1]
					// remove last item
					s.listeners = s.listeners[:len(s.listeners)-1]
					close(ch)
					break
				}
			}
		case val, ok := <-s.source:
			if !ok {
				return
			}
			for _, listener := range s.listeners {
				if listener != nil {
					select {
					case listener <- val:
					case <-ctx.Done():
						return
					}

				}
			}
		}
	}
}

func NewBroadcastServer[T any](ctx context.Context, source <-chan T) BroadcastServer[T] {
	service := &broadcastServer[T]{
		source:         source,
		listeners:      make([]chan T, 0),
		addListener:    make(chan chan T),
		removeListener: make(chan (<-chan T)),
	}
	go service.serve(ctx)
	return service
}
