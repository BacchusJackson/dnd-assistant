package repositories

import (
	"context"
	"dnd-assistant/entities"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
)

var ErrCharacterAlreadyExists = errors.New("character ID already exists")
var ErrDatabaseFail = errors.New("database transaction failed... see log for details")
var ErrCharacterNotFound = errors.New("character id not found")


// RedisRepo TODO: implement character repo interface
type RedisRepo struct {
	client redisClient
	log *log.Logger
}

type redisClient interface {
	Append(key string, value string) error
	Set(key string, value string) error
	SetMap(key string, value map[string]string) error
	Get(key string) (map[string]string, error)
	Ping() error
	Clean() error
}

type defaultClient struct {
	client *redis.Client
}

func NewDefaultClient(address string, database uint) *defaultClient {
	return &defaultClient{
		client:   redis.NewClient(&redis.Options{Addr: address, DB: int(database)}),
	}
}

// Set TODO: Add response code check
func (d defaultClient) Set(key string, value string) error {
	res, err := d.client.HSet(context.Background(), key, value).Result()
	if err != nil {
		return ErrDatabaseFail
	}
	log.Printf("HSET database response: %d\n", res)
	return nil
}

func (d defaultClient) Append(key string, value string) error {
	res, err := d.client.SAdd(context.Background(), key, value).Result()
	if err != nil {
		return ErrDatabaseFail
	}
	if res != 1 {
		log.Printf("SADD database response: %d\n", res)
		return ErrDatabaseFail
	}
	return nil
}

func (d defaultClient) SetMap(key string, value map[string]string) error {
	res, err := d.client.HSet(context.Background(), key, value).Result()
	if err != nil {
		return ErrDatabaseFail
	}
	// res is equal to the number of fields written which should be equal to the length of the map
	if res != int64(len(value)) {
		return ErrDatabaseFail
	}
	return nil
}

func (d defaultClient) Get(key string) (map[string]string, error) {
	res, err := d.client.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, ErrDatabaseFail
	}
	return res, nil
}

func (d defaultClient) Ping() error {
	res, err := d.client.Ping(context.Background()).Result()
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

func (d defaultClient) Clean() error {
	res, err := d.client.FlushDB(context.Background()).Result()
	if err != nil {
		log.Println("failed to clean")
		return ErrDatabaseFail
	}
	if res != "OK" {
		log.Printf("FLUSHDB database response: %s\n", res)
		return ErrDatabaseFail
	}
	return nil
}

func NewRedisRepo(address string, instance uint) (*RedisRepo, error) {
	repo := &RedisRepo{client: NewDefaultClient(address, instance)}
	err := repo.client.Ping()
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (r RedisRepo) CreateCharacter(c *entities.Character) error {
	r.log.Println("Create character")
	err := c.Valid()
	if err != nil {
		r.log.Println(err)
		return ErrDatabaseFail
	}
	c.Ticker = 0

	// Check for duplicate ID
	r.log.Println("Check if character exists")
	characterMap, err := r.client.Get(c.Id)
	if err != nil {
		r.log.Println("Failed to check if character exists")
		return ErrDatabaseFail
	}
	if len(characterMap) > 0 {
		r.log.Printf("Character already exists with ID %s\n", c.Id)
		return ErrCharacterAlreadyExists
	}

	// Write character to database
	r.log.Println("Write character to database")
	err = r.client.SetMap(c.Id, c.Map())
	if err != nil {
		r.log.Printf("failed to write character ID: %s to database", c.Id)
		return ErrDatabaseFail
	}

	// Write history to database
	r.log.Println("Write history to database")
	err = r.client.SetMap(c.HistoryId(), c.Map())
	if err != nil {
		r.log.Printf("failed to write history: %s for character: %s", c.HistoryId(), c.Id)
		return ErrDatabaseFail
	}

	// Write character id character list
	r.log.Println("Write character id to database")
	err = r.client.Append("characters", c.Id)
	if err != nil {
		r.log.Println("failed to write character id to characters")
		return ErrDatabaseFail
	}
	r.log.Println("Completed character creation")
	return nil
}

func (r RedisRepo) GetCharacter(id string) (*entities.Character, error) {
	panic("implement me")
}

func (r RedisRepo) UpdateCharacter(c *entities.Character) error {
	panic("implement me")
}

func (r RedisRepo) DeleteCharacter(id string) error {
	panic("implement me")
}

func (r RedisRepo) RestoreCharacter(id string, ticker uint) (*entities.Character, error) {
	panic("implement me")
}

func (r RedisRepo) CharacterIds() ([]string, error) {
	panic("implement me")
}

