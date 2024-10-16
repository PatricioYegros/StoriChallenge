package app

import (
	"embed"
	"html/template"

	"github.com/patricioyegros/storichallenge/app/funcs"
	"github.com/patricioyegros/storichallenge/app/service"
)

const (
	emailTemplate = "template/email.html"
)

//go:embed template/*
var htmlFS embed.FS

// NewService creates a new service.Service.
// injects the email sender and the apppasword in the email sender.
func NewService() *service.Service {

	return &service.Service{
		TransactionsReader: service.TransactionsReader{
			CSVReader: funcs.FileReader{},
		},
		EmailService: service.EmailService{
			Template: template.Must(template.ParseFS(htmlFS, emailTemplate)),
			EmailSender: funcs.EjectEmail{
				Email:    "patricioyegrosstori@gmail.com",
				Password: "lbeu ecrz auhd xfnv",
			},
		},
	}
}
