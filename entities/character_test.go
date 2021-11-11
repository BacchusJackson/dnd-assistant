package entities

import (
	"testing"
)

func TestCharacter_Serialization(t *testing.T) {
	hero := NewCharacter("Bruce Wayne", "1")

	jsonBytes, _ := hero.Marshal()

	hero2 := Character{}

	_ = hero2.Unmarshal(jsonBytes)

	if hero.Name != hero2.Name {
		t.Error("failed to marshal/unmarshal")
	}

	hero3 := Character{}

	err := hero3.Unmarshal([]byte("Should throw error"))
	if err == nil {
		t.Error("failed to catch unmarshal error")
	}
}

func TestCharacter_Map(t *testing.T) {
	var hero2 Character
	var hero3 Character

	hero := NewCharacter("Bruce Wayne", "bruce")

	heroMap := hero.Map()

	if hero.Name != heroMap["name"] {
		t.Error("failed to create character map")
	}

	err := hero2.Unmap(heroMap)

	if err != nil {
		t.Errorf("Failed to unmap \n%s\n", err)
	}

	if hero2.Name != hero.Name {
		t.Error("failed to unmap character")
	}

	err = hero3.Unmap(map[string]string{"throw": "error"})

	if err == nil {
		t.Error("failed to catch bad character data")
	}
}

func TestCharacter_HistoryId(t *testing.T) {
	hero := NewCharacter("Bruce Wayne", "bruce")
	hero.Ticker = 2
	if hero.HistoryId() != "bruce:2" {
		t.Error("failed to generate history ID")
	}
}

func TestCharacter_String(t *testing.T) {
	hero := NewCharacter("Bruce Wayne", "bruce")
	hero.Id = "bruce"
	t.Log("\n" + hero.String())
}

func TestAbility_String(t *testing.T) {
	hero := NewCharacter("Bruce Wayne", "bruce")
	hero.Strength = 15
	hero.Charisma = 8
	if hero.Ability(Strength).Name() != "Strength" {
		t.Error("ability name error")
	}
	hero.Dexterity = 15
	if hero.Ability(Dexterity).Score() != 15 {
		t.Error("ability score error")
	}
	if hero.Ability(Strength).String() != "Strength     [15] +2" {
		t.Log(hero.Ability(Strength).String())
		t.Error("failed to create string")
	}
	if hero.Ability(Charisma).String() != "Charisma     [8 ] -1" {
		t.Log(hero.Ability(Charisma).String())
		t.Error("failed to create string")
	}

	if hero.Ability(AbilityName(-1)).Name() != "" {
		t.Error("failed to hit default case in ability")
	}
}

func TestCharacter_Health(t *testing.T) {
	hero := NewCharacter("Bruce Wayne", "bruce")
	hero.MaxHp = 10
	hero.CurrentHp = 5
	if hero.hpPercentage() != 50 {
		t.Error("failed to calculate health percentage")
	}
	hero.CurrentHp = -5

	if hero.hpPercentage() != -50 {
		t.Error("failed to calculate negative health percentage")
	}
}

func TestCharacter_HealthString(t *testing.T) {
	hero := NewCharacter("Bruce Wayne", "bruce")
	hero.MaxHp = 24
	hero.CurrentHp = 12
	t.Log(hero.HealthString())
	if hero.HealthString() != "[=========================                         ] (12/24) 50%" {
		t.Error("failed to generate correct health string")
	}
	hero.CurrentHp = 20
	t.Log(hero.HealthString())

	hero.CurrentHp = 50
	t.Log(hero.HealthString())

	hero.CurrentHp = -30
	t.Log(hero.HealthString())
}

func TestCharacter_ModifyLevel(t *testing.T) {
	hero := NewCharacter("Bruce Wayne", "bruce")
	t.Log(hero.Levels())
	hero.ModifyLevel("Warlock", 2)
	t.Log(hero.Levels())
	hero.ModifyLevel("Barbarian", 3)
	t.Log(hero.Levels())
}
