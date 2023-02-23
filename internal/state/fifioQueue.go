package state

import (
	"fmt"
	"sync"
)

type Fifo struct {
	queue []string
	mu    sync.Mutex
}

// Create a new first-in first-out queue with a size limit.
// Note: if size is 0 then no size is set.
func NewFifoQueue() *Fifo {
	return &Fifo{
		queue: make([]string, 0),
	}
}

func (f *Fifo) findElement(v string) bool {
	for i := 0; i < len(f.queue); i++ {
		e := f.queue[i]
		if e == v {
			return true
		}
	}
	return false
}

func (f *Fifo) GetLen() int {
	defer f.mu.Unlock()
	f.mu.Lock()
	return len(f.queue)
}

func (f *Fifo) Enqueue(v string) error {
	defer f.mu.Unlock()
	f.mu.Lock()
	if len(f.queue) != 0 && f.findElement(v) {
		return fmt.Errorf("element: %s already queued", v)
	}
	f.queue = append(f.queue, v)
	return nil
}

func (f *Fifo) Dequeue() string {
	defer f.mu.Unlock()
	f.mu.Lock()
	e := f.queue[0]
	f.queue[0] = ""
	f.queue = f.queue[1:]
	return e
}
