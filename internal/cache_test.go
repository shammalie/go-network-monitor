package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestNewLocalCache(t *testing.T) {
	timeToLive := 5 * time.Second
	cache := NewLocalCache(timeToLive)

	if cache == nil || cache.timeToLive.Seconds() != timeToLive.Seconds() {
		t.Errorf("cache should be created with matching timeToLive durations")
	}
}

func TestCache_cleanup(t *testing.T) {
	cache := NewLocalCache(1 * time.Millisecond)
	cache.Set("127.0.0.1")
	time.Sleep(100 * time.Millisecond)

	if len(cache.queue) != 0 {
		t.Errorf("Cleanup function should of removed the event")
	}
}

func TestCache_getQueueSize(t *testing.T) {
	cache := NewLocalCache(10 * time.Second)
	for i := 0; i < 5; i++ {
		cache.Set(fmt.Sprintf("127.0.0.%d", i+1))
	}
	if len(cache.queue) != 5 {
		t.Errorf("cache size should be 5")
	}
}

func TestCache_Get(t *testing.T) {
	ip := "127.0.0.1"
	cache := NewLocalCache(10 * time.Second)
	cache.Set(ip)
	element, found := cache.Get(ip)
	if !found || element.ip != ip {
		t.Errorf("the ip should match")
	}
}

func TestCache_Set(t *testing.T) {
	cache := NewLocalCache(10 * time.Second)
	cache.Set("127.0.0.1")
	if len(cache.queue) != 1 {
		t.Errorf("event should of been set in queue, length should return 1")
	}
}

func BenchmarkCache(b *testing.B) {
	ip := "127.0.0.1"
	cache := NewLocalCache(5 * time.Second)
	b.ResetTimer()
	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache.Set(ip)
		}
	})
	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			event, found := cache.Get(ip)
			if found {
				_ = event
			}
		}
	})
}
