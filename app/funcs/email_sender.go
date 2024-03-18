package funcs

import (
	"errors"
)

type EmailSender interface {
	// Send sends an email to the reciever with the subject and body received by parameter.
	// Returns ErrSendingEmail if an error is produced
	Send(subject string, body string, reciever string) error
}

var ErrSendingEmail = errors.New("error while sending email")
