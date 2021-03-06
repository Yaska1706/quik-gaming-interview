package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yaska1706/quik-gaming-interview/pkg/api"
)

type Storage interface {
	CreateDebit(wallet api.Wallet) error
	CreateCredit(wallet api.Wallet) error
	GetAllByWalletID(walletid string) ([]api.Wallet, error)
}

type storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) Storage {
	return &storage{db: db}
}

func (s *storage) CreateDebit(wallet api.Wallet) error {
	if wallet := s.db.Create(&wallet); wallet.Error != nil {
		return fmt.Errorf(wallet.Error.Error())
	}
	return nil
}

func (s *storage) CreateCredit(wallet api.Wallet) error {
	if wallet := s.db.Create(&wallet); wallet.Error != nil {
		return fmt.Errorf(wallet.Error.Error())
	}
	return nil
}

func (s *storage) GetAllByWalletID(walletid string) ([]api.Wallet, error) {
	var wallets []api.Wallet
	if wallet := s.db.Where("wallet_id = ?", walletid).Find(&wallets); wallet.Error != nil {
		return nil, fmt.Errorf(wallet.Error.Error())
	}

	return wallets, nil
}
