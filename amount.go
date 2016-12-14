package poker

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Amount represents an amount of US cents.
type Amount int

// String returns a string representing the amount in the form '$2.32'.
func (a Amount) String() string {
	if a == -1 {
		return fmt.Sprintf("All In")
	}
	if a%100 == 0 {
		return fmt.Sprintf("$%.0f", float64(a)/100)
	}
	return fmt.Sprintf("$%.2f", float64(a)/100)
}

// NewAmount creates an amount from a float.
func NewAmount(amount float64) Amount {
	return Amount((amount * 100) + 0.5)
}

// ParseAmount parses an amount from a string.
// The string may contain '$' and 'USD'.
func ParseAmount(amount string) (Amount, error) {
	amount = strings.Replace(amount, "$", "", -1)
	amount = strings.Replace(amount, "USD", "", -1)
	amount = strings.Replace(amount, " ", "", -1)

	val, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println("fail", amount)
		return 0, err
	}

	return NewAmount(val), nil
}

// MarshalJSON marshals the string representation of the amount.
func (a Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

// UnmarshalJSON parses an amount from JSON.
func (a Amount) UnmarshalJSON(data []byte) error {

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	a, err := ParseAmount(str)
	return err
}
