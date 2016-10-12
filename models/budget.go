package models

import (
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

const NIL_VALUE = "nil"

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
	return strings.Join(chunks, ";")
}

func displayDate(date time.Time) string {
	if date.IsZero() {
		return NIL_VALUE
	} else {
		return date.String()
	}
}

func displayDecimal(dec decimal.Decimal) string {
	return dec.String()
}
