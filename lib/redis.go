package lib

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisOptions struct {
	URL string
}

type Redis struct {
	Client *redis.Client
}

func NewRedis(opts RedisOptions) *Redis {
	opt, err := redis.ParseURL(opts.URL)
	if err != nil {
		log.Fatalf("Redis parse url: %s", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := redis.NewClient(opt)
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect redis: %s", err)
	}
	return &Redis{Client: client}
}

func (r *Redis) Close() {
	if err := r.Client.Close(); err != nil {
		log.Errorf("Redis error in closing the connection: %s", err)
	}
}

func (r *Redis) JSONSet(key string, value interface{}, expiration time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Client.Set(context.TODO(), key, string(b), expiration).Err()
}

func (r *Redis) JSONGet(key string, result interface{}) error {
	val, err := r.Client.Get(context.TODO(), key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), result)
}
