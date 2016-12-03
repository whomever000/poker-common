package poker

import "testing"

// NewAmount() /////////////////////////////////////////////////////////////////

type testPairNewAmount struct {
	input  float64
	output int
}

var testsNewAmount = []testPairNewAmount{
	{0.42, 42},
	{42, 4200},
	{-0.42, -42},
}

func TestNewAmount(t *testing.T) {

	tests := testsNewAmount

	for i := 0; i < len(tests); i++ {
		amount := NewAmount(tests[i].input)
		if int(amount) != tests[i].output {
			t.Errorf("For %v expected %v, got %v",
				tests[i].input,
				tests[i].output,
				int(amount))
		}
	}
}

// String() ////////////////////////////////////////////////////////////////////

type testPairAmountString struct {
	input  float64
	output string
}

var testsAmountString = []testPairAmountString{
	{0.42, "$0.42"},
	{42, "$42.00"},
	{-0.42, "$-0.42"},
}

func TestAmountString(t *testing.T) {

	tests := testsAmountString

	for i := 0; i < len(tests); i++ {
		amount := NewAmount(tests[i].input)
		if amount.String() != tests[i].output {
			t.Errorf("For %v expected %v, got %v",
				tests[i].input,
				tests[i].output,
				amount.String())
		}
	}
}

// ParseAmount() ///////////////////////////////////////////////////////////////

type testPairParseAmount struct {
	input  string
	output int
}

var testsParseAmount = []testPairParseAmount{
	{"$0.42", 42},
	{"$42.00", 4200},
	{"$-0.42", -42},
	{"$123 USD", 12300},
}

var testsParseAmountError = []string{
	"$1.a3",
	"asdasd",
	"12346a",
}

func TestParseAmount(t *testing.T) {

	tests := testsParseAmount

	for i := 0; i < len(tests); i++ {
		amount, err := ParseAmount(tests[i].input)
		if err != nil || int(amount) != tests[i].output {
			t.Errorf("For %v expected %v, got %v",
				tests[i].input,
				tests[i].output,
				int(amount))
		}
	}

	// Additional test for error cases
	testsError := testsParseAmountError

	for i := 0; i < len(testsError); i++ {
		amount, err := ParseAmount(testsError[i])
		if err == nil {
			t.Errorf("For %v expected error, got no error and amount of %v",
				testsError[i], amount)
		}
		if int(amount) != 0 {
			t.Errorf("For %v expected $0.00, got %v", testsError[i], amount)
		}
	}
}
