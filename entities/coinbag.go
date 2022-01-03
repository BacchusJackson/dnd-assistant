package entities

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

type CoinBag struct {
	Id       string   `json:"id"`
	Copper   Copper   `json:"copper,string,omitempty"`
	Silver   Silver   `json:"silver,string,omitempty"`
	Gold     Gold     `json:"gold,string,omitempty"`
	Platinum Platinum `json:"platinum,string,omitempty"`
}

func NewCoinBag() *CoinBag {
	return &CoinBag{
		Id:       uuid.NewString(),
		Copper:   0,
		Silver:   0,
		Gold:     0,
		Platinum: 0,
	}
}

func (c CoinBag) EntityId() string {
	return fmt.Sprintf("coinbag.%s", c.Id)
}

func (c CoinBag) Map() map[string]string {
	return map[string]string{
		"id":       c.Id,
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
