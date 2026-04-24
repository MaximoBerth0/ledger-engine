package values

type TransactionID string

type TransactionStatus string

const (
	TransactionStatusPending  TransactionStatus = "pending"
	TransactionStatusPosted   TransactionStatus = "posted"
	TransactionStatusReverted TransactionStatus = "reverted"
)
