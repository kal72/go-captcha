package redisstore

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type Redis struct {
	client *goredis.Client
	prefix string
	ctx    context.Context
}
type Config struct {
	Addr     string
	Password string
	DB       int    // default 0
	Prefix   string // prefix for key captcha
}

func New(cfg Config) *Redis {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return &Redis{
		client: rdb,
		prefix: cfg.Prefix,
		ctx:    context.Background(),
	}
}

func (r *Redis) makeKey(key string) string {
	if r.prefix == "" {
		return "captcha:" + key
	}
	return r.prefix + ":" + key
}

func (r *Redis) Set(key string, value string, ttl time.Duration) error {
	redisKey := r.makeKey(key)
	return r.client.Set(r.ctx, redisKey, value, ttl).Err()
}

func (r *Redis) Get(key string) (string, error) {
	redisKey := r.makeKey(key)
	val, err := r.client.Get(r.ctx, redisKey).Result()
	if err == goredis.Nil {
		return "", ErrNotFound
	}
	return val, err
}

func (r *Redis) Delete(key string) error {
	redisKey := r.makeKey(key)
	return r.client.Del(r.ctx, redisKey).Err()
}

var ErrNotFound = goredis.Nil
