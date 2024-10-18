package repository

import (
	"github.com/patricioyegros/storichallenge/app/models"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	Create(db *gorm.DB, transactions []models.Transaction) error
}

type TransactionRepository struct{}

func (repository TransactionRepository) Create(db *gorm.DB, transactions []models.Transaction) error {
	return db.Create(transactions).Error
}
