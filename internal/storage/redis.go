package storage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Storage interface {
	Set(ctx context.Context, key, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type Redis struct {
	client *redis.Client
}

func NewRedis(addr, password string) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})
	return &Redis{client: rdb}
}

func (r *Redis) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
