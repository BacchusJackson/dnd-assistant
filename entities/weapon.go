package entities

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strings"
)

type Weapon struct {
	Name        string           `json:"name"`
	Id          string           `json:"id"`
	Description string           `json:"description"`
	Properties  []WeaponProperty `json:"properties,string,omitempty"`
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
		w.Properties = append(w.Properties, prop)
	}
	return w
}

func (w Weapon) PropertiesString() string {
	var props []string
	for _, prop := range w.Properties {
		props = append(props, string(prop))
	}
	if len(props) == 0 {
		return "-"
	}
	return strings.Join(props, ", ")
}

func (w Weapon) Valid() error {
	err := ValidId(w.Id)

	if err != nil {
		log.Println("invalid ID")
		return ErrInvalidWeapon
	}

	return nil
}

func (w Weapon) String() string {

	return strings.Join([]string{w.Name, w.PropertiesString(), w.Description}, " | ")
}
