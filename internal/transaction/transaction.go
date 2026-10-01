package transaction

import (
	"errors"
	"time"

	"wallet/internal/money"
)

var (
	ErrInvalidAmount = errors.New("amount must be positive")
	ErrInvalidType   = errors.New("unknown transaction type")
	ErrEmptyID       = errors.New("id must not be empty")
	ErrEmptyAccount  = errors.New("account id must not be empty")
)

type Type string

const (
	Deposit     Type = "deposit"
	Withdrawal  Type = "withdrawal"
	TransferIn  Type = "transfer_in"
	TransferOut Type = "transfer_out"
)

type Transaction struct {
	ID             string
	AccountID      string
	TransferID     string
	CounterpartyID string
	Type           Type
	Amount         money.Money
	BalanceAfter   money.Money
	Description    string
	CreatedAt      time.Time
}

// Params holds the input for New. Required: ID, AccountID, Type, Amount.
// TransferID and CounterpartyID are set for the two halves of a transfer.
type Params struct {
	ID             string
	AccountID      string
	TransferID     string
	CounterpartyID string
	Type           Type
	Amount         money.Money
	BalanceAfter   money.Money
	Description    string
}

func New(p Params) (*Transaction, error) {
	if p.ID == "" {
		return nil, ErrEmptyID
	}
	if p.AccountID == "" {
		return nil, ErrEmptyAccount
	}
	if p.Amount <= 0 {
		return nil, ErrInvalidAmount
	}

	switch p.Type {
	case Deposit, Withdrawal, TransferIn, TransferOut:
	default:
		return nil, ErrInvalidType
	}

	return &Transaction{
		ID:             p.ID,
		AccountID:      p.AccountID,
		TransferID:     p.TransferID,
		CounterpartyID: p.CounterpartyID,
		Type:           p.Type,
		Amount:         p.Amount,
		BalanceAfter:   p.BalanceAfter,
		Description:    p.Description,
		CreatedAt:      time.Now(),
	}, nil
}
