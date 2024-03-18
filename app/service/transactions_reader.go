package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/patricioyegros/storichallenge/app/funcs"
	"github.com/patricioyegros/storichallenge/app/models"

	"github.com/shopspring/decimal"
)

const (
	IDIndex                = 0
	IDHeader               = "Id"
	DateIndex              = 1
	DateHeader             = "Date"
	AmountIndex            = 2
	AmountHeader           = "Transaction"
	numberOfElementsPerRow = 3

	receivedDateSeparator = "/"
	wantDateSeparator     = "-"
	dateLayout            = "06/20/2024"
	dayAndMonthLen        = 2
)

type TransactionsReader struct {
	CSVReader funcs.FileReader
}

var ErrReadingTransactions = errors.New("error parsing transactions csv")
var ErrParsingID = errors.New("error converting transaction id")
var ErrParsingDate = errors.New("error converting transaction date")
var ErrParsingAmount = errors.New("error converting transaction amount to decimal")

// Read reads a CSV file that contains a list of transactions and
// returns the list of transactions of the CSV file
// or ErrReadingTransactions if an error is produced
func (reader TransactionsReader) Read(csvFileName string) ([]models.Transaction, error) {
	var csvReader funcs.FileReader
	csvReader = reader.CSVReader
	csvRows, err := csvReader.Read(csvFileName)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrReadingTransactions, err.Error())
	}

	return reader.parse(csvRows, csvFileName)
}

// Parse the list of rows of the csv file into a list of models.Transaction,
// converting each cell of the row in the correct type.
// Returns the list of transactions or ErrParsingCsv if there is an error during type conversion
func (reader TransactionsReader) parse(csvRows [][]string, csvFileName string) ([]models.Transaction, error) {
	// create list with the same size as csvRows
	transactions := make([]models.Transaction, 0, len(csvRows))

	// transform each row into a models.Transaction
	for lineNumber, row := range csvRows {
		if len(row) != numberOfElementsPerRow {
			return nil, fmt.Errorf("%w: error parsing line %d: %d elements expected, got %d", ErrReadingTransactions, lineNumber+1, numberOfElementsPerRow, len(row))
		}

		idString := row[IDIndex]

		id, err := parseTransactionID(idString)
		if err != nil {
			return nil, errorParsingCSV(IDHeader, idString, lineNumber)
		}

		dateString := row[DateIndex]

		date, err := parseDate(dateString)
		if err != nil {
			return nil, errorParsingCSV(DateHeader, dateString, lineNumber)
		}

		amountString := row[AmountIndex]

		amount, err := parseAmount(amountString)
		if err != nil {
			return nil, errorParsingCSV(AmountHeader, amountString, lineNumber)
		}

		transactions = append(transactions, models.Transaction{
			IDInFile: id,
			FileName: csvFileName,
			Date:     date,
			Amount:   amount,
		})
	}

	return transactions, nil
}

// Returns a ErrParsingCsv with more information for the user
func errorParsingCSV(header string, stringValue string, lineNumber int) error {
	return fmt.Errorf(
		"%w: error parsing %s %q in line %d",
		ErrReadingTransactions, header, stringValue, lineNumber+1,
	)
}

// parseTransactionID transforms a string representing an id to uint
// Returns ErrParsingID if the transformation if not possible
func parseTransactionID(idString string) (uint, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrParsingID, err.Error())
	}

	if id < 0 {
		return 0, fmt.Errorf("%w: %s", ErrParsingID, "value cannot be less than 0")
	}

	return uint(id), nil
}

// parseDate transforms a string in format "7/15" or "7/15/2023" to a time.Time
// Returns ErrParsingDate if the transformation if not possible
func parseDate(dateString string) (time.Time, error) {
	dateSplitted := strings.Split(dateString, receivedDateSeparator)
	if len(dateSplitted) < dayAndMonthLen {
		return time.Time{}, fmt.Errorf("%w: at least month and year are expected", ErrParsingDate)
	}

	// add a 0 to the begging of the date to ensure it has two digits in case month is < 10
	for i, datePart := range dateSplitted {
		if len(datePart) == 1 {
			dateSplitted[i] = "0" + datePart
		}
	}

	// if dateString doesn't have the year, assume actual year
	if len(dateSplitted) == dayAndMonthLen {
		dateSplitted = append(dateSplitted, strconv.Itoa(time.Now().Year()))
	}

	dateString = dateSplitted[2] + wantDateSeparator + dateSplitted[0] + wantDateSeparator + dateSplitted[1]

	ans, err := time.Parse(time.DateOnly, dateString)
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: %s", ErrParsingDate, err.Error())
	}

	return ans, nil
}

// parseAmount converts a string representing an amount of money to a decimal.Decimal
// Returns ErrParsingAmount if the transformation if not possible
func parseAmount(s string) (decimal.Decimal, error) {
	ans, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Decimal{}, fmt.Errorf("%w: %s", ErrParsingAmount, err.Error())
	}

	return ans, nil
}
