package entities

import (
	"github.com/google/uuid"
	"testing"
)

func TestCoinBag(t *testing.T) {
	bag := NewCoinBag()
	bag.Copper = Copper(30)
	bag.Silver = Silver(23)
	bag.Gold = Gold(12)
	bag.Platinum = Platinum(3)

	t.Log("Bag Key:", bag.EntityKey())
	t.Log(bag)
}

func TestCoinBag_Map(t *testing.T) {
	bag := NewCoinBag()
	bag.Copper = Copper(2)
	bag.Silver = Silver(32)
	bag.Gold = Gold(54)
	bag.Platinum = Platinum(4)
	t.Log(bag.Map())
}

func TestParseCoinBag(t *testing.T) {
	bagMap := map[string]string{
		"id":       uuid.NewString(),
		"copper":   "10",
		"silver":   "11",
		"gold":     "12",
		"platinum": "13",
	}

	// Invalid map
	bagMap["copper"] = "zzz"
	_, err := ParseCoinBag(bagMap)
	checkError(t, ErrInvalidEntity, err)

	// Good Map
	bagMap["copper"] = "10"
	_, err = ParseCoinBag(bagMap)
	checkError(t, nil, err)
}

func TestCoinBag_Validate(t *testing.T) {
	bag := NewCoinBag()
	bag.id = ""
	err := bag.Validate()
	checkError(t, ErrInvalidEntity, err)
}
