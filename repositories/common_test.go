package repositories

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

const RedisAddress string = "localhost:6379"
const RedisTestDB uint = 2

var MockIntCmd = redis.NewIntCmd(context.Background(), nil)
var MockStringStringMapCmd = redis.NewStringStringMapCmd(context.Background(), nil)
var MockStatusCmd = redis.NewStatusCmd(context.Background(), nil)
var MockMap map[string]string
var MockString string
var MockError error

// NewMockRedisClient a mocked client that doesn't hit a live database
func NewMockRedisClient() *RedisClient {
	client := NewDefaultRedisClient(RedisAddress, RedisTestDB)
	client.redis = &MockReaderWriter{}

	return client
}

type MockReaderWriter struct{}

func (m MockReaderWriter) HSet(_ context.Context, _ string, _ ...interface{}) *redis.IntCmd {
	return MockIntCmd
}
func (m MockReaderWriter) SAdd(_ context.Context, _ string, _ ...interface{}) *redis.IntCmd {
	return MockIntCmd
}

func (m MockReaderWriter) HGetAll(_ context.Context, _ string) *redis.StringStringMapCmd {
	return MockStringStringMapCmd
}

func (m MockReaderWriter) Ping(_ context.Context) *redis.StatusCmd {
	return MockStatusCmd
}

func (m MockReaderWriter) FlushDB(_ context.Context) *redis.StatusCmd {
	return MockStatusCmd
}

func checkError(t *testing.T, expected interface{}, got interface{}) {
	if got != expected {
		t.Errorf("expected %v ... got %v\n", expected, got)
	}
}

type MockEntity struct{}

func (m MockEntity) Map() map[string]string {
	return MockMap
}

func (m MockEntity) EntityKey() string {
	return MockString
}

func (m MockEntity) Validate() error {
	return MockError
}
