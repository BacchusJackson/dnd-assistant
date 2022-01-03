package entities

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

type CoinBag struct {
	id       string   `json:"id"`
	Copper   Copper   `json:"copper,string,omitempty"`
	Silver   Silver   `json:"silver,string,omitempty"`
	Gold     Gold     `json:"gold,string,omitempty"`
	Platinum Platinum `json:"platinum,string,omitempty"`
}

func NewCoinBag() *CoinBag {
	return &CoinBag{
		id:       uuid.NewString(),
		Copper:   0,
		Silver:   0,
		Gold:     0,
		Platinum: 0,
	}
}

func (c CoinBag) Validate() error {
	err := Validate(c.EntityKey())
	if err != nil {
		return ErrInvalidEntity
	}
	return nil
}

func ParseCoinBag(m map[string]string) (*CoinBag, error) {

	bag := &CoinBag{}
	bag.id = m["id"]
	values := []int{0, 0, 0, 0}
	for i, coin := range []string{"copper", "silver", "gold", "platinum"} {
		value, err := strconv.Atoi(m[coin])
		if err != nil {
			return nil, ErrInvalidEntity
		}
		values[i] = value
	}

	bag.Copper = Copper(values[0])
	bag.Silver = Silver(values[1])
	bag.Gold = Gold(values[2])
	bag.Platinum = Platinum(values[3])
	return bag, bag.Validate()
}

func (c CoinBag) EntityKey() string {
	return fmt.Sprintf("coinbag.%s", c.id)
}

func (c CoinBag) Map() map[string]string {
	return map[string]string{
		"id":       c.id,
		"copper":   strconv.Itoa(int(c.Copper)),
		"silver":   strconv.Itoa(int(c.Silver)),
		"gold":     strconv.Itoa(int(c.Gold)),
		"platinum": strconv.Itoa(int(c.Platinum)),
	}
}

func (c CoinBag) String() string {
	return fmt.Sprintf("Copper: %d | Silver: %d | Gold: %d | Platinum %d",
		c.Copper, c.Silver, c.Gold, c.Platinum)
}
