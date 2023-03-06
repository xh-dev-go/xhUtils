package channels

import "sync"

type MChannel[T any] struct {
	out     []chan T
	channel chan T
	in      []InChan[T]
}

type InChan[T any] struct {
	Channel chan<- T
	Process func()
	Done    func()
}

func NewMChannel[T any]() *MChannel[T] {
	return &MChannel[T]{
		in:      make([]InChan[T], 0),
		out:     make([]chan T, 0),
		channel: make(chan T),
	}
}

func (m *MChannel[T]) inChan() chan<- T {
	c := make(chan T)
	d := make(chan struct{})
	var blockWg sync.WaitGroup
	blockWg.Add(1)
	inChan := InChan[T]{

		Channel: c,
		Process: func() {
			blockWg.Done()
			for {
				select {
				case v := <-c:
					for _, c := range m.out {
						c <- v
					}
				case <-d:
					return
				}
			}
		},
		Done: func() {
			d <- struct{}{}
		},
	}
	go inChan.Process()
	blockWg.Wait()

	m.in = append(m.in, inChan)
	return inChan.Channel
}

func (m *MChannel[T]) outChan() <-chan T {
	c := make(chan T)
	m.out = append(m.out, c)

	return c
}
