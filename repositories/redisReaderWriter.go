package repositories

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// RedisReaderWriter can be implemented to mock redis specific commands for testing
type RedisReaderWriter interface {
	HSet(ctx context.Context, key string, value ...interface{}) *redis.IntCmd
	SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	FlushDB(ctx context.Context) *redis.StatusCmd
	HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd
	Ping(ctx context.Context) *redis.StatusCmd
}
