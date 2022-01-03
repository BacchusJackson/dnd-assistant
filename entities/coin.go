package entities

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Coin interface {
	Value() (copper uint, silver uint, gold uint, platinum uint)
	String() string
}

type Copper uint

func (c Copper) Value() (copper uint, silver uint, gold uint, platinum uint) {
	return uint(c), uint(c / 10), uint(c / 100), uint(c / 10000)
}

func (c Copper) String() string {
	msg := message.NewPrinter(language.AmericanEnglish)
	return msg.Sprintf("%d CP\n", c)
}

type Silver uint

func (s Silver) Value() (copper uint, silver uint, gold uint, platinum uint) {
	return uint(s * 10), uint(s), uint(s / 10), uint(s / 1000)
}

func (s Silver) String() string {
	msg := message.NewPrinter(language.AmericanEnglish)
	return msg.Sprintf("%d SP\n", s)
}

type Gold uint

func (g Gold) Value() (copper uint, silver uint, gold uint, platinum uint) {
	return uint(g * 100), uint(g * 10), uint(g), uint(g / 100)
}

func (g Gold) String() string {
	msg := message.NewPrinter(language.AmericanEnglish)
	return msg.Sprintf("%d GP\n", g)
}

type Platinum uint

func (p Platinum) Value() (copper uint, silver uint, gold uint, platinum uint) {
	return uint(p * 10000), uint(p * 1000), uint(p * 100), uint(p)
}

func (p Platinum) String() string {
	msg := message.NewPrinter(language.AmericanEnglish)
	return msg.Sprintf("%d PP\n", p)
}
