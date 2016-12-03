package poker

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/whomever000/poker-common/card"
)

type Round struct {
	Cards   []card.Card
	Pot     Amount
	Actions []PlayerAction
}

type Result struct {
	Winner    PlayerPosition
	Pot       Amount
	ShowDowns []PlayerCards
}

type Date time.Time

func (d Date) String() string {
	return time.Time(d).Format("2006/01/02 15:04:05 MST")
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

type Hand struct {
	Client     string
	Table      Table
	HandID     int
	Date       Date
	Button     PlayerPosition
	SmallBlind PlayerPosition
	BigBlind   PlayerPosition
	ThisPlayer *PlayerCards
	Players    []Player
	Rounds     []Round
	Result     *Result
}

func (h *Hand) String() string {
	var str string

	// Header
	str += fmt.Sprintf("%v Hand #%v: %v ", h.Client, h.HandID, h.Table.Game)
	str += fmt.Sprintf("(%v) - %v\n", h.Table.Stakes,
		h.Date)
	str += fmt.Sprintf("Table '%v' %v-max Seat #%v is the button\n",
		h.Table.Name, h.Table.Size, h.Button)

	// Player names and stacks
	for i := 0; i < h.Table.Size; i++ {
		str += fmt.Sprintf("Seat %v: %v (%v in chips)\n", i+1, h.Players[i].Name,
			h.Players[i].Stack)
	}

	// Make sure small and big blind have been posted
	if len(h.Rounds) == 0 ||
		len(h.Rounds[0].Actions) < 2 {
		return str
	}

	// Print small and big blind
	smallBlind := h.SmallBlind.Player(h).Name
	bigBlind := h.BigBlind.Player(h).Name

	str += fmt.Sprintf("%v: posts small blind %v\n", smallBlind,
		h.Table.Stakes.SmallBlind)
	str += fmt.Sprintf("%v: posts big blind %v\n", bigBlind,
		h.Table.Stakes.BigBlind)

	// Print betting rounds
	for r := 0; r < len(h.Rounds); r++ {
		var i int
		if r == 0 {
			str += fmt.Sprintf("*** HOLE CARDS ***\n")
		} else if r == 1 {
			str += fmt.Sprintf("*** FLOP *** [%v %v %v]\n", h.Rounds[r].Cards[2],
				h.Rounds[r].Cards[3], h.Rounds[r].Cards[4])
		} else if r == 2 {
			str += fmt.Sprintf("*** TURN *** [%v %v %v][%v]\n", h.Rounds[r].Cards[2],
				h.Rounds[r].Cards[3], h.Rounds[r].Cards[4], h.Rounds[r].Cards[5])
		} else if r == 3 {
			str += fmt.Sprintf("*** RIVER *** [%v %v %v %v][%v]\n", h.Rounds[r].Cards[2],
				h.Rounds[r].Cards[3], h.Rounds[r].Cards[4], h.Rounds[r].Cards[5],
				h.Rounds[r].Cards[6])
		}

		for ; i < len(h.Rounds[r].Actions); i++ {
			player := h.Rounds[r].Actions[i].Position.Player(h).Name
			action := h.Rounds[r].Actions[i].Action

			str += fmt.Sprintf("%v: %v\n", player, action)
		}
	}

	if h.Result.Winner == 0 {
		return str
	}

	// Print showdown
	if len(h.Result.ShowDowns) > 0 {
		str += fmt.Sprintf("*** SHOW DOWN ***\n")
	}

	for i := 0; i < len(h.Result.ShowDowns); i++ {
		player := h.Result.ShowDowns[i].Position.Player(h).Name
		c1 := h.Result.ShowDowns[i].Cards[0]
		c2 := h.Result.ShowDowns[i].Cards[1]

		str += fmt.Sprintf("%v: shows [%v %v]\n", player, c1, c2)
	}

	// Print winnings
	str += fmt.Sprintf("%v collected %v from pot", h.Result.Winner.Player(h).Name,
		h.Result.Pot)

	return str
}
