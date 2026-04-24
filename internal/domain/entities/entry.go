package domain

import (
	"ledger-engine/internal/domain/values"
	"time"
)

type Entry struct {
	ID            values.EntryID
	TransactionID values.TransactionID
	AccountID     values.AccountID
	Direction     values.EntryDirection
	Amount        values.Money
	CreatedAt     time.Time
}
