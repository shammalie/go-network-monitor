package internal

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

const (
	redisAddr     = "REDIS_URL"
	redisPassword = "REDIS_PASSWORD"
	redisDb       = "REDIS_DB"
	channelName   = "network-monitor-servers"
)

type RedisClient struct {
	client       *redis.Client
	itemDuration time.Duration
	ctxTime      int
	queue        map[string]interface{}
	mu           sync.Mutex
}

func NewRedisClient(ctxTimeSeconds int, duration time.Duration) *RedisClient {
	var wg sync.WaitGroup
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
		client:       redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: db}),
		itemDuration: duration,
		ctxTime:      ctxTimeSeconds,
		queue:        map[string]interface{}{},
	}
	rdb.loadKeys()
	wg.Add(1)
	go func() {
		defer wg.Done()
		rdb.appChannel()
	}()
	fmt.Println("started redis client")
	return rdb
}

func (r *RedisClient) loadKeys() {
	var cursor uint64
	defer r.mu.Unlock()
	r.mu.Lock()
	for {
		var keys []string
		var err error
		keys, cursor, err = r.client.Scan(context.Background(), cursor, "prefix:*", 0).Result()
		if err != nil {
			panic(err)
		}
		for _, key := range keys {
			r.queue[key] = struct{}{}
		}
		if cursor == 0 { // no more keys
			break
		}
	}
}

func (r *RedisClient) appChannel() {
	sub := r.client.Subscribe(context.Background(), channelName)
	defer sub.Close()
	subCh := sub.Channel()
	for msg := range subCh {
		r.mu.Lock()
		r.queue[msg.Payload] = msg.Payload
		r.mu.Unlock()
	}
}

func (r *RedisClient) getInMemoryQueue(key string) bool {
	defer r.mu.Unlock()
	r.mu.Lock()
	return r.queue[key] == nil
}

func (r *RedisClient) Set(key string, value interface{}) error {
	defer r.mu.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.ctxTime)*time.Second)
	defer cancel()
	err := r.client.Set(ctx, key, value, r.itemDuration).Err()
	if err != nil {
		return err
	}
	r.mu.Lock()
	r.queue[key] = value
	err = r.client.Publish(ctx, channelName, key).Err()
	if err != nil {
		return err
	}
	return nil
}
