package repositories

import (
	"dnd-assistant/entities"
	"log"
)

type RedisRepo struct {
	redis Client
}

func NewDefaultRedisRepo(address string, database uint) *RedisRepo {
	return &RedisRepo{redis: NewDefaultRedisClient(address, database)}
}

func (r RedisRepo) Write(e entities.Entity) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	err = r.redis.Write(e.EntityKey(), e.Map())
	if err != nil {
		log.Println(err)
		return ErrRepoWriteFail
	}
	return nil
}

func (r RedisRepo) Read(key string) (entities.Entity, error) {
	eType, _, err := entities.Unpack(key)

	if err != nil {
		return nil, err
	}

	entityMap, err := r.redis.Read(key)
	if err != nil {
		return nil, err
	}

	switch eType {
	case "coin_bag":
		return entities.ParseCoinBag(entityMap)
	default:
		return nil, ErrEntityNotFound
	}

}
