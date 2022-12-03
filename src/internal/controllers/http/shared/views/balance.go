package views

import "github.com/0B1t322/BWG-Test/internal/models/aggregate"

type BalanceView struct {
	ID      string  `json:"id"      format:"uuid"`
	Balance float64 `json:"balance"`
}

func BalanceViewFrom(balance aggregate.Balance) BalanceView {
	return BalanceView{
		ID:      balance.ID.String(),
		Balance: balance.Balance.Balance,
	}
}
