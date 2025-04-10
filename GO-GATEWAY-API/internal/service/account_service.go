package service

import ("github.com/devfullcycle/go-gateway-api/internal/domain"

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}