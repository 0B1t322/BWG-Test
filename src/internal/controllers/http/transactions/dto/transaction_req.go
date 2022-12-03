package dto

type (
	OperationType string

	TransactionReq struct {
		UserID    string        `json:"useId"     format:"uuid"`
		Amount    float64       `json:"amount"    format:"float"`
		Operation OperationType `json:"operation"                enums:"add,sub"`
	}

	TransactionsReq struct {
		Transactions []TransactionReq `json:"transactions"`
	}
)

const (
	Add OperationType = "add"
	Sub OperationType = "sub"
)
