package repositories

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"testing"
)

const RedisAddress string = "localhost:6379"

var mockResponse = "error"

// Mocked Repository for testing
func newMockedRepo() *RedisClient {
	client := NewDefaultRedisClient(RedisAddress, 2)
	client.redis = &MockClient{}
	return client
}

type MockClient struct{}

func (m MockClient) HSet(ctx context.Context, key string, value ...interface{}) *redis.IntCmd {
	intCmd := redis.NewIntCmd(ctx, nil)
	switch key {
	case "error":
		intCmd.SetErr(ErrDatabaseFail)
		return intCmd
	case "fail":
		intCmd.SetVal(-1)
		return intCmd
	default:
		intCmd.SetVal(1)
		return intCmd
	}
}
func (m MockClient) SAdd(ctx context.Context, key string, value ...interface{}) *redis.IntCmd {
	intCmd := redis.NewIntCmd(ctx, nil)
	switch key {
	case "error":
		intCmd.SetErr(ErrDatabaseFail)
		return intCmd
	case "fail":
		intCmd.SetVal(-1)
		return intCmd
	default:
		intCmd.SetVal(1)
		return intCmd
	}
}

func (m MockClient) HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {
	cmd := redis.NewStringStringMapCmd(ctx, nil)
	switch key {
	case "error":
		cmd.SetErr(ErrDatabaseFail)
		return cmd
	default:
		cmd.SetVal(map[string]string{"field-1": "value-1"})
		return cmd
	}
}

func (m MockClient) Ping(ctx context.Context) *redis.StatusCmd {
	cmd := redis.NewStatusCmd(ctx, nil)

	switch mockResponse {
	case "error":
		cmd.SetErr(ErrDatabaseFail)
		return cmd
	case "fail":
		cmd.SetVal("fail")
		return cmd
	default:
		cmd.SetVal("PONG")
		return cmd
	}
}

func (m MockClient) FlushDB(ctx context.Context) *redis.StatusCmd {
	cmd := redis.NewStatusCmd(ctx, nil)
	switch mockResponse {
	case "error":
		cmd.SetErr(ErrDatabaseFail)
		return cmd
	case "fail":
		cmd.SetVal("fail")
		return cmd
	default:
		cmd.SetVal("OK")
		return cmd
	}
}

// Internal test functions
func clean() {
	client := NewDefaultRedisClient(RedisAddress, 2)
	err := client.Clean()
	if err != nil {
		fmt.Printf("Failed to clean Redis Database: %s\n", err)
		os.Exit(-1)
	}
}

func checkError(t *testing.T, expected interface{}, got interface{}) {
	if got != expected {
		t.Errorf("expected %v ... got %v\n", expected, got)
	}
}

// Unit Tests
func TestRedisRepo_Append(t *testing.T) {
	clean()
	client := newMockedRepo()

	err := client.Append("error", "")
	checkError(t, ErrDatabaseFail, err)

	err = client.Append("fail", "")
	checkError(t, ErrDatabaseFail, err)

	client = NewDefaultRedisClient(RedisAddress, 2)
	err = client.Append("test-key", "test-value")
	checkError(t, nil, err)
}

func TestRedisRepo_Set(t *testing.T) {
	clean()
	client := newMockedRepo()
	err := client.Set("error", "test-field", "test-value")
	checkError(t, ErrDatabaseFail, err)

	err = client.Set("fail", "test-field", "test-value")
	checkError(t, ErrDatabaseFail, err)

	client = NewDefaultRedisClient(RedisAddress, 2)

	err = client.Set("test-key", "test-field", "test-value")
	checkError(t, nil, err)
}

func TestRedisRepo_SetMap(t *testing.T) {
	clean()
	client := newMockedRepo()

	err := client.SetMap("error", nil)
	checkError(t, ErrDatabaseFail, err)

	err = client.SetMap("fail", nil)
	checkError(t, ErrDatabaseFail, err)

	client = NewDefaultRedisClient(RedisAddress, 2)
	err = client.SetMap("test-key", map[string]string{
		"field-1": "value-1",
		"field-2": "value-2",
	})
	checkError(t, nil, err)
}

func TestRedisRepo_Get(t *testing.T) {
	clean()
	client := newMockedRepo()
	res, err := client.Get("error")
	checkError(t, ErrDatabaseFail, err)

	client = NewDefaultRedisClient(RedisAddress, 2)
	err = client.Set("test-key", "test-field", "test-value")
	checkError(t, nil, err)

	res, err = client.Get("test-key")
	checkError(t, nil, err)
	checkError(t, "test-value", res["test-field"])

}

func TestRedisRepo_Ping(t *testing.T) {
	client := newMockedRepo()
	mockResponse = "error"
	err := client.Ping()
	checkError(t, ErrDatabaseFail, err)

	mockResponse = "fail"
	err = client.Ping()
	checkError(t, ErrDatabaseFail, err)

	client = NewDefaultRedisClient(RedisAddress, 2)
	err = client.Ping()
	checkError(t, nil, err)
}

func TestRedisRepo_Clean(t *testing.T) {
	client := newMockedRepo()
	mockResponse = "error"
	err := client.Clean()
	checkError(t, ErrDatabaseFail, err)

	mockResponse = "fail"
	err = client.Clean()
	checkError(t, ErrDatabaseFail, err)

	client = NewDefaultRedisClient(RedisAddress, 2)
	err = client.Clean()
	checkError(t, nil, err)
}
