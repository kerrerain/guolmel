package models

import (
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

const NIL_VALUE = "nil"
const BUDGET_STATE_FLAG = "[BUDGET_STATE]"
const SEPARATOR = ";"
const DATE_LAYOUT = "2006-01-02T15:04:05.000Z"

type Budget struct {
	Expenses               []Expense
	StartDate              time.Time
	EndDate                time.Time
	LastModificationDate   time.Time
	TotalExpenses          decimal.Decimal
	TotalEarnings          decimal.Decimal
	TotalUncheckedExpenses decimal.Decimal
	InitialBalance         decimal.Decimal
	Difference             decimal.Decimal
	CurrentBalance         decimal.Decimal
}

func (b *Budget) ToMessage() string {
	chunks := []string{
		displayDate(b.StartDate),
		displayDate(b.LastModificationDate),
		displayDate(b.EndDate),
		displayDecimal(b.InitialBalance),
		displayDecimal(b.CurrentBalance),
	}
	return strings.Join(chunks, SEPARATOR)
}

func BudgetFromString(str string) *Budget {
	chunks := strings.Split(str, SEPARATOR)

	startDate, _ := time.Parse(DATE_LAYOUT, chunks[0])
	lastModificationDate, _ := time.Parse(DATE_LAYOUT, chunks[1])
	endDate, _ := time.Parse(DATE_LAYOUT, chunks[2])
	initialBalance, _ := decimal.NewFromString(chunks[3])
	currentBalance, _ := decimal.NewFromString(chunks[4])

	return &Budget{
		StartDate:            startDate,
		LastModificationDate: lastModificationDate,
		EndDate:              endDate,
		InitialBalance:       initialBalance,
		CurrentBalance:       currentBalance,
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
