package account

import (
	"errors"
	"testing"
	"wallet/internal/money"
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
			wantErr:     ErrInsufficientFunds,
			wantBalance: 10000,
		},
		{
			name:        "отрицательная сумма",
			balance:     10000,
			amount:      -500,
			wantErr:     ErrInvalidAmount,
			wantBalance: 10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := &Account{ID: "acc-1", Owner: "Anna", Balance: tt.balance}

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
