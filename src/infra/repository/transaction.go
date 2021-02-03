package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/joaoeliandro/banking-pix-microservice/domain/model"
)

type TransactionRepositoryDatabase struct {
	Database *gorm.DB
}

func (repository *TransactionRepositoryDatabase) Register(transaction *model.Transaction) error {
	err := repository.Database.Create(transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *TransactionRepositoryDatabase) Save(transaction *model.Transaction) error {
	err := repository.Database.Save(transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *TransactionRepositoryDatabase) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	repository.Database.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, errors.New("no transaction was found")
	}

	return &transaction, nil
}
