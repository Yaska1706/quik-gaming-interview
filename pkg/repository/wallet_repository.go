package repository

import "strconv"

type WalletRepository struct{}

var walletRepository *WalletRepository

func GetWalletRepository() *WalletRepository {
	if walletRepository == nil {
		walletRepository = &WalletRepository{}
	}
	return walletRepository
}
func (r *WalletRepository) Get(id string) (*Wallet, error) {
	var wallet Wallet
	where := Wallet{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &wallet, []string{})
	if err != nil {
		return nil, err
	}
	return &wallet, err
}

func (r *WalletRepository) All() (*[]Wallet, error) {
	var tasks []Wallet
	err := Find(&Wallet{}, &tasks, []string{}, "id asc")
	return &tasks, err
}

func (r *WalletRepository) Query(q *Wallet) (*[]Wallet, error) {
	var tasks []Wallet
	err := Find(&q, &tasks, []string{}, "id asc")
	return &tasks, err
}

func (r *WalletRepository) Add(task *Wallet) error {
	err := Create(&task)
	err = Save(&task)
	return err
}
