package poker

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Table represents a poker table.
type Table struct {

	// Name of the table.
	Name string

	// Stakes of the table.
	Stakes Stakes

	// Size of the table (number of seats).
	Size int

	// Game is the game type.
	Game Game
}

// Stakes represents table stakes.
type Stakes struct {

	// SmallBlind is the size of the small blind.
	SmallBlind Amount

	// BigBlind is the size of the big blind.
	BigBlind Amount
}

// String returns a string representation of the stakes in the form
// '$0.01/$0.02 USD'.
func (s Stakes) String() string {
	return fmt.Sprintf("%v/%v USD", s.SmallBlind, s.BigBlind)
}

// ParseStakes parses a string to a stakes object. The string must be in a format
// similar to '$0.01/$0.02 USD'.
func ParseStakes(stakes string) (Stakes, error) {

	var s Stakes
	var err error

	strs := strings.Split(stakes, "/")
	s.SmallBlind, err = ParseAmount(strs[0])
	if err != nil {
		return Stakes{}, err
	}
	s.BigBlind, err = ParseAmount(strs[1])
	if err != nil {
		return Stakes{}, err
	}
	return s, nil
}

// MarshalJSON marshals the string representation of stakes.
func (s Stakes) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// UnmarshalJSON parses stakes from JSON.
func (s *Stakes) UnmarshalJSON(b []byte) error {

	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	*s, err = ParseStakes(str)
	return err
}

// Internally used game type
type game int

// List of games
const (
	// Texas Hold'em No Limit game
	TexasHoldEmNoLimit game = iota

	// Unknown game
	Unknown
)

// Game is the interface to a value which describes a game.
// The Game interface exists to disallow external code to create new types of
// games.
type Game interface {
	Game() game
	String() string
}

// Game is present to satisfy Game interface.
// The Game interface has this function to disallow external code to create new
// types of games.
func (g game) Game() game {
	return g
}

// String returns the name of the game.
func (g game) String() string {
	switch g {
	case TexasHoldEmNoLimit:
		return "Texas Hold'em No Limit"
	default:
		return "Unknown game"
	}
}

// ParseGame parses a game mode.
func ParseGame(gameStr string) (Game, error) {
	switch gameStr {
	case "Texas Hold'em No Limit":
		return TexasHoldEmNoLimit, nil
	case "No Limit Hold'em":
		return TexasHoldEmNoLimit, nil
	default:
		return Unknown, fmt.Errorf("warning: Failed to parse game")
	}
}

// MarshalJSON marshals the string representation of the game.
func (g game) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

// UnmarshalJSON parses a game from JSON.
func (g *game) UnmarshalJSON(b []byte) error {

	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	gg, err := ParseGame(str)
	*g = gg.(game)
	return err
}
