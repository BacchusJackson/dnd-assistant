package repositories

import (
	"github.com/go-redis/redis/v8"
	"testing"
)

// Unit Tests
func TestRedisRepo_Append(t *testing.T) {
	client := NewMockRedisClient()

	// Redis Command Failure
	MockIntCmd.SetErr(redis.ErrClosed)
	err := client.Append("error", "")
	checkError(t, ErrDatabaseFail, err)

	// Unexpected Return Code
	MockIntCmd.SetErr(nil)
	MockIntCmd.SetVal(-1)
	err = client.Append("key", "value")
	checkError(t, ErrDatabaseFail, err)

	// Live data test
	client = NewDefaultRedisClient(RedisAddress, 2)
	err = client.Clean()
	checkError(t, nil, err)
	err = client.Append("test-key", "test-value")
	checkError(t, nil, err)
}

func TestRedisRepo_Update(t *testing.T) {
	client := NewMockRedisClient()

	// Redis Command Failure
	MockIntCmd.SetErr(redis.ErrClosed)
	MockIntCmd.SetVal(0)
	err := client.Update("error", "test-field", "test-value")
	checkError(t, ErrDatabaseFail, err)

	// Unexpected Return Code
	MockIntCmd.SetErr(nil)
	MockIntCmd.SetVal(-1)
	err = client.Update("fail", "test-field", "test-value")
	checkError(t, ErrDatabaseFail, err)

	// Live test
	client = NewDefaultRedisClient(RedisAddress, RedisTestDB)
	err = client.Clean()
	checkError(t, nil, err)
	err = client.Update("test-key", "test-field", "test-value")
	checkError(t, nil, err)
}

func TestRedisClient_Write(t *testing.T) {
	client := NewMockRedisClient()

	// Redis Command Failure
	MockIntCmd.SetErr(redis.ErrClosed)
	MockIntCmd.SetVal(0)
	err := client.Write("error", nil)
	checkError(t, ErrDatabaseFail, err)

	// Unexpected Return Code
	MockIntCmd.SetErr(nil)
	MockIntCmd.SetVal(-1)
	err = client.Write("fail", nil)
	checkError(t, ErrDatabaseFail, err)

	// Live Data Test
	client = NewDefaultRedisClient(RedisAddress, RedisTestDB)
	err = client.Clean()
	checkError(t, nil, err)
	err = client.Write("test-key", map[string]string{
		"field-1": "value-1",
		"field-2": "value-2",
	})
	checkError(t, nil, err)
}

func TestRedisClient_Read(t *testing.T) {
	client := NewMockRedisClient()

	// Redis Command Failure
	MockStringStringMapCmd.SetErr(redis.ErrClosed)
	MockStringStringMapCmd.SetVal(nil)
	res, err := client.Read("error")
	checkError(t, ErrDatabaseFail, err)

	// Live Test
	client = NewDefaultRedisClient(RedisAddress, RedisTestDB)
	err = client.Update("test-key", "test-field", "test-value")
	checkError(t, nil, err)

	res, err = client.Read("test-key")
	checkError(t, nil, err)
	checkError(t, "test-value", res["test-field"])

}

func TestRedisClient_Ping(t *testing.T) {
	client := NewMockRedisClient()

	// Redis Command Failure
	MockStatusCmd.SetErr(redis.ErrClosed)
	MockStatusCmd.SetVal("")
	err := client.Ping()
	checkError(t, ErrDatabaseFail, err)

	// Unexpected Return
	MockStatusCmd.SetErr(nil)
	MockStatusCmd.SetVal("Database Down")
	err = client.Ping()
	checkError(t, ErrDatabaseFail, err)

	// Live Test
	client = NewDefaultRedisClient(RedisAddress, RedisTestDB)
	err = client.Ping()
	checkError(t, nil, err)
}

func TestRedisClient_Clean(t *testing.T) {
	client := NewMockRedisClient()

	// Redis Command Failure
	MockStatusCmd.SetErr(redis.ErrClosed)
	MockStatusCmd.SetVal("")
	err := client.Clean()
	checkError(t, ErrDatabaseFail, err)

	// Unexpected Return
	MockStatusCmd.SetErr(nil)
	MockStatusCmd.SetVal("")
	err = client.Clean()
	checkError(t, ErrDatabaseFail, err)

	// Live Test
	client = NewDefaultRedisClient(RedisAddress, RedisTestDB)
	err = client.Clean()
	checkError(t, nil, err)
}
