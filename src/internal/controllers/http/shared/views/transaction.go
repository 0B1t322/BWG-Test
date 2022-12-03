package views

import (
	"time"

	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
)

type (
	TransactionView struct {
		ID         string    `json:"id"         format:"uuid"`
		UserID     string    `json:"userId"     format:"uuid"`
		Amount     float64   `json:"amount"     format:"float"`
		Status     string    `json:"status"                        enums:"created,success,denied"`
		CreatedAt  time.Time `json:"createdAt"  format:"date-time"`
		ExecutedAt time.Time `json:"executedAt" format:"date-time"`
	}

	TransactionsView struct {
		Transactions []TransactionView `json:"transactions"`
	}
)

func TransactionStatusToString(status aggregate.TransactionStatus) string {
	switch status {
	case aggregate.TransactionStatusCreated:
		return "created"
	case aggregate.TransactionStatusSuccess:
		return "success"
	case aggregate.TransactionStatusDenied:
		return "denied"
	}
	return ""
}

func TransactionViewFrom(transaction aggregate.Transaction) TransactionView {
	return TransactionView{
		ID:         transaction.ID.String(),
		UserID:     transaction.UserID.String(),
		Amount:     transaction.Amount,
		Status:     TransactionStatusToString(transaction.Status),
		CreatedAt:  transaction.CreatedAt,
		ExecutedAt: transaction.ExecutedAt,
	}
}

func TransactionsViewFrom(transactions []aggregate.Transaction) TransactionsView {
	transactionsView := TransactionsView{
		Transactions: make([]TransactionView, 0, len(transactions)),
	}

	for _, transaction := range transactions {
		transactionsView.Transactions = append(
			transactionsView.Transactions,
			TransactionViewFrom(transaction),
		)
	}
	return transactionsView
}
