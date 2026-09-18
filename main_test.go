package main

import (
	"errors"
	"testing"
)

func TestDeposit(t *testing.T) {
	acc := &Account{ID: "acc-1", Owner: "Anna", Balance: 0}

	err := acc.Deposit(10000)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if acc.Balance != 10000 {
		t.Errorf("balance = %v, expected balance = 10000", acc.Balance)
	}
}

func TestDepositNegative(t *testing.T) {
	acc := &Account{ID: "acc-1", Owner: "Anna", Balance: 0}

	err := acc.Deposit(-100)

	if !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("got %v, want %v", err, ErrInvalidAmount)
	}

	if acc.Balance != 0 {
		t.Errorf("balance = %v, expected balance = 0", acc.Balance)
	}
}

func TestWithdraw(t *testing.T) {
	acc := &Account{ID: "acc-1", Owner: "Anna", Balance: 10000}

	err := acc.Withdraw(3000)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if acc.Balance != 7000 {
		t.Errorf("balance = %v, expected balance = 7000", acc.Balance)
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {
	acc := &Account{ID: "acc-1", Owner: "Anna", Balance: 10000}

	err := acc.Withdraw(99999)

	if !errors.Is(err, ErrInsufficientFunds) {
		t.Fatalf("got %v, want %v", err, ErrInsufficientFunds)
	}

	if acc.Balance != 10000 {
		t.Errorf("balance = %v, expected balance = 10000", acc.Balance)
	}
}
