package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/patricioyegros/storichallenge/app/models"
	"github.com/patricioyegros/storichallenge/app/repository"
)

type ITransactionService interface {
	// Apply recieves the transactions total by parameter and applies it to the user's balance
	Apply(transactions []models.Transaction, transactionBalance float64, csvFileName string, db *sql.DB) error
}

type TransactionService struct {
	DB                    *sql.DB
	UserRepository        repository.UserRepository
	TransactionRepository repository.TransactionRepository
}

var ErrAplyingTransactions = errors.New("error aplying transactions")

func (service TransactionService) Apply(
	transactions []models.Transaction,
	transactionsBalance float64,
	csvFileName string,
	db *sql.DB,
) error {
	//First search user and obtain user_id with email
	userId, err := service.UserRepository.Search(db, transactions[0].Email)
	if err != nil {
		log.Printf("Error searching user. ERROR: %s", err)
		return err
	}
	err = service.InsertInDB(db, csvFileName, transactionsBalance, userId)
	if err != nil {
		log.Printf("Error inserting new Row in Transactions Table. ERROR: %s", err)
		return err
	}
	err = service.UserRepository.UpdateBalance(db, transactionsBalance, userId)
	if err != nil {
		log.Printf("Error updating balance of user")
		return err
	}
	return nil
}

func (service TransactionService) InsertInDB(db *sql.DB, csvFileName string, transactionsBalance float64, user int64) error {
	query := `INSERT INTO Transactions (file_name, created_date, amount, user_owner) VALUES (?, NOW(), ?, ?)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, csvFileName, transactionsBalance, user)
	if err != nil {
		log.Printf("Error %s when inserting row into transactions table", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d transactions created", rows)
	return nil
}
