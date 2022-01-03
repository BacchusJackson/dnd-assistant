package repositories

import (
	"dnd-assistant/entities"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"testing"
)

func TestRedisRepo_Write(t *testing.T) {
	repo := NewDefaultRedisRepo(RedisAddress, RedisTestDB)
	repo.redis = NewMockRedisClient()

	// Fail Validation
	item := MockEntity{}
	MockError = entities.ErrInvalidEntity
	err := repo.Write(item)
	checkError(t, entities.ErrInvalidEntity, err)

	// Fail Client
	MockIntCmd.SetErr(redis.ErrClosed)
	MockIntCmd.SetVal(0)
	MockError = nil
	err = repo.Write(item)
	checkError(t, ErrRepoWriteFail, err)

	// Successful Write
	MockMap = map[string]string{"name": "Bob the Destroyer"}
	MockIntCmd.SetErr(nil)
	MockIntCmd.SetVal(int64(len(MockMap)))
	err = repo.Write(item)
	checkError(t, nil, err)
}

func TestRedisRepo_Read(t *testing.T) {
	repo := NewDefaultRedisRepo(RedisAddress, RedisTestDB)
	repo.redis = NewMockRedisClient()

	// Fail Unpack
	_, err := repo.Read("")
	checkError(t, entities.ErrInvalidKey, err)

	// Fail Client
	MockError = nil
	MockStringStringMapCmd.SetErr(redis.ErrClosed)
	_, err = repo.Read(fmt.Sprintf("coin_bag.%s", uuid.NewString()))
	checkError(t, ErrDatabaseFail, err)

	// Default Case
	MockStringStringMapCmd.SetErr(nil)
	_, err = repo.Read(fmt.Sprintf("nothing.%s", uuid.NewString()))
	checkError(t, ErrEntityNotFound, err)

	// Coin Bag
	_, err = repo.Read(fmt.Sprintf("coin_bag.%s", uuid.NewString()))
	checkError(t, entities.ErrInvalidEntity, err)

}
