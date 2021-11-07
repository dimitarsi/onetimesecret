package request

import "github.com/go-redis/redis"

type GenericSecretRequest struct {
	Redis redis.Client
}