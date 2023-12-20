package repository

import (
	"errors"
	"go-parallel-transactions/domain"
)

func CreateUser(payload *domain.CreateUserRequest, channel *chan *domain.User, userDB *[]*domain.User) *domain.User {

	// No need to use a Write lock. Channel performs synchronization
	user := &domain.User{
		Id: int64(len(*userDB) +1),
		Name: payload.Name,
		IsVerified: false,
		Balance: 1000,
	}
	*userDB = append(*userDB, user);
	*channel <- user
	return user
}


func VerifyUser(channel chan *domain.User){
	// No need to use a Write lock. Channel performs synchronization
	user := <- channel
	user.IsVerified = true
}

func NewTransaction(transaction *domain.NewTransactionRequest, channel *chan *domain.NewTransactionRequest, userDB *[]*domain.User) error {
	senderIndex := transaction.SenderId-1
	receiverIndex := transaction.ReceiverId -1

	if senderIndex < 0 || senderIndex >= int64(len(*userDB)) {
		return errors.New("senderId not found")
	}

	if receiverIndex < 0 || receiverIndex >= int64(len(*userDB)) {
		return errors.New("receiverId not found")
	}

	sender := (*userDB)[senderIndex]	
	if sender.Balance < transaction.Amount {
		return errors.New("insufficient balance")
	}

	*channel <- transaction

	return nil
}

func ProcessTransaction(transactionChannel *chan *domain.NewTransactionRequest, userChannel *chan *domain.User, userDB *[]*domain.User){
	transaction := <- *transactionChannel

	sender := (*userDB)[transaction.SenderId-1]
	receiver := (*userDB)[transaction.ReceiverId-1]
	// push unverified sender to verification queue
	if !sender.IsVerified {
		*userChannel <- sender
		return
	}

	if transaction.Amount < sender.Balance {
		sender.Balance -= transaction.Amount

		receiver.Balance += transaction.Amount
	}
}