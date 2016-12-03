package poker

import (
	"encoding/json"
	"strconv"

	"github.com/whomever000/poker-common/card"
)

// Player represents a poker player.
type Player struct {

	// Name is the player's name.
	Name string

	// Stack is the player's stack at hand start.
	Stack Amount
}

// UnmarshalJSON parses a player from JSON.
func (p *Player) UnmarshalJSON(b []byte) error {

	var data struct {
		name  string
		stack string
	}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	p.Name = data.name
	p.Stack, err = ParseAmount(data.stack)
	return err
}

// PlayerPosition is a position of a player.
// PlayerPositions are 1-indexed, and a value of 0 means 'no player'.
type PlayerPosition int

// Player returns the player given his position.
// h is the hand in question.
func (p PlayerPosition) Player(h *Hand) *Player {

	// Perform boundary check.
	if p < 1 || int(p) > len(h.Players) {
		return nil
	}

	return &h.Players[p-1]
}

// NextPlayerPosition returns the next player position at the table.
func NextPlayerPosition(pos PlayerPosition, tableSize int) PlayerPosition {
	newPos := PlayerPosition(pos + 1)
	if int(newPos) > tableSize {
		newPos = PlayerPosition(1)
	}

	return newPos
}

// PreviousPlayerPosition returns the previous player position at the table.
func PreviousPlayerPosition(pos PlayerPosition, tableSize int) PlayerPosition {
	newPos := PlayerPosition(pos - 1)
	if int(newPos) < 1 {
		newPos = PlayerPosition(tableSize)
	}

	return newPos
}

// PlayerAction is a player's action.
type PlayerAction struct {

	// Position is the player's position.
	Position PlayerPosition

	// Action is the player's action.
	Action Action
}

// UnmarshalJSON parses a player action from JSON.
func (p *PlayerAction) UnmarshalJSON(b []byte) error {

	var data struct {
		position string
		action   string
	}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	val, err := strconv.ParseInt(data.position, 10, 64)
	if err != nil {
		return err
	}
	p.Position = PlayerPosition(val)

	err = json.Unmarshal(b, &p.Action)
	return err
}

// PlayerCards is a player's cards.
type PlayerCards struct {

	// Position is the player's position.
	Position PlayerPosition

	// Cards are the player's cards.
	Cards []card.Card
}

// UnmarshalJSON parses player cards from JSON.
func (p *PlayerCards) UnmarshalJSON(b []byte) error {

	var data struct {
		position string
		action   string
	}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	val, err := strconv.ParseInt(data.position, 10, 64)
	if err != nil {
		return err
	}
	p.Position = PlayerPosition(val)

	err = json.Unmarshal(b, &p.Cards)
	return err
}
