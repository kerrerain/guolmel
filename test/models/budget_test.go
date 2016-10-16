package models_test

import (
	"github.com/magleff/guolmel/models"
	"github.com/shopspring/decimal"
	"testing"
)

func TestToMessage(t *testing.T) {
	// Arrange
	budget := &models.Budget{
		Expenses: []models.Expense{
			models.Expense{
				Amount: decimal.NewFromFloat(0.15),
			},
			models.Expense{
				Amount: decimal.NewFromFloat(-3.15),
			},
		},
		InitialBalance: decimal.NewFromFloat(0.25),
		CurrentBalance: decimal.NewFromFloat(0.25),
	}

	// Act
	actual := budget.ToMessage()

	// Assert
	expected := "nil;nil;nil;0.25;0;0;0;0;0.25|0.15;nil;;false|-3.15;nil;;false"

	if expected != actual {
		t.Error("Expected: " + expected + " Actual: " + actual)
	}
}

func TestBudgetFromString(t *testing.T) {
	// Arrange
	str := "nil;nil;nil;0.25;0;0;0;0;0.25|0.15;nil;;false|-3.15;nil;;true"

	// Act
	budget := models.BudgetFromString(str)

	// Assert
	if len(budget.Expenses) != 2 {
		t.Error("There should be 2 expenses.")
	}

	if decimal.NewFromFloat(0.25).Cmp(budget.InitialBalance) != 0 {
		t.Error("The initial balance should be parsed.")
	}

	if budget.Expenses[1].Checked != true {
		t.Error("Should parse the checked attribute.")
	}
}

func TestComputeInformation(t *testing.T) {
	// Arrange
	budget := &models.Budget{
		InitialBalance: decimal.NewFromFloat(1114.25),
		Expenses: []models.Expense{
			models.Expense{Amount: decimal.NewFromFloat(20.51), Checked: false},
			models.Expense{Amount: decimal.NewFromFloat(-30.68), Checked: true},
			models.Expense{Amount: decimal.NewFromFloat(10.05), Checked: false},
			models.Expense{Amount: decimal.NewFromFloat(-18.36), Checked: false},
		},
	}

	// Act
	budget.ComputeInformation()

	// Assert
	if decimal.NewFromFloat(-49.04).Cmp(budget.TotalExpenses) != 0 {
		t.Error("Should compute the total of expenses.")
	}

	if decimal.NewFromFloat(30.56).Cmp(budget.TotalEarnings) != 0 {
		t.Error("Should compute the total of earnings.")
	}

	if decimal.NewFromFloat(-18.36).Cmp(budget.TotalUncheckedExpenses) != 0 {
		t.Error("Should compute the total of unchecked expenses.")
	}

	if decimal.NewFromFloat(-18.48).Cmp(budget.Difference) != 0 {
		t.Error("Should compute the difference.")
	}

	if decimal.NewFromFloat(1095.77).Cmp(budget.CurrentBalance) != 0 {
		t.Error("Should compute the current balance.")
	}
}
