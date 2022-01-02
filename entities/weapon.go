package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strings"
)

type Weapon struct {
	Name        string   `json:"name"`
	Id          string   `json:"id"`
	Description string   `json:"description"`
	Properties  []string `json:"properties,string,omitempty"`
}

var ErrInvalidWeapon = errors.New("invalid weapon format")

type WeaponProperty string

const (
	Finesse   WeaponProperty = "Finesse"
	Heavy     WeaponProperty = "Heavy"
	Light     WeaponProperty = "Light"
	Loading   WeaponProperty = "Loading"
	Range     WeaponProperty = "Range"
	Reach     WeaponProperty = "Reach"
	Special   WeaponProperty = "Special"
	Thrown    WeaponProperty = "Thrown"
	TwoHanded WeaponProperty = "Two Handed"
	Versatile WeaponProperty = "Versatile"
)

func NewWeapon(name string, description string, props ...WeaponProperty) *Weapon {
	w := &Weapon{Name: name, Description: description}
	w.Id = fmt.Sprintf("weapon.%s", uuid.NewString())
	for _, prop := range props {
		w.Properties = append(w.Properties, string(prop))
	}
	return w
}

func ParseWeapon(m map[string]string) (*Weapon, error) {
	weapon := &Weapon{}

	propJsonBytes := []byte(m["properties"])
	m["properties"] = ""

	jsonBytes, _ := json.Marshal(m)
	err := json.Unmarshal(jsonBytes, weapon)
	err = json.Unmarshal(propJsonBytes, &weapon.Properties)
	return weapon, err
}

func (w Weapon) PropertiesString() string {

	if len(w.Properties) == 0 {
		return "-"
	}
	return strings.Join(w.Properties, ", ")
}

func (w Weapon) Valid() error {
	err := ValidId(w.Id)

	if err != nil {
		log.Println("invalid ID")
		return ErrInvalidWeapon
	}

	return nil
}

func (w Weapon) Map() map[string]string {
	output := map[string]string{}

	// Props have to be converted to json string
	propBytes, _ := json.Marshal(w.Properties)
	jsonBytes, _ := json.Marshal(w)
	_ = json.Unmarshal(jsonBytes, &output)

	output["properties"] = string(propBytes)
	return output
}

func (w Weapon) String() string {

	return strings.Join([]string{w.Name, w.PropertiesString(), w.Description}, " | ")
}
