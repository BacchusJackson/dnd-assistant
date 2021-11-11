package entities

import (
	"encoding/json"
	"fmt"
)

type CoinBag struct {
	Copper   uint `json:"copper,string,omitempty"`
	Silver   uint `json:"silver,string,omitempty"`
	Gold     uint `json:"gold,string,omitempty"`
	Platinum uint `json:"platinum,string,omitempty"`
}

func (c CoinBag) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c *CoinBag) Unmarshal(jsonBytes []byte) error {
	err := json.Unmarshal(jsonBytes, &c)
	return err
}

func (c CoinBag) string() string {
	return fmt.Sprintf("Copper: %d | Silver: %d | Gold: %d | Platinum %d",
		c.Copper, c.Silver, c.Gold, c.Platinum)
}
