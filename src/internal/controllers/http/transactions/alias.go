package transactions

import (
	"github.com/0B1t322/BWG-Test/internal/controllers/http/shared/views"
	"github.com/0B1t322/BWG-Test/internal/controllers/http/transactions/dto"
)

type (
	TransactionReq     = dto.TransactionReq
	TransactionsReq    = dto.TransactionsReq
	GetTransactionsReq = dto.GetTransactionsReq
)

type (
	TransactionsView = views.TransactionsView
)

var (
	TransactionsViewFrom = views.TransactionsViewFrom
)
