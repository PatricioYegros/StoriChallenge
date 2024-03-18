package models

import (
	"github.com/shopspring/decimal"
)

type User struct {
	Email   string
	Balance decimal.Decimal
}
