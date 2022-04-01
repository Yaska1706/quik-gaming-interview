package api

import (
	"errors"
	"strconv"
)

type DebitWalletRequest struct {
	Amount   string `json:"amount"`
	WalletID string `json:"wallet_id"`
}

type CreditWalletRequest struct {
	Amount   string `json:"amount"`
	WalletID string `json:"wallet_id"`
}

type WalletBalane struct {
	WalletID int `json:"wallet_id"`
	Balance  int `json:"balance"`
}

type WalletService interface {
	AddDebit(walletID string, request DebitWalletRequest) error
	AddCredit(walletID string, request CreditWalletRequest) error
}

type WalletRepository interface {
	CreateDebit(wallet Wallet) error
	CreateCredit(wallet Wallet) error
	GetAllByWalletID(walletid string, wallets []Wallet) ([]Wallet, error)
}

type walletService struct {
	storage WalletRepository
}

func NewWalletService(walletrepo WalletRepository) WalletService {
	return &walletService{storage: walletrepo}
}

const (
	defaultdebit  = "0"
	defaultcredit = "0"
)

func (w *walletService) AddDebit(walletID string, request DebitWalletRequest) error {

	if err := w.validateAmount(request.Amount); err != nil {
		return err
	}

	newdebit := Wallet{
		WalletID: walletID,
		Debit:    request.Amount,
		Credit:   defaultcredit,
	}

	if err := w.storage.CreateDebit(newdebit); err != nil {
		return err
	}
	return nil
}

func (w *walletService) AddCredit(walletID string, request CreditWalletRequest) error {

	if err := w.validateAmount(request.Amount); err != nil {
		return err
	}

	newcredit := Wallet{
		WalletID: walletID,
		Credit:   request.Amount,
		Debit:    defaultdebit,
	}

	if err := w.storage.CreateCredit(newcredit); err != nil {
		return err
	}
	return nil
}

func (w *walletService) validateAmount(amount string) error {
	if amount == "" {
		return errors.New("debit cannot be blank")
	}

	newamount, _ := strconv.ParseFloat(amount, 64)
	if newamount < 0 {
		return errors.New("cannot be a negative")
	}

	return nil
}
