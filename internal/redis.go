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
	client    *redis.Client
	ctxTime   int
	size      int64
	mu        sync.Mutex
	warmCache *Cache
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
		client:    redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: db}),
		ctxTime:   ctxTimeSeconds,
		warmCache: NewLocalCache(5 * time.Second),
	}
	fmt.Println("loaded redis client")
	return rdb
}

func getTimeoutContext(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}

func (r *RedisClient) loadCacheEntries() ([]ipEvent, error) {
	var entries []ipEvent
	ctx := context.Background()
	iter := r.client.Scan(ctx, 0, "", 0).Iterator()
	for iter.Next(ctx) {
		ip := iter.Val()
		cacheEvent, err := r.Get(ip)
		if err != nil {
			return nil, err
		}
		entries = append(entries, *cacheEvent)
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	err := r.getDbSize()
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *RedisClient) getDbSize() error {
	defer r.mu.Unlock()
	ctx, cancel := getTimeoutContext(time.Duration(r.ctxTime) * time.Second)
	defer cancel()
	size, err := r.client.DBSize(ctx).Result()
	if err != nil {
		return err
	}
	r.mu.Lock()
	r.size = size
	return nil
}

func (r *RedisClient) Get(key string) (*ipEvent, error) {
	if event, found := r.warmCache.Get(key); found {
		return event, nil
	}
	ctx, cancel := getTimeoutContext(time.Duration(r.ctxTime) * time.Second)
	defer cancel()
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, err
	} else if err != nil {
		panic(err)
	}
	var event ipEvent
	err = json.Unmarshal([]byte(val), &event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *RedisClient) Set(key string, value interface{}) error {
	r.warmCache.Set(value)
	ctx, cancel := getTimeoutContext(time.Duration(r.ctxTime) * time.Second)
	defer cancel()
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = r.client.Set(ctx, key, b, 6*time.Hour).Err()
	if err != nil {
		return err
	}
	err = r.getDbSize()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Del(key string) error {
	ctx, cancel := getTimeoutContext(time.Duration(r.ctxTime) * time.Second)
	defer cancel()
	_, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	err = r.getDbSize()
	if err != nil {
		return err
	}
	return nil
}
