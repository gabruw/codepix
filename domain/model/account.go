package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/asaskevich/govalidator"
)

type Account struct {
	Base      `valid:"required"`
	Bank      *Bank     `valid:"-"`
	PixKeys   []*PixKey `valid:"-"`
	Number    string    `json:"number" valid:"notnull"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}

	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	account := Account{
		Bank:      bank,
		Number:    number,
		OwnerName: ownerName,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}
