package account

import "github.com/arensusu/pionex/domain"

type AccountService struct {
	client domain.Client
}

func NewAccountService(c domain.Client) *AccountService {
	return &AccountService{client: c}
}
