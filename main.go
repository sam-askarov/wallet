package main

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInvalidAmount     = errors.New("amount must be positive")
	ErrAccountExists     = errors.New("account already exists")
	ErrAccountNotFound   = errors.New("account not found")
)

type Money int64

func (m Money) String() string {
	sign := ""

	if m < 0 {
		sign = "-"
		m = -m
	}
	return fmt.Sprintf("%s%d.%02d", sign, m/100, m%100)
}

type Account struct {
	ID      string
	Owner   string
	Balance Money
}

func (a *Account) Deposit(amount Money) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount Money) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if amount > a.Balance {
		return ErrInsufficientFunds
	}

	a.Balance -= amount
	return nil
}

type Wallet struct {
	accounts map[string]*Account
}

func NewWallet() *Wallet {
	return &Wallet{
		accounts: make(map[string]*Account),
	}
}

func (w *Wallet) CreateAccount(id, owner string) (*Account, error) {
	if _, ok := w.accounts[id]; ok {
		return nil, ErrAccountExists
	}

	acc := &Account{
		ID:      id,
		Owner:   owner,
		Balance: 0,
	}
	w.accounts[id] = acc

	return acc, nil
}

func (w *Wallet) GetAccount(id string) (*Account, error) {
	acc, ok := w.accounts[id]
	if !ok {
		return nil, ErrAccountNotFound
	}

	return acc, nil
}

func main() {
	w := NewWallet()

	acc, err := w.CreateAccount("acc-1", "Anna")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if err := acc.Deposit(100000); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(acc.Balance) // 1000.00

	// тот же счёт достаём из кошелька
	same, _ := w.GetAccount("acc-1")
	fmt.Println(same.Balance) // 1000.00 — тот же объект!

	// дубликат
	_, err = w.CreateAccount("acc-1", "Boris")
	fmt.Println(err) // account already exists

	// несуществующий
	_, err = w.GetAccount("acc-99")
	fmt.Println(err)
}
