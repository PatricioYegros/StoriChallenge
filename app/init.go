package app

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/patricioyegros/storichallenge/app/funcs"
	"github.com/patricioyegros/storichallenge/app/repository"
	"github.com/patricioyegros/storichallenge/app/service"
)

const (
	emailTemplate    = "template/email.html"
	DBUrlEnvVar      = "MYSQL_URL"
	DBPortEnvVar     = "MYSQL_PORT"
	DBUserEnvVar     = "MYSQL_USER"
	DBPasswordEnvVar = "MYSQL_PASSWORD"
	DBNameEnvVar     = "MYSQL_NAME"
)

var (
	ErrDatabaseNotConfigured = errors.New("database env variables not configured")
)

//go:embed template/email.html
var htmlFS embed.FS

// NewService creates a new service.Service.
// injects the email sender and the apppasword in the email sender.
func NewService(db *sql.DB) *service.Service {

	return &service.Service{
		TransactionsReader: service.TransactionsReader{
			CSVReader: funcs.FileReader{},
		},
		TransactionService: service.TransactionService{
			DB:                    db,
			UserRepository:        repository.UserRepository{},
			TransactionRepository: repository.TransactionRepository{},
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

// Creating DB Connection
func NewDBConnection() (*sql.DB, error) {
	dbUrl := os.Getenv(DBUrlEnvVar)
	dbPort := os.Getenv(DBPortEnvVar)
	dbUser := os.Getenv(DBUserEnvVar)
	dbPassword := os.Getenv(DBPasswordEnvVar)
	dbName := os.Getenv(DBNameEnvVar)

	if dbUrl == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return nil, ErrDatabaseNotConfigured
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbUrl, dbPort, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
