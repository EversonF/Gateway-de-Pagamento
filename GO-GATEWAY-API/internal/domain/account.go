package domain

import ("time"
import "github.com/google/uuid"
import "crypto/rand"
import "encoding/hex"
import "sync")

type Account struct {
	ID string
	Name string
	Email string
	APIKey string
	Balance float64
	mu sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateAPIKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func NewAccount(name, email string) *Account {
	
	account := &Account{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		Balance: 0.0,
		APIKey: generateAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

func (a *Account) AddBalance(amount float64){
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	a.UpdatedAt = time.Now()
	a.mu.Unlock()
}