package api

import (
	"errors"
	"fmt"
	"regexp"
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
	GetBalance(walletid string) (string, error)
}

type WalletRepository interface {
	CreateDebit(wallet Wallet) error
	CreateCredit(wallet Wallet) error
	GetAllByWalletID(walletid string) ([]Wallet, error)
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

func (w *walletService) GetBalance(walletid string) (string, error) {

	wallets, err := w.storage.GetAllByWalletID(walletid)
	if err != nil {
		return "", err
	}
	var storedebit []string
	var storecredit []string

	for _, wallet := range wallets {
		storecredit = append(storecredit, wallet.Credit)
		storedebit = append(storedebit, wallet.Debit)
	}

	creditsum := findSum(storecredit)
	debitsum := findSum(storedebit)

	balance := findbalance(creditsum, debitsum)

	return balance, nil
}

func (w *walletService) validateAmount(amount string) error {
	if amount == "" {
		return errors.New("debit cannot be blank")
	}

	value := regexp.MustCompile("[+]?([0-9]*[.])?[0-9]+")
	if !value.MatchString(amount) {
		return errors.New("value must be a number")
	}

	newamount, _ := strconv.ParseFloat(amount, 64)
	if newamount < 0 {
		return errors.New("cannot be a negative")
	}

	return nil
}

func findSum(array []string) string {
	arraytofloat := []float64{}
	for _, value := range array {
		if n, err := strconv.ParseFloat(value, 64); err == nil {
			arraytofloat = append(arraytofloat, n)
		}
	}
	var sum float64

	for _, value := range arraytofloat {
		sum += value
	}

	return fmt.Sprint(sum)
}

func findbalance(credit, debit string) string {
	newcredit, _ := strconv.ParseFloat(credit, 64)
	newdebit, _ := strconv.ParseFloat(debit, 64)

	balance := newcredit - newdebit

	return fmt.Sprint(balance)
}
