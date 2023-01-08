package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
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
	client        *redis.Client
	subscriptions *redis.PubSub
	ctxTime       int
}

type ipEvent struct {
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
	client := redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: db})
	rdb := &RedisClient{
		client:        client,
		subscriptions: client.Subscribe(context.Background(), channelName),
		ctxTime:       ctxTimeSeconds,
	}
	rdb.loadKeys()
	fmt.Println("loaded redis client")
	return rdb
}

func getTimeoutContext(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}

func (r *RedisClient) loadKeys() {
	ctx, cancel := getTimeoutContext(time.Duration(r.ctxTime) * time.Second)
	defer cancel()
	iter := r.client.Scan(ctx, 0, "", 0).Iterator()
	for iter.Next(ctx) {
		e := ipEvent{
			Ip: iter.Val(),
		}
		timestampStr, err := r.Get(e.Ip)
		if err != nil {
			fmt.Println(err)
		}
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			panic(err)
		}
		e.Timestamp = timestamp
		b, err := json.Marshal(e)
		if err != nil {
			panic(err)
		}
		err = r.client.Publish(ctx, channelName, string(b)).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

func (r *RedisClient) inCache(key string) bool {
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
	err := r.client.Set(ctx, key, value, duration).Err()
	if err != nil {
		return err
	}
	b, err := json.Marshal(ipEvent{Ip: key, Timestamp: value.(int64)})
	if err != nil {
		return err
	}
	err = r.client.Publish(ctx, channelName, string(b)).Err()
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
	return nil
}
