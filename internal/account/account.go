package account

import (
	"errors"
	"wallet/internal/money"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInvalidAmount     = errors.New("amount must be positive")
)

type Account struct {
	ID      string
	Owner   string
	Balance money.Money
}

func (a *Account) Deposit(amount money.Money) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount money.Money) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if amount > a.Balance {
		return ErrInsufficientFunds
	}

	a.Balance -= amount
	return nil
}
