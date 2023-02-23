package state

import (
	"container/list"
	"sync"
)

// Priority list will reorder elements based on occurance with the aim to optimise
// the looping requirements when attempting to match elements.
type Prioritylist struct {
	list *list.List
	size int
	mu   sync.Mutex
}

// Create a new priority list with provided elements.
func NewPrioritylist(elements []string) *Prioritylist {
	q := &Prioritylist{
		list: list.New(),
		size: len(elements),
	}
	for _, e := range elements {
		q.list.PushBack(e)
	}
	return q
}

// Lookup elements in the list and return a boolean outcome.
// When an element is matched it will be moved to the top of the list and return true.
func (p *Prioritylist) Lookup(v string) bool {
	defer p.mu.Unlock()
	p.mu.Lock()
	if v == "" {
		return false
	}
	for e := p.list.Front(); e != nil; e = e.Next() {
		if e.Value != v {
			continue
		}
		p.list.MoveToFront(e)
		return true
	}
	return false
}
