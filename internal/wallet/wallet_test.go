package wallet

import (
	"errors"
	"testing"
	"wallet/internal/account"
	"wallet/internal/money"
)

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
			wantErr:         ErrSameAccount,
			wantBalanceFrom: 10000,
		},
		{
			name:            "несуществующий получатель",
			balanceFrom:     1000,
			fromID:          "acc-1",
			toID:            "acc-99",
			amount:          5000,
			wantErr:         ErrAccountNotFound,
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
			w := New()

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
