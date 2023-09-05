package domain

import (
	"github.com/google/uuid"
	"time"
)

type TransactionType byte
type AccountType byte

const (
	TED TransactionType = 'T'
	PIX TransactionType = 'P'

	PJ AccountType = 'J'
	PF AccountType = 'F'
)

func CreateUserAccount(name, cpf, number, email string) UserAccount {
	var userId = uuid.New()
	return UserAccount{
		ID:        userId,
		Name:      name,
		Number:    number,
		Cpf:       cpf,
		Email:     email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func CreateTransaction(
	requestID string,
	from UserAccount,
	to UserAccount,
	value float64,
	transactionType TransactionType,
) PaymentTransaction {
	var newTransactionID uuid.UUID = uuid.New()
	// TODO handle validation and log
	var newRequestId, _ = uuid.Parse(requestID)
	return PaymentTransaction{
		ID:        newTransactionID,
		RequestID: newRequestId,
		From:      from,
		To:        to,
		Value:     value,
		Type:      transactionType,
		CreatedAt: time.Now().UTC(),
	}
}

func (t PaymentTransaction) End() PaymentTransaction {
	t.FinishedAt = time.Now().UTC()
	return t
}

type UserAccount struct {
	ID        uuid.UUID
	Name      string
	Cpf       string
	Number    string
	Email     string
	Type      AccountType
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaymentTransaction struct {
	ID         uuid.UUID
	RequestID  uuid.UUID
	From       UserAccount
	To         UserAccount
	Value      float64
	Type       TransactionType
	CreatedAt  time.Time
	FinishedAt time.Time
}
