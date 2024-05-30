package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) *Account {
	if client == nil {
		return nil
	}
	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

func (a *Account) Credit(amout float64) {
	a.Balance += amout
	a.UpdatedAt = time.Now()
}

func (a *Account) Debit(amout float64) {
	a.Balance -= amout
	a.UpdatedAt = time.Now()
}
