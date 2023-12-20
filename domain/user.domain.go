package domain

type User struct{
	Id int64	`json:"id"`
	Balance float64	`json:"balance"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
}


type CreateUserRequest struct {
	Name string `json:"name" validator:"required"`
}



type NewTransactionRequest struct {
	SenderId int64 `json:"senderId" validator:"required"`
	ReceiverId int64 `json:"receiverId" validator:"required"`
	Amount float64 `json:"amount" validator:"required"`
}
