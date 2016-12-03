package card

import (
	"fmt"
	"strings"
)

type card int

type Card interface {
	Card() card
	String() string
}

func (c card) Card() card {
	return c
}

func (c card) String() string {
	switch c {
	case Card2c:
		return "2c"
	case Card3c:
		return "3c"
	case Card4c:
		return "4c"
	case Card5c:
		return "5c"
	case Card6c:
		return "6c"
	case Card7c:
		return "7c"
	case Card8c:
		return "8c"
	case Card9c:
		return "9c"
	case CardTc:
		return "Tc"
	case CardJc:
		return "Jc"
	case CardQc:
		return "Qc"
	case CardKc:
		return "Kc"
	case CardAc:
		return "Ac"

	// Diamonds
	case Card2d:
		return "2d"
	case Card3d:
		return "3d"
	case Card4d:
		return "4d"
	case Card5d:
		return "5d"
	case Card6d:
		return "6d"
	case Card7d:
		return "7d"
	case Card8d:
		return "8d"
	case Card9d:
		return "9d"
	case CardTd:
		return "Td"
	case CardJd:
		return "Jd"
	case CardQd:
		return "Qd"
	case CardKd:
		return "Kd"
	case CardAd:
		return "Ad"

	// Hearts
	case Card2h:
		return "2h"
	case Card3h:
		return "3h"
	case Card4h:
		return "4h"
	case Card5h:
		return "5h"
	case Card6h:
		return "6h"
	case Card7h:
		return "7h"
	case Card8h:
		return "8h"
	case Card9h:
		return "9h"
	case CardTh:
		return "Th"
	case CardJh:
		return "Jh"
	case CardQh:
		return "Qh"
	case CardKh:
		return "Kh"
	case CardAh:
		return "Ah"

	// Spades
	case Card2s:
		return "2s"
	case Card3s:
		return "3s"
	case Card4s:
		return "4s"
	case Card5s:
		return "5s"
	case Card6s:
		return "6s"
	case Card7s:
		return "7s"
	case Card8s:
		return "8s"
	case Card9s:
		return "9s"
	case CardTs:
		return "Ts"
	case CardJs:
		return "Js"
	case CardQs:
		return "Qs"
	case CardKs:
		return "Ks"
	case CardAs:
		return "As"
	}

	return "Invalid"
}

func ParseCard(str string) (Card, error) {

	// Expect 2 cards.
	if len(str) != 2 {
		return CardInvalid, fmt.Errorf("failed to parse card: expected len=2, got %v",
			len(str))
	}

	str = strings.ToLower(str)

	var value int
	var color int

	switch int(str[0]) {
	case '2', '3', '4', '5', '6', '7', '8', '9':
		value = int(str[0]) - int('0')
	case 't':
		value = 10
	case 'j':
		value = 11
	case 'q':
		value = 12
	case 'k':
		value = 13
	case 'a':
		value = 14
	default:
		return CardInvalid, fmt.Errorf("failed to parse card: %v", str)
	}

	switch int(str[1]) {
	case 'c':
		color = 0
	case 'd':
		color = 1
	case 'h':
		color = 2
	case 's':
		color = 3
	default:
		return CardInvalid, fmt.Errorf("failed to parse card: %v", str)
	}

	fmt.Println(value, color, (value-2)+(14*color))
	return card((value - 2) + (13 * color)), nil
}

const (
	CardInvalid card = iota - 1

	// Clubs
	Card2c
	Card3c
	Card4c
	Card5c
	Card6c
	Card7c
	Card8c
	Card9c
	CardTc
	CardJc
	CardQc
	CardKc
	CardAc

	// Diamonds
	Card2d
	Card3d
	Card4d
	Card5d
	Card6d
	Card7d
	Card8d
	Card9d
	CardTd
	CardJd
	CardQd
	CardKd
	CardAd

	// Hearts
	Card2h
	Card3h
	Card4h
	Card5h
	Card6h
	Card7h
	Card8h
	Card9h
	CardTh
	CardJh
	CardQh
	CardKh
	CardAh

	// Spades
	Card2s
	Card3s
	Card4s
	Card5s
	Card6s
	Card7s
	Card8s
	Card9s
	CardTs
	CardJs
	CardQs
	CardKs
	CardAs
)
