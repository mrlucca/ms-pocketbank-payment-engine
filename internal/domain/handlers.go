package domain

func HandleTransactionFromUserCpf(
	userRepository UserRepository,
	transactionRepository TransactionRepository,
	sender TransactionResponseSender,
	inputDTO TransactionFromCpfDTO,
) {
	var (
		from UserAccount
		to   UserAccount
		err  error
	)
	from, err = userRepository.FromCpf(inputDTO.FromAccountID)
	if err != nil {
		panic("TODO handle user not found from account")
	}

	to, err = userRepository.FromCpf(inputDTO.ToCpf)
	if err != nil {
		panic("TODO handle user not found")
	}
	handlePaymentTransaction(
		userRepository,
		transactionRepository,
		sender,
		from,
		to,
		inputDTO.BaseTransaction,
	)
}

func HandleTransactionFromUserEmail(
	userRepository UserRepository,
	transactionRepository TransactionRepository,
	sender TransactionResponseSender,
	inputDTO TransactionFromEmailDTO,
) {
	var (
		from UserAccount
		to   UserAccount
		err  error
	)
	from, err = userRepository.FromCpf(inputDTO.FromAccountID)
	if err != nil {
		panic("TODO handle user not found from account")
	}

	to, err = userRepository.FromCpf(inputDTO.ToEmail)
	if err != nil {
		panic("TODO handle user not found")
	}
	handlePaymentTransaction(
		userRepository,
		transactionRepository,
		sender,
		from,
		to,
		inputDTO.BaseTransaction,
	)
}

func handlePaymentTransaction(
	userRepository UserRepository,
	transactionRepository TransactionRepository,
	sender TransactionResponseSender,
	from, to UserAccount,
	request BaseTransaction,
) {
	var (
		transaction PaymentTransaction
		fromBalance float64
	)

	fromBalance = userRepository.GetBalance(from)
	if fromBalance < request.Value {
		// TODO HANDLE WHEN NO CONTAINS BALANCE FOR COMPLETE TRANSACTION
		return
	}

	transaction = CreateTransaction(
		request.RequestID,
		from,
		to,
		request.Value,
		request.Type,
	)
	go transactionRepository.Save(transaction)
	go sender.Send(transaction)
}

type BaseTransaction struct {
	RequestID     string
	FromAccountID string
	Type          TransactionType
	Value         float64
}

type TransactionFromCpfDTO struct {
	BaseTransaction
	ToCpf string
}

type TransactionFromEmailDTO struct {
	BaseTransaction
	ToEmail string
}
