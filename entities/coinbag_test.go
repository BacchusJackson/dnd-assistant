package entities

import "testing"

func TestCoinBag_Serialization(t *testing.T) {
	bag := CoinBag{}

	bag.Copper = 10
	bag.Silver = 14
	bag.Gold = 12
	bag.Platinum = 1

	jsonBytes, err := bag.Marshal()
	if err != nil {
		t.Error(err)
	}

	var bag2 CoinBag

	err = bag2.Unmarshal([]byte{})
	if err == nil {
		t.Error("failed to catch json error")
	}

	err = bag2.Unmarshal(jsonBytes)
	if err != nil {
		t.Error(err)
	}

	if bag2.Copper != bag.Copper {
		t.Error("serialization error: Copper")
	}
	if bag2.Silver != bag.Silver {
		t.Error("serialization error: Silver")
	}
	if bag2.Gold != bag.Gold {
		t.Error("serialization error: Gold")
	}
	if bag2.Platinum != bag.Platinum {
		t.Error("serialization error: Platinum")
	}

	t.Log(bag)
}
