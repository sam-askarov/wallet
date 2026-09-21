package wallet

import (
	"errors"
	"wallet/internal/account"
	"wallet/internal/money"
)

var (
	ErrAccountExists   = errors.New("account already exists")
	ErrAccountNotFound = errors.New("account not found")
	ErrSameAccount     = errors.New("transfer to the same account")
)

type Wallet struct {
	accounts map[string]*account.Account
}

func NewWallet() *Wallet {
	return &Wallet{
		accounts: make(map[string]*account.Account),
	}
}

func (w *Wallet) CreateAccount(id, owner string) (*account.Account, error) {
	if _, ok := w.accounts[id]; ok {
		return nil, ErrAccountExists
	}

	acc := &account.Account{
		ID:      id,
		Owner:   owner,
		Balance: 0,
	}
	w.accounts[id] = acc

	return acc, nil
}

func (w *Wallet) GetAccount(id string) (*account.Account, error) {
	acc, ok := w.accounts[id]
	if !ok {
		return nil, ErrAccountNotFound
	}

	return acc, nil
}

func (w *Wallet) Transfer(fromID, toID string, amount money.Money) error {
	if fromID == toID {
		return ErrSameAccount
	}

	if amount <= 0 {
		return account.ErrInvalidAmount
	}

	from, err := w.GetAccount(fromID)
	if err != nil {
		return err
	}

	to, err := w.GetAccount(toID)
	if err != nil {
		return err
	}

	if err := from.Withdraw(amount); err != nil {
		return err
	}

	return to.Deposit(amount)
}
