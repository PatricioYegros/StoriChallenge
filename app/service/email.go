package service

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"time"

	"github.com/patricioyegros/storichallenge/app/funcs"
	"github.com/shopspring/decimal"
)

type IEmailService interface {
	// Send formats the data received by parameters into the Stori mail and sends it by email
	Send(
		transactionBalance decimal.Decimal,
		transactionsPerMonth []TransactionsPerMonth,
		avgDebit, avgCredit decimal.Decimal,
		receiver string,
	) error
}

type EmailService struct {
	EmailSender funcs.EmailSender
	Template    *template.Template
}

const emailSubject = "Stori Challenge - Transaction summary"

type transactionsPerMonthEmailData struct {
	Month string
	Value string
}

type emailData struct {
	Date                     string
	UserBalance              string
	TransactionsBalance      string
	AvgDebit                 string
	AvgCredit                string
	TransactionsPerMonthList []transactionsPerMonthEmailData
}

// send formats the data received by parameters into the Stori mail and sends it
func (emailService EmailService) Send(
	transactionsBalance decimal.Decimal,
	transactionsPerMonth []TransactionsPerMonth,
	avgDebit, avgCredit decimal.Decimal,
	receiver string,
) error {
	var htmlBuffer bytes.Buffer

	err := emailService.Template.Execute(&htmlBuffer, emailData{
		Date:                     time.Now().Format(time.DateTime),
		TransactionsBalance:      transactionsBalance.String(),
		AvgDebit:                 avgDebit.String(),
		AvgCredit:                avgCredit.String(),
		TransactionsPerMonthList: transactionsPerMonthToEmailData(transactionsPerMonth),
	})
	if err != nil {
		return fmt.Errorf("%w: %s", funcs.ErrSendingEmail, err.Error())
	}

	return emailService.EmailSender.Send(emailSubject, htmlBuffer.String(), receiver)
}

// transactionsPerMonthToEmailData transforms a list of transactions per month to
// a list of transactionsPerMonthEmailData that can be render into the email template
func transactionsPerMonthToEmailData(transactionsPerMonthList []TransactionsPerMonth) []transactionsPerMonthEmailData {
	currentYear := time.Now().Year()

	result := make([]transactionsPerMonthEmailData, 0, len(transactionsPerMonthList))

	for _, transactionsPerMonth := range transactionsPerMonthList {
		var monthText string

		// add year if not current year
		if transactionsPerMonth.Month.Year() != currentYear {
			monthText = fmt.Sprintf("%s %d", transactionsPerMonth.Month.Month().String(), transactionsPerMonth.Month.Year())
		} else {
			monthText = transactionsPerMonth.Month.Month().String()
		}

		result = append(result, transactionsPerMonthEmailData{
			Month: monthText,
			Value: strconv.Itoa(transactionsPerMonth.Amount),
		})
	}

	return result
}
