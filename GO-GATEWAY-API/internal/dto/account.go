package dto

import (
	"time")

type CreateAccount struct {
	Name string `json:"name"` 
	Email string `json:"email"`
}

type AccountOutputInput struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Balance float64 `json:"balance"`
	APIKey string `json:"api_key, omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToAccount(input CreateAccountInput) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

func FromAccount(account *domain.Account) AccountOutputInput {
	return AccountOutputInput{
		ID: account.ID,
		Name: account.Name,
		Email: account.Email,
		APIKey: account.APIKey,
		Balance: account.Balance,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}