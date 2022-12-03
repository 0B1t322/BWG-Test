package dto

type GetTransactionsReq struct {
	UserID string `json:"-" format:"uuid" uri:"userId"`
}
