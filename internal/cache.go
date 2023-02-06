package internal

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cache struct {
	timeToLive time.Duration
	queue      map[string]*Element
	mu         sync.Mutex
}

type Element struct {
	expireAt int64
	id       primitive.ObjectID
	ip       string
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

func (c *Cache) Get(key string) (*Element, bool) {
	defer c.mu.Unlock()
	c.mu.Lock()
	if element := c.queue[key]; element != nil {
		return element, true
	}
	return nil, false
}

func (c *Cache) Set(value interface{}) {
	defer c.mu.Unlock()
	c.mu.Lock()
	switch e := value.(type) {
	case Element:
		if c.queue[e.ip] != nil {
			return
		}
		for _, v := range c.queue {
			if e.ip == v.ip {
				return
			}
		}
		e.expireAt = time.Now().UTC().Add(c.timeToLive).UnixMilli()
		c.queue[e.ip] = &e
	}
}
