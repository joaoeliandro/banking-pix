package usecase

import (
	"errors"

	"github.com/joaoeliandro/banking-pix-microservice/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixKeyRepository      model.PixKeyRepositoryInterface
}

func (transactionUseCase *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKind string, description string) (*model.Transaction, error) {
	account, err := transactionUseCase.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := transactionUseCase.PixKeyRepository.FindKeyByKind(pixKeyTo, pixKeyKind)

	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)
	if err != nil {
		return nil, err
	}

	transactionUseCase.TransactionRepository.Save(transaction)

	if transaction.ID == "" {
		return transaction, nil
	}

	return nil, errors.New("unable to process this transaction")
}

func (transactionUseCase *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := transactionUseCase.TransactionRepository.Find(transactionId)

	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err = transactionUseCase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (transactionUseCase *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := transactionUseCase.TransactionRepository.Find(transactionId)

	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionCompleted

	err = transactionUseCase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (transactionUseCase *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := transactionUseCase.TransactionRepository.Find(transactionId)

	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = transactionUseCase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
