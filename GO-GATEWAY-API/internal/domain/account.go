package domain

import ("time"
import "github.com/google/uuid")

type Account struct {
	ID string
	Name string
	Email string
	APIKey string
	Balance float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(name, email string) *Account {
	
	account := &Account{
		ID: uuid.New().String(),
	}
}