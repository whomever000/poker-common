package poker

import "fmt"

// Action is the interface to a player action.
type Action interface {

	// String returns string representation of the action.
	String() string

	// Amount returns the associated amount (not relevant to all action types).
	Amount() Amount
}

// Fold action /////////////////////////////////////////////////////////////////

// NewFoldAction creates a new fold action.
func NewFoldAction() Action {
	return &foldAction{}
}

type foldAction struct{}

func (a *foldAction) Amount() Amount {
	return 0
}
func (a *foldAction) String() string {
	return "fold"
}

// Check action ////////////////////////////////////////////////////////////////

// NewCheckAction creates a new check action.
func NewCheckAction() Action {
	return &checkAction{}
}

type checkAction struct{}

func (a *checkAction) Amount() Amount {
	return 0
}
func (a *checkAction) String() string {
	return "check"
}

// Call action /////////////////////////////////////////////////////////////////

// NewCallAction creates a new call action.
func NewCallAction(amount Amount) Action {
	return &callAction{amount}
}

type callAction struct {
	amount Amount
}

func (a *callAction) Amount() Amount {
	return a.amount
}
func (a *callAction) String() string {
	return fmt.Sprintf("call (%v)", a.amount)
}

// Raise action ////////////////////////////////////////////////////////////////

// NewRaiseAction creates a new raise action.
func NewRaiseAction(amount Amount) Action {
	return &raiseAction{amount}
}

type raiseAction struct {
	amount Amount
}

func (a *raiseAction) Amount() Amount {
	return a.amount
}
func (a *raiseAction) String() string {
	return fmt.Sprintf("raise (%v)", a.amount)
}

// Bet action //////////////////////////////////////////////////////////////////

// NewBetAction creates a new bet action.
func NewBetAction(amount Amount) Action {
	return &betAction{amount}
}

type betAction struct {
	amount Amount
}

func (a *betAction) Amount() Amount {
	return a.amount
}
func (a *betAction) String() string {
	return fmt.Sprintf("bet (%v)", a.amount)
}
