package transactionqueue

import (
	"context"
	"sync"

	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type TransactionExecuter interface {
	ExecuteTransactions(
		ctx context.Context,
		transactions []aggregate.Transaction,
	) error
}

type TransactionQueue struct {
	queue map[uuid.UUID][]aggregate.Transaction
	sync.Mutex

	executer TransactionExecuter
}

func NewTransactionQueue(executer TransactionExecuter) *TransactionQueue {
	return &TransactionQueue{
		queue:    make(map[uuid.UUID][]aggregate.Transaction),
		executer: executer,
	}
}

func (tq *TransactionQueue) AddTransactions(transactions ...aggregate.Transaction) {
	tq.Lock()
	defer tq.Unlock()

	tmap := transactionsToMap(transactions)

	for userID := range tmap {
		tq.queue[userID] = append(tq.queue[userID], tmap[userID]...)
	}
}

func (tq *TransactionQueue) ExecuteTransactions() {
	tq.Lock()
	defer tq.Unlock()

	for userID, transactions := range tq.queue {
		go func(transactions []aggregate.Transaction) {
			if err := tq.executer.ExecuteTransactions(
				context.Background(),
				transactions,
			); err != nil {
				logrus.WithFields(
					logrus.Fields{
						"service": "TransactionQueue",
					},
				).Error(err)
			}
		}(transactions)

		delete(tq.queue, userID)
	}
}

func transactionsToMap(transactions []aggregate.Transaction) map[uuid.UUID][]aggregate.Transaction {
	tmap := make(map[uuid.UUID][]aggregate.Transaction)
	for _, transaction := range transactions {
		tmap[transaction.UserID] = append(tmap[transaction.UserID], transaction)
	}
	return tmap
}
