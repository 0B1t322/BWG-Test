package aggregate

type TransactionSort []Transaction

func (t TransactionSort) Len() int {
	return len(t)
}

func (t TransactionSort) Less(i, j int) bool {
	return t[i].CreatedAt.Before(t[j].CreatedAt)
}

func (t TransactionSort) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
