package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Expense struct {
	Amount      decimal.Decimal
	Date        time.Time
	Description string
	Checked     bool
}
