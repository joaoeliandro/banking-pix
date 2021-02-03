package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/joaoeliandro/banking-pix-microservice/domain/model"
)

type PixKeyRepositoryDatabase struct {
	Database *gorm.DB
}

func (repository PixKeyRepositoryDatabase) AddBank(bank *model.Bank) error {
	err := repository.Database.Create(bank).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository PixKeyRepositoryDatabase) AddAccount(account *model.Account) error {
	err := repository.Database.Create(account).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository PixKeyRepositoryDatabase) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := repository.Database.Create(pixKey).Error
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

func (repository PixKeyRepositoryDatabase) FindKeyById(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	repository.Database.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, errors.New("no key was found")
	}

	return &pixKey, nil
}

func (repository PixKeyRepositoryDatabase) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	repository.Database.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, errors.New("no account found")
	}

	return &account, nil
}

func (repository PixKeyRepositoryDatabase) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	repository.Database.Preload("Bank").First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, errors.New("no bank found")
	}

	return &bank, nil
}
