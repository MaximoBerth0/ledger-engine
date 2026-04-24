package domain

import (
	"ledger-engine/internal/domain/values"
	"time"
)

type Transaction struct {
	ID             values.TransactionID
	Status         values.TransactionStatus
	RevertOf       *values.TransactionID // nil unless this is a reversal
	IdempotencyKey string
	Metadata       map[string]any
	PostedAt       time.Time
	CreatedAt      time.Time
}
