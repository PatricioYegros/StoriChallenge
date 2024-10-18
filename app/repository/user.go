package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type IUserRepository interface {
	//Search makes a SELECT statement where obtains the user_id of a specific email
	Search(db *sql.DB, email string) (userId int64, err error)
	UpdateBalance(db *sql.DB, balance float64) error
}

type UserRepository struct{}

func (repository UserRepository) Search(db *sql.DB, email string) (userId int64, err error) {
	err = db.QueryRow("SELECT user_id FROM user WHERE email = ?", email).Scan(&userId)
	if err != nil {
		log.Printf("Error %s when searching user", err)
		return 0, err
	}
	return userId, nil
}

func (repository UserRepository) UpdateBalance(db *sql.DB, balance float64, userId int64) error {
	query := "UPDATE user SET balance = ? where user_id = ?"
	_, err := db.Exec(query, balance, userId)
	if err != nil {
		log.Printf("Error %s when updating balance of user", err)
		return err
	}
	return nil
}
