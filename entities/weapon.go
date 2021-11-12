package entities

import (
	"encoding/json"
	"github.com/google/uuid"
	"strings"
)

type Weapon struct {
	Name        string           `json:"name"`
	Id          string           `json:"id"`
	Description string           `json:"description"`
	Properties  []WeaponProperty `json:"properties,string,omitempty"`
}

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
	w.Id = uuid.NewString()
	for _, prop := range props {
		w.Properties = append(w.Properties, prop)
	}
	return w
}

func (w Weapon) Marshal() ([]byte, error) {
	return json.Marshal(w)
}

func (w *Weapon) Unmarshal(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, &w)
}

func (w Weapon) PropertiesString() string {
	var props []string
	for _, prop := range w.Properties {
		props = append(props, string(prop))
	}
	if len(props) == 0 {
		return ""
	}
	return strings.Join(props, " | ")
}

func (w Weapon) Valid() error {
	_, err := uuid.Parse(w.Id)
	if err != nil {
		return err
	}

	return nil
}

func (w Weapon) String() string {
	var b strings.Builder
	b.WriteString(w.Name + "\n")
	b.WriteString(w.PropertiesString() + "\n")
	b.WriteString(w.Description + "\n")
	return b.String()
}
