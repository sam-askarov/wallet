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
	ErrSameAccount       = errors.New("transfer to the same account")
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

func (w *Wallet) Transfer(fromID, toID string, amount Money) error {
	if fromID == toID {
		return ErrSameAccount
	}

	if amount <= 0 {
		return ErrInvalidAmount
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

func main() {
	w := NewWallet()

	acc1, _ := w.CreateAccount("acc-1", "Anna")
	acc2, _ := w.CreateAccount("acc-2", "Boris")

	_ = acc1.Deposit(100000) // 1000.00

	fmt.Println("=== до переводов ===")
	fmt.Println("acc-1:", acc1.Balance)
	fmt.Println("acc-2:", acc2.Balance)

	// 1. перевод самому себе
	err := w.Transfer("acc-1", "acc-1", 5000)
	fmt.Println("\n1. самому себе:", err)
	fmt.Println("   acc-1:", acc1.Balance)

	// 2. отрицательная сумма
	err = w.Transfer("acc-1", "acc-2", -100)
	fmt.Println("\n2. отрицательная сумма:", err)
	fmt.Println("   acc-1:", acc1.Balance)

	// 3. несуществующий получатель
	err = w.Transfer("acc-1", "acc-99", 5000)
	fmt.Println("\n3. нет получателя:", err)
	fmt.Println("   acc-1:", acc1.Balance)

	// 4. не хватает денег
	err = w.Transfer("acc-1", "acc-2", 999999)
	fmt.Println("\n4. мало денег:", err)
	fmt.Println("   acc-1:", acc1.Balance)

	// 5. успешный перевод
	err = w.Transfer("acc-1", "acc-2", 30000)
	fmt.Println("\n5. успех:", err)
	fmt.Println("   acc-1:", acc1.Balance)
	fmt.Println("   acc-2:", acc2.Balance)
}
