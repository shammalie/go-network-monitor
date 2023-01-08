package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	redisAddr     = "REDIS_URL"
	redisPassword = "REDIS_PASSWORD"
	redisDb       = "REDIS_DB"
	channelName   = "network-monitor-servers"
)

type RedisClient struct {
	client  *redis.Client
	ctxTime int
	cache   map[string]interface{}
	mu      sync.Mutex
}

type ipEvent struct {
	Id        primitive.ObjectID
	Ip        string
	Timestamp int64
}

func NewRedisClient(ctxTimeSeconds int) *RedisClient {
	if ctxTimeSeconds == 0 {
		fmt.Println("setting ctx time to 5 seconds.")
		ctxTimeSeconds = 5
	}
	addr := viper.GetString(redisAddr)
	if redisAddr == "" {
		panic("no redis url provided")
	}
	password := viper.GetString(redisPassword)
	db := viper.GetInt(redisDb)
	rdb := &RedisClient{
		client:  redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: db}),
		ctxTime: ctxTimeSeconds,
		cache:   make(map[string]interface{}, 1000),
	}
	fmt.Println("loaded redis client")
	return rdb
}

func getTimeoutContext(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}

func (r *RedisClient) loadCacheEntries() ([]ipEvent, error) {
	var entries []ipEvent
	ctx, cancel := getTimeoutContext(time.Duration(r.ctxTime) * time.Second)
	defer cancel()
	iter := r.client.Scan(ctx, 0, "", 0).Iterator()
	for iter.Next(ctx) {
		var event ipEvent
		event.Ip = iter.Val()
		cacheEvent, err := r.Get(event.Ip)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(cacheEvent), &event)
		if err != nil {
			return nil, err
		}
		entries = append(entries, event)
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *RedisClient) inLocalCache(key string) bool {
	defer r.mu.Unlock()
	r.mu.Lock()
	return r.cache[key] != nil
}

func (r *RedisClient) getIpEventFromLocalCache(key string) *ipEvent {
	defer r.mu.Unlock()
	r.mu.Lock()
	entry := r.cache[key]
	switch event := entry.(type) {
	case ipEvent:
		return &event
	default:
		return nil
	}
}

func (r *RedisClient) addToLocalCache(key string, value interface{}) {
	defer r.mu.Unlock()
	r.mu.Lock()
	r.cache[key] = value
}

func (r *RedisClient) removeFromLocalCache(key string) {
	defer r.mu.Unlock()
	r.mu.Lock()
	delete(r.cache, key)
}

func (r *RedisClient) inCache(key string) bool {
	if r.inLocalCache(key) {
		return true
	}
	_, err := r.Get(key)
	if err == redis.Nil {
		return false
	} else if err != nil {
		panic(err)
	}
	return true
}

func (r *RedisClient) Get(key string) (string, error) {
	ctx, cancel := getTimeoutContext(time.Duration(r.ctxTime) * time.Second)
	defer cancel()
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil {
		panic(err)
	}
	return val, nil
}

func (r *RedisClient) Set(key string, value interface{}, duration time.Duration) error {
	ctx, cancel := getTimeoutContext(time.Duration(r.ctxTime) * time.Second)
	defer cancel()
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = r.client.Set(ctx, key, b, duration).Err()
	if err != nil {
		return err
	}
	r.addToLocalCache(key, value)
	return nil
}

func (r *RedisClient) Del(key string) error {
	ctx, cancel := getTimeoutContext(time.Duration(r.ctxTime) * time.Second)
	defer cancel()
	_, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	r.removeFromLocalCache(key)
	return nil
}
