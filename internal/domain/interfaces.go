package domain

type UserRepository interface {
	FromCpf(cpf string) (UserAccount, error)
	FromEmail(email string) (UserAccount, error)
	FromID(id string) (UserAccount, error)
	GetBalance(account UserAccount) float64
}

type TransactionRepository interface {
	Save(transaction PaymentTransaction)
}

type TransactionResponseSender interface {
	Send(transaction PaymentTransaction)
}
