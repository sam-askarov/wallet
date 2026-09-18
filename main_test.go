package main

import (
	"errors"
	"testing"
	"wallet/internal/account"
	"wallet/internal/money"
	"wallet/internal/wallet"
)

func TestDeposit(t *testing.T) {
	acc := &account.Account{ID: "acc-1", Owner: "Anna", Balance: 0}

	err := acc.Deposit(10000)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if acc.Balance != 10000 {
		t.Errorf("balance = %v, expected balance = 10000", acc.Balance)
	}
}

func TestDepositNegative(t *testing.T) {
	acc := &account.Account{ID: "acc-1", Owner: "Anna", Balance: 0}

	err := acc.Deposit(-100)

	if !errors.Is(err, account.ErrInvalidAmount) {
		t.Fatalf("got %v, want %v", err, account.ErrInvalidAmount)
	}

	if acc.Balance != 0 {
		t.Errorf("balance = %v, expected balance = 0", acc.Balance)
	}
}

func TestWithdrawTable(t *testing.T) {
	tests := []struct {
		name        string
		balance     money.Money
		amount      money.Money
		wantErr     error
		wantBalance money.Money
	}{
		{
			name:        "успешное снятие",
			balance:     10000,
			amount:      3000,
			wantErr:     nil,
			wantBalance: 7000,
		},
		{
			name:        "недостаточно средств",
			balance:     10000,
			amount:      99999,
			wantErr:     account.ErrInsufficientFunds,
			wantBalance: 10000,
		},
		{
			name:        "отрицательная сумма",
			balance:     10000,
			amount:      -500,
			wantErr:     account.ErrInvalidAmount,
			wantBalance: 10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := &account.Account{ID: "acc-1", Owner: "Anna", Balance: tt.balance}

			err := acc.Withdraw(tt.amount)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if acc.Balance != tt.wantBalance {
				t.Errorf("balance = %v, want %v", acc.Balance, tt.wantBalance)
			}
		})
	}
}

func TestTransferTable(t *testing.T) {
	tests := []struct {
		name            string
		balanceFrom     money.Money
		balanceTo       money.Money
		fromID          string
		toID            string
		amount          money.Money
		wantErr         error
		wantBalanceFrom money.Money
		wantBalanceTo   money.Money
	}{
		{
			name:            "успешный перевод",
			balanceFrom:     10000,
			balanceTo:       0,
			fromID:          "acc-1",
			toID:            "acc-2",
			amount:          7000,
			wantErr:         nil,
			wantBalanceFrom: 3000,
			wantBalanceTo:   7000,
		},
		{
			name:            "перевод самому себе",
			balanceFrom:     10000,
			fromID:          "acc-1",
			toID:            "acc-1",
			amount:          5000,
			wantErr:         wallet.ErrSameAccount,
			wantBalanceFrom: 10000,
		},
		{
			name:            "несуществующий получатель",
			balanceFrom:     1000,
			fromID:          "acc-1",
			toID:            "acc-99",
			amount:          5000,
			wantErr:         wallet.ErrAccountNotFound,
			wantBalanceFrom: 1000,
		},
		{
			name:            "недостаточно средств",
			balanceFrom:     10000,
			balanceTo:       0,
			fromID:          "acc-1",
			toID:            "acc-2",
			amount:          999900,
			wantErr:         account.ErrInsufficientFunds,
			wantBalanceFrom: 10000,
			wantBalanceTo:   0,
		},
		{
			name:            "отрицательная сумма",
			balanceFrom:     10000,
			balanceTo:       0,
			fromID:          "acc-1",
			toID:            "acc-2",
			amount:          -1000,
			wantErr:         account.ErrInvalidAmount,
			wantBalanceFrom: 10000,
			wantBalanceTo:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wallet.NewWallet()

			acc1, _ := w.CreateAccount("acc-1", "Anna")
			acc2, _ := w.CreateAccount("acc-2", "Jane")
			acc1.Balance = tt.balanceFrom
			acc2.Balance = tt.balanceTo

			err := w.Transfer(tt.fromID, tt.toID, tt.amount)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if acc1.Balance != tt.wantBalanceFrom {
				t.Errorf("balance = %v, want %v", acc1.Balance, tt.wantBalanceFrom)
			}

			if acc2.Balance != tt.wantBalanceTo {
				t.Errorf("balance = %v, want %v", acc2.Balance, tt.wantBalanceTo)
			}
		})
	}
}
