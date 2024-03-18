package models

import (
	"time"

	"github.com/shopspring/decimal"
)

const AmountBytes = 64

type Transaction struct {
	IDInFile uint
	FileName string
	Date     time.Time
	Amount   decimal.Decimal
}
