package transactions

import (
	"sort"

	"github.com/0B1t322/BWG-Test/internal/controllers/http/transactions/dto"
	transactionsrv "github.com/0B1t322/BWG-Test/internal/domain/transaction/service"
	"github.com/0B1t322/BWG-Test/internal/models/aggregate"
	"github.com/0B1t322/BWG-Test/internal/services/transactionqueue"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionsController struct {
	transactionsService transactionsrv.TransactionService
	queue               *transactionqueue.TransactionQueue
}

func NewTransactionsController(
	transactionsService transactionsrv.TransactionService,
	queue *transactionqueue.TransactionQueue,
) *TransactionsController {
	return &TransactionsController{
		transactionsService: transactionsService,
		queue:               queue,
	}
}

func (t *TransactionsController) Build(r gin.IRouter) {
	transactions := r.Group("/transactions")
	{
		transactions.GET("/:userId", t.GetTransactions)
		transactions.POST("", t.AddTransactions)
	}
}

// GetTransactions
// @Summary Get transactions
// @Description Get transactions for user
// @Tags transactions
// @Produce json
// @Param userId path string true "User ID"
// @Router /transactions/{userId} [get]
// @Success 200 {object} TransactionsView
func (t *TransactionsController) GetTransactions(c *gin.Context) {
	var req GetTransactionsReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	transactions, err := t.transactionsService.GetAllTransactionsForUser(c, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, TransactionsViewFrom(transactions))
}

// AddTransactions
// @Summary Add transactions
// @Description Add transactions for user
// @Tags transactions
// @Accept json
// @Produce json
// @Param body body TransactionsReq true "Add transactions request"
// @Router /transactions [post]
// @Success 200 {object} TransactionsView
func (t *TransactionsController) AddTransactions(c *gin.Context) {
	var req TransactionsReq
	{
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	var ts []aggregate.Transaction

	for _, r := range req.Transactions {
		id, err := uuid.Parse(r.UserID)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		created, err := t.transactionsService.CreateTransaction(
			c,
			r.Amount,
			OperationTypeFromString(string(r.Operation)),
			id,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ts = append(ts, created)
	}

	sort.Sort(aggregate.TransactionSort(ts))

	t.queue.AddTransactions(ts...)
	t.queue.ExecuteTransactions()
}

func OperationTypeFromString(opType string) aggregate.OperationType {
	switch opType {
	case string(dto.Add):
		return aggregate.OperationTypeAdd
	case string(dto.Sub):
		return aggregate.OperationTypeSub
	}

	return aggregate.OperationTypeAdd
}
