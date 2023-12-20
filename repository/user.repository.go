package repository

import (
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