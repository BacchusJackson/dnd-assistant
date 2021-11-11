package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

const (
	Strength AbilityName = iota
	Dexterity
	Constitution
	Intelligence
	Wisdom
	Charisma
)

type AbilityName int

type Ability struct {
	score int
	name  string
}

type Character struct {
	Id               string          `json:"id"`
	Name             string          `json:"name"`
	Race             string          `json:"race"`
	Ticker           uint            `json:"ticker,string"` // Used to log history
	ProficiencyBonus uint            `json:"proficiency_bonus,string"`
	MaxHp            int             `json:"max_hp,string"`
	CurrentHp        int             `json:"current_hp,string"`
	Strength         int             `json:"strength,string"`
	Dexterity        int             `json:"dexterity,string"`
	Constitution     int             `json:"constitution,string"`
	Intelligence     int             `json:"intelligence,string"`
	Wisdom           int             `json:"wisdom,string"`
	Charisma         int             `json:"charisma,string"`
	Athletics        bool            `json:"athletics,string"`
	Acrobatics       bool            `json:"acrobatics,string"`
	SleightOfHand    bool            `json:"sleight_of_hand,string"`
	Stealth          bool            `json:"stealth,string"`
	Arcana           bool            `json:"arcana,string"`
	History          bool            `json:"history,string"`
	Investigation    bool            `json:"investigation,string"`
	Nature           bool            `json:"nature,string"`
	Religion         bool            `json:"religion,string"`
	AnimalHandling   bool            `json:"animal_handling,string"`
	Insight          bool            `json:"insight,string"`
	Medicine         bool            `json:"medicine,string"`
	Perception       bool            `json:"perception,string"`
	Survival         bool            `json:"survival,string"`
	Deception        bool            `json:"deception,string"`
	Intimidation     bool            `json:"intimidation,string"`
	Performance      bool            `json:"performance,string"`
	Persuasion       bool            `json:"persuasion,string"`
	ClassLevels      map[string]uint `json:"character_levels,omitempty"`
}

func (c Character) Marshal() ([]byte, error) {
	// Convert to JSON
	return json.Marshal(c)
}

func (c *Character) Unmarshal(jsonBytes []byte) error {

	err := json.Unmarshal(jsonBytes, &c)
	if err != nil {
		return err
	}
	if c.Id == "" {
		return errors.New("bad character data")
	}
	return nil
}

func (c *Character) Map() map[string]string {
	var m map[string]string

	jsonBytes, _ := c.Marshal()

	_ = json.Unmarshal(jsonBytes, &m)

	return m
}

func (c *Character) Unmap(m map[string]string) error {
	jsonBytes, _ := json.Marshal(m)

	_ = c.Unmarshal(jsonBytes)

	if c.Id == "" {
		return errors.New("bad character data")
	}
	return nil
}

func (c *Character) HistoryId() string {
	return fmt.Sprintf("%s:%d", c.Id, c.Ticker)
}

func (c Character) String() string {
	out := fmt.Sprintf("%s (%s) %s\n", c.Name, c.Id, c.HealthString())
	out += c.Ability(Strength).String() + "\n"
	out += c.Ability(Dexterity).String() + "\n"
	out += c.Ability(Constitution).String() + "\n"
	out += c.Ability(Intelligence).String() + "\n"
	out += c.Ability(Wisdom).String() + "\n"
	out += c.Ability(Charisma).String() + "\n"
	return out
}

func (c Character) HealthString() string {
	percentage := c.hpPercentage()
	ticks := ""

	if percentage > 0 && percentage < 100 {
		ticks = strings.Repeat("=", int(math.Round(float64(percentage/2))))
	}
	if percentage >= 100 {
		ticks = strings.Repeat("=", 49)
		ticks += "+"
	}

	return fmt.Sprintf("[%-50s] (%d/%d) %d%%", ticks, c.CurrentHp, c.MaxHp, percentage)
}

func (c Character) hpPercentage() int {
	return int(math.Round((float64(c.CurrentHp) / float64(c.MaxHp)) * 100.00))
}

func (c Character) Ability(a AbilityName) Ability {
	switch a {
	case Strength:
		return Ability{name: "Strength", score: c.Strength}
	case Dexterity:
		return Ability{name: "Dexterity", score: c.Dexterity}
	case Constitution:
		return Ability{name: "Constitution", score: c.Constitution}
	case Intelligence:
		return Ability{name: "Intelligence", score: c.Intelligence}
	case Wisdom:
		return Ability{name: "Wisdom", score: c.Wisdom}
	case Charisma:
		return Ability{name: "Charisma", score: c.Charisma}
	default:
		return Ability{}
	}
}

func (a Ability) Modifier() int {
	return (a.score - 10) / 2
}

func (a Ability) Name() string {
	return a.name
}

func (a Ability) Score() int {
	return a.score
}

func (a Ability) String() string {
	sign := "+"
	if a.Modifier() == 0 {
		sign = " "
	}
	if a.Modifier() < 0 {
		sign = ""
	}
	return fmt.Sprintf("%-12s [%-2d] %s%d", a.name, a.score, sign, a.Modifier())
}

func (c *Character) ModifyLevel(class string, level uint) {

	c.ClassLevels[class] = level
}

func (c Character) Levels() string {
	if len(c.ClassLevels) == 0 {
		return "None"
	}
	out := "\n"
	maxLength := 0
	for _, key := range reflect.ValueOf(c.ClassLevels).MapKeys() {
		if len(key.String()) > maxLength {
			maxLength = len(key.String())
		}
	}
	for class, level := range c.ClassLevels {
		out += class + strings.Repeat(" ", maxLength-len(class)+1) + strconv.Itoa(int(level)) + "\n"
	}

	return out
}

func NewCharacter(name string, id string) *Character {
	return &Character{
		Id:               id,
		Name:             name,
		Ticker:           0,
		ProficiencyBonus: 0,
		MaxHp:            5,
		CurrentHp:        5,
		Strength:         10,
		Dexterity:        10,
		Constitution:     10,
		Intelligence:     10,
		Wisdom:           10,
		Charisma:         10,
		Athletics:        false,
		Acrobatics:       false,
		SleightOfHand:    false,
		Stealth:          false,
		Arcana:           false,
		History:          false,
		Investigation:    false,
		Nature:           false,
		Religion:         false,
		AnimalHandling:   false,
		Insight:          false,
		Medicine:         false,
		Perception:       false,
		Survival:         false,
		Deception:        false,
		Intimidation:     false,
		Performance:      false,
		Persuasion:       false,
		ClassLevels:      map[string]uint{},
	}
}
