package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionsPerMonthToEmailData(t *testing.T) {
	tests := []struct {
		name string
		got  []TransactionsPerMonth
		want []transactionsPerMonthEmailData
	}{
		{"empty list", []TransactionsPerMonth{}, []transactionsPerMonthEmailData{}},
		{
			"list with a single transaction in a single month",
			[]TransactionsPerMonth{
				{Month: time.Date(time.Now().Year(), 6, 1, 0, 0, 0, 0, time.UTC), Amount: 1},
			},
			[]transactionsPerMonthEmailData{{Month: "June", Value: "1"}},
		},
		{
			"list with multiple transaction in a single month",
			[]TransactionsPerMonth{
				{Month: time.Date(time.Now().Year(), 6, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			},
			[]transactionsPerMonthEmailData{{Month: "June", Value: "2"}},
		},
		{
			"list with multiple transactions in multiple moths",
			[]TransactionsPerMonth{
				{Month: time.Date(time.Now().Year(), 5, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
				{Month: time.Date(time.Now().Year(), 6, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			},
			[]transactionsPerMonthEmailData{
				{Month: "May", Value: "2"},
				{Month: "June", Value: "2"},
			},
		},
		{
			"list with multiple months of old year",
			[]TransactionsPerMonth{
				{Month: time.Date(1996, 5, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
				{Month: time.Date(1996, 6, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			},
			[]transactionsPerMonthEmailData{
				{Month: "May 1996", Value: "2"},
				{Month: "June 1996", Value: "2"},
			},
		},
		{
			"list with different years",
			[]TransactionsPerMonth{
				{Month: time.Date(time.Now().Year(), 6, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
				{Month: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			},
			[]transactionsPerMonthEmailData{
				{Month: "June", Value: "2"},
				{Month: "June 2023", Value: "2"},
			},
		},
		{
			"list 3 months",
			[]TransactionsPerMonth{
				{Month: time.Date(time.Now().Year(), 5, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
				{Month: time.Date(time.Now().Year(), 6, 1, 0, 0, 0, 0, time.UTC), Amount: 3},
				{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 4},
			},
			[]transactionsPerMonthEmailData{
				{Month: "May", Value: "2"},
				{Month: "June", Value: "3"},
				{Month: "July", Value: "4"},
			},
		},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, transactionsPerMonthToEmailData(tt.got))
		})
	}
}
