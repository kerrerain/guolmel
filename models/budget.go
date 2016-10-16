package models

import (
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
	"time"
)

const NIL_VALUE = "nil"
const BUDGET_STATE_FLAG = "[BUDGET_STATE]"
const SEPARATOR = ";"
const LINE_BREAK = "|"
const DATE_LAYOUT = "2006-01-02T15:04:05.000Z"

type Budget struct {
	Expenses               []Expense
	StartDate              time.Time
	EndDate                time.Time
	LastModificationDate   time.Time
	InitialBalance         decimal.Decimal
	TotalExpenses          decimal.Decimal
	TotalEarnings          decimal.Decimal
	TotalUncheckedExpenses decimal.Decimal
	Difference             decimal.Decimal
	CurrentBalance         decimal.Decimal
}

func (b *Budget) ToMessage() string {
	chunks := []string{
		displayDate(b.StartDate),
		displayDate(b.LastModificationDate),
		displayDate(b.EndDate),
		displayDecimal(b.InitialBalance),
		displayDecimal(b.TotalExpenses),
		displayDecimal(b.TotalEarnings),
		displayDecimal(b.TotalUncheckedExpenses),
		displayDecimal(b.Difference),
		displayDecimal(b.CurrentBalance),
	}
	return strings.Join(chunks, SEPARATOR) +
		displayExpenses(b.Expenses)
}

func (b *Budget) AddExpense(expense Expense) {
	b.Expenses = append(b.Expenses, expense)
	b.ComputeInformation()
}

func (b *Budget) ComputeInformation() {
	b.TotalEarnings = ComputeTotalEarnings(b.Expenses)
	b.TotalExpenses = ComputeTotalExpenses(b.Expenses)
	b.TotalUncheckedExpenses = ComputeTotalUncheckedExpenses(b.Expenses)
	b.Difference = b.TotalEarnings.Add(b.TotalExpenses)
	b.CurrentBalance = b.InitialBalance.Add(b.Difference)
}

func ComputeTotalEarnings(expenses []Expense) decimal.Decimal {
	totalEarnings := decimal.NewFromFloat(0.00)
	for _, entry := range expenses {
		if entry.Amount.Cmp(decimal.NewFromFloat(0)) > 0 {
			totalEarnings = totalEarnings.Add(entry.Amount)
		}
	}
	return totalEarnings
}

func ComputeTotalExpenses(expenses []Expense) decimal.Decimal {
	totalExpenses := decimal.NewFromFloat(0.00)
	for _, entry := range expenses {
		if entry.Amount.Cmp(decimal.NewFromFloat(0)) <= 0 {
			totalExpenses = totalExpenses.Add(entry.Amount)
		}
	}
	return totalExpenses
}

func ComputeTotalUncheckedExpenses(expenses []Expense) decimal.Decimal {
	totalUncheckedExpenses := decimal.NewFromFloat(0.00)
	filteredExpenses := Filter(expenses, func(obj Expense) bool {
		return !obj.Checked
	})
	for _, entry := range filteredExpenses {
		if entry.Amount.Cmp(decimal.NewFromFloat(0)) <= 0 {
			totalUncheckedExpenses = totalUncheckedExpenses.Add(entry.Amount)
		}
	}
	return totalUncheckedExpenses
}

func Filter(vs []Expense, f func(Expense) bool) []Expense {
	vsf := make([]Expense, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func BudgetFromString(str string) *Budget {
	lines := strings.Split(str, LINE_BREAK)
	chunks := strings.Split(lines[0], SEPARATOR)

	startDate, _ := time.Parse(DATE_LAYOUT, chunks[0])
	lastModificationDate, _ := time.Parse(DATE_LAYOUT, chunks[1])
	endDate, _ := time.Parse(DATE_LAYOUT, chunks[2])
	initialBalance, _ := decimal.NewFromString(chunks[3])
	totalExpenses, _ := decimal.NewFromString(chunks[4])
	totalEarnings, _ := decimal.NewFromString(chunks[5])
	totalUncheckedExpenses, _ := decimal.NewFromString(chunks[6])
	difference, _ := decimal.NewFromString(chunks[7])
	currentBalance, _ := decimal.NewFromString(chunks[8])

	var expenses []Expense

	if len(lines) > 1 {
		expenses, _ = parseExpensesFromString(lines[1:])
	}

	return &Budget{
		StartDate:              startDate,
		LastModificationDate:   lastModificationDate,
		EndDate:                endDate,
		InitialBalance:         initialBalance,
		TotalExpenses:          totalExpenses,
		TotalEarnings:          totalEarnings,
		TotalUncheckedExpenses: totalUncheckedExpenses,
		Difference:             difference,
		CurrentBalance:         currentBalance,
		Expenses:               expenses,
	}
}

func parseExpensesFromString(chunks []string) ([]Expense, error) {
	expenses := []Expense{}

	for _, chunk := range chunks {
		expenses = append(expenses, parseSingleExpense(chunk))
	}

	return expenses, nil
}

func parseSingleExpense(input string) Expense {
	chunks := strings.Split(input, SEPARATOR)

	amount, _ := decimal.NewFromString(chunks[0])
	date, _ := time.Parse(DATE_LAYOUT, chunks[1])
	description := chunks[2]
	checked, _ := strconv.ParseBool(chunks[3])

	return Expense{
		Amount:      amount,
		Date:        date,
		Description: description,
		Checked:     checked,
	}
}

func displayDate(date time.Time) string {
	if date.IsZero() {
		return NIL_VALUE
	} else {
		return date.Format(DATE_LAYOUT)
	}
}

func displayDecimal(dec decimal.Decimal) string {
	return dec.String()
}

func displayExpenses(expenses []Expense) string {
	chunks := []string{}

	if len(expenses) == 0 {
		return ""
	}

	for _, expense := range expenses {
		chunks = append(chunks, displaySingleExpense(expense))
	}

	return LINE_BREAK + strings.Join(chunks, LINE_BREAK)
}

func displaySingleExpense(expense Expense) string {
	chunks := []string{
		displayDecimal(expense.Amount),
		displayDate(expense.Date),
		expense.Description,
		strconv.FormatBool(expense.Checked),
	}
	return strings.Join(chunks, SEPARATOR)
}
