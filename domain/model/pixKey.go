package model

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/asaskevich/govalidator"
)

type PixKeyRepositoryInteface interface {
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Account   *Account `valid:"-"`
	Key       string   `json:"key" valid:"notnull"`
	Kind      string   `json:"kind" valid:"notnull"`
	Status    string   `json:"status" valid:"notnull"`
	AccountID string   `json:"account_id" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalid type key")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewPixKey(account *Account, key string, kind string) (*PixKey, error) {
	pixKey := PixKey{
		Key:     key,
		Kind:    kind,
		Account: account,
		Status:  "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
