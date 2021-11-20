package repositories

import "dnd-assistant/entities"

// RedisRepo TODO: implement character repo interface
type RedisRepo struct { }

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

func (r RedisRepo) CreateCharacter(c *entities.Character) error {
	panic("implement me")
}