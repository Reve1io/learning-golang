package pointers

type Account struct {
	Owner   string
	Balance int
}

func Deposit(acc *Account, value int) {

	if value < 0 {
		return
	}

	acc.Balance += value
}
