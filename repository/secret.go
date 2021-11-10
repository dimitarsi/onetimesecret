package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type SecretRepository interface {
	GetDel(key string) (map[string]string, error)
	Set(key string, value interface{}) error
}

type RedisSecretsRepository struct {
	Redis *redis.Client
	Expires time.Duration
}

func (r *RedisSecretsRepository) GetDel(key string) (map[string]string, error) {
	ctx := context.Background()
	rawData, err := r.Redis.GetDel(ctx, key).Result()

	if err != nil {
		return nil, err
	}

	data := make(map[string]string)

	err = json.Unmarshal([]byte(rawData), &data)

	return data, err
}

func (r *RedisSecretsRepository) Set(key string, value interface{}) error {

	strValue, err := json.Marshal(value)

	if err != nil {
		return err
	}
	
	ctx := context.Background()
	return r.Redis.Set(ctx, key, strValue, r.Expires).Err()
}

func NewRedisSecretsRepository() *RedisSecretsRepository {
	return &RedisSecretsRepository {
		Redis: redis.NewClient(&redis.Options{
			Addr:     "redisdb:6379",
			Password: "",
			DB:       0,
		}),
		Expires: time.Minute * 20,
	}
}