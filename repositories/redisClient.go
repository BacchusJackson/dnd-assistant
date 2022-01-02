package repositories

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
)

// RedisClient implements the client interface
// This abstracts the redis specific commands and returns
type RedisClient struct {
	redis RedisReaderWriter
}

var ErrDatabaseFail = errors.New("database transaction failed... see log for details")

func NewDefaultRedisClient(address string, database uint) *RedisClient {
	return &RedisClient{
		redis: redis.NewClient(&redis.Options{Addr: address, DB: int(database)}),
	}
}

func (r RedisClient) Append(key string, value string) error {
	res, err := r.redis.SAdd(context.Background(), key, value).Result()
	if err != nil {
		log.Printf("Redis Append Error: %s\n", err)
		return ErrDatabaseFail
	}
	if res != 1 {
		log.Printf("SADD Response: %d\n", res)
		return ErrDatabaseFail
	}
	return nil
}

func (r RedisClient) Update(key string, field string, value string) error {
	res, err := r.redis.HSet(context.Background(), key, field, value).Result()
	if err != nil {
		log.Printf("Redis repo HSET failed: %s\n", err)
		return ErrDatabaseFail
	}

	if res != 1 {
		log.Printf("HSET Response: %d\n", res)
		return ErrDatabaseFail
	}
	log.Printf("HSET database response: %d\n", res)
	return nil
}

func (r RedisClient) Write(key string, values map[string]string) error {
	res, err := r.redis.HSet(context.Background(), key, values).Result()
	if err != nil {
		log.Printf("Redis HSET error: %s\n", err)
		return ErrDatabaseFail
	}
	if res != int64(len(values)) {
		log.Printf("Redis HSET Response: %d\n", res)
		return ErrDatabaseFail
	}
	return nil
}

func (r RedisClient) Read(key string) (map[string]string, error) {
	res, err := r.redis.HGetAll(context.Background(), key).Result()
	if err != nil {
		log.Printf("Redis HGETALL error: %s\n", res)
		return nil, ErrDatabaseFail
	}

	log.Printf("Redis HGETALL Response: %s\n", res)
	return res, nil
}

func (r RedisClient) Ping() error {
	res, err := r.redis.Ping(context.Background()).Result()
	if err != nil {
		log.Println("failed to ping")
		return ErrDatabaseFail
	}
	if res != "PONG" {
		log.Printf("response: %s\n", res)
		return ErrDatabaseFail
	}
	return nil
}

func (r RedisClient) Clean() error {
	res, err := r.redis.FlushDB(context.Background()).Result()
	if err != nil {
		log.Println("failed to cleanTestDatabase")
		return ErrDatabaseFail
	}
	if res != "OK" {
		log.Printf("FLUSHDB database response: %s\n", res)
		return ErrDatabaseFail
	}
	return nil
}
