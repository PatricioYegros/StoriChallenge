package service

import (
	"testing"
	"time"

	"github.com/patricioyegros/storichallenge/app/models"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	reader := TransactionsReader{}

	fileName := "test.csv"

	tests := []struct {
		name       string
		got        [][]string
		want       []models.Transaction
		err        error
		errMessage string
	}{
		{
			"empty list",
			[][]string{},
			[]models.Transaction{},
			nil,
			"",
		},
		{
			"empty row",
			[][]string{
				{},
			},
			nil,
			ErrReadingTransactions,
			"error parsing line 1: 3 elements expected, got 0",
		},
		{
			"row with less elements",
			[][]string{
				{"asd", "asd"},
			},
			nil,
			ErrReadingTransactions,
			"error parsing line 1: 3 elements expected, got 2",
		},
		{
			"row with incorrect id",
			[][]string{
				{"asd", "7/15", "60.3"},
			},
			nil,
			ErrReadingTransactions,
			"error parsing Id \"asd\" in line 1",
		},
		{
			"row with incorrect date",
			[][]string{
				{"1", "asd", "60.3"},
			},
			nil,
			ErrReadingTransactions,
			"error parsing Date \"asd\" in line 1",
		},
		{
			"row with incorrect transaction",
			[][]string{
				{"1", "7/15", "asd"},
			},
			nil,
			ErrReadingTransactions,
			"error parsing Transaction \"asd\" in line 1",
		},
		{
			"one row",
			[][]string{
				{"0", "7/15", "+60.5"},
			},
			[]models.Transaction{
				{IDInFile: 0, FileName: fileName, Date: time.Date(time.Now().Year(), 7, 15, 0, 0, 0, 0, time.UTC), Amount: decimal.NewFromFloat32(60.5)},
			},
			nil,
			"",
		},
		{
			"correct rows",
			[][]string{
				{"0", "7/15", "+60.5"},
				{"1", "7/28", "-10.3"},
				{"2", "8/2", "-20.46"},
				{"3", "8/13", "+10"},
			},
			[]models.Transaction{
				{IDInFile: 0, FileName: fileName, Date: time.Date(time.Now().Year(), 7, 15, 0, 0, 0, 0, time.UTC), Amount: decimal.NewFromFloat32(60.5)},
				{IDInFile: 1, FileName: fileName, Date: time.Date(time.Now().Year(), 7, 28, 0, 0, 0, 0, time.UTC), Amount: decimal.NewFromFloat32(-10.3)},
				{IDInFile: 2, FileName: fileName, Date: time.Date(time.Now().Year(), 8, 2, 0, 0, 0, 0, time.UTC), Amount: decimal.NewFromFloat32(-20.46)},
				{IDInFile: 3, FileName: fileName, Date: time.Date(time.Now().Year(), 8, 13, 0, 0, 0, 0, time.UTC), Amount: decimal.NewFromInt(10)},
			},
			nil,
			"",
		},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			ans, err := reader.parse(tt.got, fileName)

			require.ErrorIs(t, err, tt.err)

			if err == nil {
				assert.ElementsMatch(t, ans, tt.want)
			} else {
				require.ErrorContains(t, err, tt.errMessage)
			}
		})
	}
}
