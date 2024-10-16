package service

import (
	"sort"
	"time"

	"github.com/elliotchance/pie/v2"
	"github.com/patricioyegros/storichallenge/app/models"
	"github.com/shopspring/decimal"
)

type Service struct {
	TransactionsReader TransactionsReader
	EmailService       IEmailService
}

type TransactionsPerMonth struct {
	Month  time.Time
	Amount int
}

// Process read csv file called csvFileName,
// calculates total balance, number of transactions grouped by month
// and average credit and debit
// and send this information by email to destinationEmail
func (service Service) Process(csvFileName, destinationEmail string) error {
	transactions, err := service.TransactionsReader.Read(csvFileName)
	if err != nil {
		return err
	}

	transactionsBalance := service.CalculateTotalBalance(transactions)

	avgDebit, avgCredit := service.CalculateAverageDebitAndCredit(transactions)

	return service.EmailService.Send(
		transactionsBalance,
		service.CalculateTransactionsPerMonth(transactions),
		avgDebit, avgCredit, destinationEmail,
	)
}

// CalculateTotalBalance calculates the total balance from a list of transactions as the sum of all transactions
func (service Service) CalculateTotalBalance(transactions []models.Transaction) float64 {
	amounts := transactionsToAmounts(transactions)

	var totalAmount float64

	if len(amounts) == 0 {
		return 0
	}

	if len(amounts) == 1 {
		totalAmount, _ = amounts[0].Float64()
		return totalAmount
	}

	totalAmount, _ = decimal.Sum(
		amounts[0],
		amounts[1:]...,
	).Float64()

	return totalAmount
}

// CalculateTransactionsPerMonth calculates the amount of transactions
// for each month present in the list of transactions.
// The returned list is in ascending order by month.
func (service Service) CalculateTransactionsPerMonth(transactions []models.Transaction) []TransactionsPerMonth {
	// group by year and month
	grouped := pie.GroupBy(transactions, func(transaction models.Transaction) time.Time {
		return time.Date(transaction.Date.Year(), transaction.Date.Month(), 1, 0, 0, 0, 0, time.UTC)
	})

	ans := make([]TransactionsPerMonth, 0, len(grouped))

	// count amount of transactions per group
	for month, transactions := range grouped {
		ans = append(ans, TransactionsPerMonth{Month: month, Amount: len(transactions)})
	}

	// order by month
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Month.Compare(ans[j].Month) == -1
	})

	return ans
}

// CalculateAverageDebitAndCredit calculates the average debit and average credit from a list of transactions
func (service Service) CalculateAverageDebitAndCredit(transactions []models.Transaction) (float64, float64) {
	amounts := transactionsToAmounts(transactions)

	debits := pie.Filter(amounts, func(amount decimal.Decimal) bool {
		return amount.LessThan(decimal.Zero)
	})

	credits := pie.Filter(amounts, func(amount decimal.Decimal) bool {
		return amount.GreaterThan(decimal.Zero)
	})

	averageDebit, _ := calculateAverage(debits).Float64()
	averageCredit, _ := calculateAverage(credits).Float64()

	return averageDebit, averageCredit
}

// calculateAverage calculates the average of a list of amount
func calculateAverage(amounts []decimal.Decimal) decimal.Decimal {
	if len(amounts) == 0 {
		return decimal.Zero
	}

	if len(amounts) == 1 {
		return amounts[0]
	}

	return decimal.Avg(amounts[0], amounts[1:]...)
}

// transactionsToAmounts transforms a list of transactions to a list of its amounts
func transactionsToAmounts(transactions []models.Transaction) []decimal.Decimal {
	return pie.Map(transactions, func(transaction models.Transaction) decimal.Decimal {
		return transaction.Amount
	})
}
