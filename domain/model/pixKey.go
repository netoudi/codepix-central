package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" gorm:"type:float" valid:"notnull"`
	Key       string   `json:"key" gorm:"type:float"  valid:"notnull"`
	Account   *Account `json:"account" valid:"-"`
	AccountID string   `json:"account_id" gorm:"column:type:float"  valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	FindBank(id string) (*Bank, error)
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)
	if err != nil {
		return err
	}

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalid status")
	}

	return nil
}

func NewPixKey(bank *Bank, number string, ownerName string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    "",
		Key:     "",
		Account: nil,
		Status:  "",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
