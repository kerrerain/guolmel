package cmd_test

import (
	"github.com/magleff/guolmel/cmd"
	"github.com/shopspring/decimal"
	"testing"
)

func TestParseAmount(t *testing.T) {
	// Arrange
	amount := "0.345"
	args := []string{amount}

	// Act
	expense, err := cmd.ParseExpenseFromArguments(args)

	// Assert
	if err != nil {
		t.Error("The given string " + amount + " should not raise an error")
	}

	if decimal.NewFromFloat(-0.34).Cmp(expense.Amount) != 0 {
		t.Error("Should parse a string and truncate to 2 decimals.")
	}
}

func TestParsePositiveAmount(t *testing.T) {
	// Arrange
	amount := "+0.345"
	args := []string{amount}

	// Act
	expense, err := cmd.ParseExpenseFromArguments(args)

	// Assert
	if err != nil {
		t.Error("The given string " + amount + " should not raise an error")
	}

	if decimal.NewFromFloat(0.34).Cmp(expense.Amount) != 0 {
		t.Error("Should parse a positive amount and truncate to 2 decimals.")
	}
}

func TestParseInvalidAmount(t *testing.T) {
	// Arrange
	amount := "aAnjbssvg"
	args := []string{amount}

	// Act
	_, err := cmd.ParseExpenseFromArguments(args)

	// Assert
	if err == nil {
		t.Error("The given string "+amount+" is not a number.",
			"An invalid string should raise an error.")
	}
}

func TestParseCommaAmount(t *testing.T) {
	// Arrange
	amount := "0,3"
	args := []string{amount}

	// Act
	expense, err := cmd.ParseExpenseFromArguments(args)

	// Assert
	if err != nil {
		t.Error("The given string " + amount + " should not raise an error")
	}

	if decimal.NewFromFloat(-0.3).Cmp(expense.Amount) != 0 {
		t.Error("Should parse a string even with a comma.")
	}
}
