package internal

import (
	"sync"
	"time"
)

type Cache struct {
	timeToLive time.Duration
	queue      map[string]*Element
	mu         sync.Mutex
}

type Element struct {
	expireAt int64
	event    *ipEvent
}

func NewLocalCache(timeToLive time.Duration) *Cache {
	var wg sync.WaitGroup
	cache := &Cache{
		timeToLive: timeToLive,
		queue:      make(map[string]*Element),
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range time.NewTicker(1 * time.Millisecond).C {
			if cache.getQueueSize() == 0 {
				continue
			}
			cache.cleanup()
		}
	}()
	return cache
}

func (c *Cache) cleanup() {
	defer c.mu.Unlock()
	c.mu.Lock()
	for key, value := range c.queue {
		if time.Now().UTC().UnixMilli() >= value.expireAt {
			delete(c.queue, key)
		}
	}
}

func (c *Cache) getQueueSize() int {
	defer c.mu.Unlock()
	c.mu.Lock()
	return len(c.queue)
}

func (c *Cache) Get(key string) (*ipEvent, bool) {
	defer c.mu.Unlock()
	c.mu.Lock()
	if element := c.queue[key]; element != nil {
		return element.event, true
	}
	return nil, false
}

func (c *Cache) Set(value interface{}) {
	defer c.mu.Unlock()
	c.mu.Lock()
	switch event := value.(type) {
	case ipEvent:
		c.queue[event.Ip] = &Element{
			expireAt: time.Now().UTC().Add(c.timeToLive).UnixMilli(),
			event:    &event,
		}
	}
}
