package repository

import (
	"fmt"

	"github.com/gabruw/codepix/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (repo PixKeyRepositoryDb) Register(transaction *model.Transaction) error {
	err := repo.Db.Create(transaction).Error
	if err != nil {
		return nil
	}

	return nil
}

func (repo PixKeyRepositoryDb) Save(transaction *model.Transaction) error {
	err := repo.Db.Save(transaction).Error
	if err != nil {
		return nil
	}

	return nil
}

func (repo PixKeyRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	repo.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction found")
	}

	return &transaction, nil
}
