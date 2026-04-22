package domain

import (
	"ledger-engine/internal/domain/values"
	"time"
)

type Account struct {
	ID        values.AccountID
	Type      values.AccountType
	Currency  values.Currency
	Status    values.AccountStatus
	Metadata  map[string]any
	CreatedAt time.Time
	ClosedAT  *time.Time
}
