package cache

import (
	"github.com/go-redis/redis"
)

type IRedis interface {
	Client() *redis.Client
	Connect() error
}
