package domain

type User struct{
	Id int64	`json:"id"`
	Balance float64	`json:"balance"`
	Name string `json:"name"`
	IsVerified bool `json:"isVerified"`
}


type CreateUserRequest struct {
	Name string `validator:"required"`
}
