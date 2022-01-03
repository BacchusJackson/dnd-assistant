package entities

import "testing"

func TestCoinBag(t *testing.T) {
	bag := NewCoinBag()
	bag.Copper = Copper(30)
	bag.Silver = Silver(23)
	bag.Gold = Gold(12)
	bag.Platinum = Platinum(3)

	t.Log("Bag ID:", bag.EntityId())
	t.Log(bag)
}
