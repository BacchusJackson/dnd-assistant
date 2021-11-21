package repositories

import (
	"dnd-assistant/entities"
	"fmt"
	"log"
	"os"
	"testing"
)

const RedisAddress string = "localhost:6379"
var repo *RedisRepo
// VoidWriter Used to ignore log output in tests
type VoidWriter struct { }
func (v VoidWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func init() {
	var err error
	repo, err = NewRedisRepo(RedisAddress, 2)
	repo.log = log.New(os.Stdout, "TEST: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		fmt.Println("Tests can't run. Database failed to connect")
		fmt.Println(err)
		os.Exit(-1)
	}
}

func cleanDatabase() {
	_ = repo.client.Clean()
}

func TestRedisRepo_CreateCharacter(t *testing.T) {
	defer cleanDatabase()
	character := entities.NewCharacter("Bruce Wayne", "bw")
	// Attempt invalid character
	character.MaxHp = 0
	err := repo.CreateCharacter(character)
	if err != ErrDatabaseFail {
		t.Error("expected error for invalid character")
	}

	// Attempt valid character
	character.MaxHp = 10
	err = repo.CreateCharacter(character)
	if err != nil {
		t.Error(err)
	}

	// Attempt duplicate character id
	err = repo.CreateCharacter(character)
	if err != ErrCharacterAlreadyExists {
		t.Error("expected error for duplicate character id")
	}

}

