package api

import "time"

type Model struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at"`
}

type Wallet struct {
	Model
	WalletID string `gorm:"column:wallet_id;" json:"wallet_id"`
	Debit    string `gorm:"column:debit;" json:"debit"`
	Credit   string `gorm:"column:credit;" json:"credit"`
}
