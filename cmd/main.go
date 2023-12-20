package main

import (
	"fmt"
	"net/http"
	"time"
	"log"

	"go-parallel-transactions/domain"
	"go-parallel-transactions/router"
	"go-parallel-transactions/repository"
)

var (
	UserDB []*domain.User= make([]*domain.User, 0)
	UserChannel chan *domain.User = make(chan *domain.User, 1)
	TransactionChannel chan *domain.NewTransactionRequest = make(chan *domain.NewTransactionRequest, 1)
	workerCount int = 2
)


func main(){
	go StartWorkers()

	router.New(&UserDB, &UserChannel, &TransactionChannel)

	fmt.Println("Server listening on port 60061")
	log.Fatal(http.ListenAndServe(":60061", nil))
}


func StartWorkers(){
	for i := 0; i<workerCount; i++ {
		go func(){
			ticker := time.NewTicker(time.Duration(workerCount))
			defer ticker.Stop()
			for range ticker.C {
				repository.VerifyUser(UserChannel)
			}
		}()

		go func(){
			ticker := time.NewTicker(time.Duration(workerCount))
			defer ticker.Stop()
			for range ticker.C {
				repository.ProcessTransaction(&TransactionChannel, &UserChannel, &UserDB)
			}
		}()
	}
}

